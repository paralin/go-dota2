package dota2

import (
	"context"
	"errors"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/sirupsen/logrus"

	devents "github.com/paralin/go-dota2/events"
	bgcm "github.com/paralin/go-dota2/protocol"
	gcm "github.com/paralin/go-dota2/protocol"
	gcsdkm "github.com/paralin/go-dota2/protocol"
	gcsm "github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-dota2/socache"
	"github.com/paralin/go-dota2/state"
)

// AppID is the ID for dota2
const AppID = 570

// ErrNotReady is returned when the dota client is not ready.
var ErrNotReady = errors.New("the dota client is not ready to accept requests yet, or has just become unready")

// handlerMap is the map of message types to handler functions.
type handlerMap map[uint32]func(packet *gamecoordinator.GCPacket) error

// Dota2 handles the dota game handler.
type Dota2 struct {
	le     logrus.FieldLogger
	client *steam.Client
	cache  *socache.SOCache

	connectionCtxMtx    sync.Mutex
	connectionCtx       context.Context
	connectionCtxCancel context.CancelFunc

	stateMtx sync.Mutex
	state    state.Dota2State

	handlers handlerMap

	pendReqMtx sync.Mutex
	pendReqID  uint32
	pendReq    map[uint32]map[uint32]responseHandler
}

// New builds a new Dota2 handler.
func New(client *steam.Client, le logrus.FieldLogger) *Dota2 {
	c := &Dota2{
		le:      le,
		cache:   socache.NewSOCache(le),
		client:  client,
		pendReq: make(map[uint32]map[uint32]responseHandler),

		state: state.Dota2State{
			ConnectionStatus: gcsdkm.GCConnectionStatus_GCConnectionStatus_NO_SESSION,
		},
	}
	c.buildHandlerMap()
	client.GC.RegisterPacketHandler(c)
	return c
}

// GetCache returns the SO Cache.
func (d *Dota2) GetCache() *socache.SOCache {
	return d.cache
}

// Close kills any ongoing calls.
func (d *Dota2) Close() {
	d.connectionCtxMtx.Lock()
	if d.connectionCtxCancel != nil {
		d.connectionCtxCancel()
	}
	d.connectionCtxMtx.Unlock()
}

// buildHandlerMap builds the map of bound handler functions.
func (d *Dota2) buildHandlerMap() {
	d.handlers = handlerMap{
		// Welcome and conn status
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome):          d.handleClientWelcome,
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus): d.handleConnectionStatus,

		// Caching
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed):   d.handleCacheSubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_UpdateMultiple):    d.handleCacheUpdateMultiple,
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheUnsubscribed): d.handleCacheUnsubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_Destroy):           d.handleCacheDestroy,

		// System events
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingRequest): d.handlePingRequest,

		// Chat events
		uint32(gcm.EDOTAGCMsg_k_EMsgGCChatMessage): d.getEventEmitter(func() devents.Event {
			return &devents.ChatMessage{}
		}),
		uint32(gcm.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse): d.getEventEmitter(func() devents.Event {
			return &devents.JoinedChatChannel{}
		}),

		// Invites
		uint32(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated): d.getEventEmitter(func() devents.Event {
			return &devents.InvitationCreated{}
		}),
	}

	d.registerGeneratedHandlers()
}

// write sends a message to the game coordinator.
func (d *Dota2) write(messageType uint32, msg proto.Message) {
	d.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppID, messageType, msg))
}

// emit emits an event.
func (d *Dota2) emit(event interface{}) {
	d.client.Emit(event)
}

// accessState safely accesses the Dota2 state. return true if the state was changed / otherwise updated during the call.
func (d *Dota2) accessState(cb func(nextState *state.Dota2State) (bool, error)) error {
	d.stateMtx.Lock()
	defer d.stateMtx.Unlock()

	lastState := d.state
	changed, err := cb(&d.state)
	if err != nil {
		return err
	}
	if changed {
		d.emit(devents.ClientStateChanged{
			OldState: lastState,
			NewState: d.state,
		})
	}
	return nil
}

// unmarshalBody attempts to unmarshal a packet body.
func (d *Dota2) unmarshalBody(packet *gamecoordinator.GCPacket, msg proto.Message) (parseErr error) {
	defer func() {
		if parseErr != nil {
			d.le.WithError(parseErr).WithField("msgtype", packet.MsgType).Warn("unable to parse message")
		}
	}()

	return proto.Unmarshal(packet.Body, msg)
}

// HandleGCPacket handles an incoming game coordinator packet.
func (d *Dota2) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppID {
		return
	}

	le := d.le.WithField("msgtype", packet.MsgType)
	handler, ok := d.handlers[packet.MsgType]
	if ok && handler != nil {
		if err := handler(packet); err != nil {
			le.WithError(err).Warn("error handling gc msg")
			ok = false
		}
	}

	respHandled := d.handleResponsePacket(packet)
	if !ok && !respHandled {
		le.Debug("unhandled gc packet")
		d.emit(&devents.UnhandledGCPacket{
			Packet: packet,
		})
	}
}

// handlePingRequest handles an incoming ping request from the gc.
func (d *Dota2) handlePingRequest(packet *gamecoordinator.GCPacket) error {
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingResponse), &gcsdkm.CMsgGCClientPing{})
	return nil
}

// getEventEmitter returns a handler that emits an event, used by the generated code.
func (d *Dota2) getEventEmitter(ctor func() devents.Event) func(packet *gamecoordinator.GCPacket) error {
	return func(packet *gamecoordinator.GCPacket) error {
		obj := ctor()
		if err := d.unmarshalBody(packet, obj.GetEventBody()); err != nil {
			return err
		}

		d.emit(obj)
		return nil
	}
}
