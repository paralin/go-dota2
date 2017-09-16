package dota2

import (
	"context"
	"errors"
	"sync"

	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/protocol/gamecoordinator"
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	devents "github.com/paralin/go-dota2/events"
	// gcmm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	gcclm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	gccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	gcsdkm "github.com/paralin/go-dota2/protocol/gcsdk_gcmessages"
	gcsm "github.com/paralin/go-dota2/protocol/gcsystemmsgs"
	"github.com/paralin/go-dota2/state"
)

// AppID is the ID for dota2
const AppID = 570

// NotReadyErr is returned when the dota client is not ready.
var NotReadyErr error = errors.New("the dota client is not ready to accept requests yet, or has just become unready")

// handlerMap is the map of message types to handler functions.
type handlerMap map[uint32]func(packet *gamecoordinator.GCPacket) error

// Dota2 handles the dota game handler.
type Dota2 struct {
	le     *logrus.Entry
	client *steam.Client

	connectionCtxMtx    sync.Mutex
	connectionCtx       context.Context
	connectionCtxCancel context.CancelFunc

	stateMtx sync.Mutex
	state    state.Dota2State

	handlers handlerMap

	profileResponseHandlersMtx sync.Mutex
	profileResponseHandlers    map[uint32][]chan<- *gccm.CMsgDOTAProfileCard
}

// New builds a new Dota2 handler.
func New(client *steam.Client, le *logrus.Entry) *Dota2 {
	c := &Dota2{
		le:     le,
		client: client,
		state:  state.Dota2State{ConnectionStatus: gcsdkm.GCConnectionStatus_GCConnectionStatus_NO_SESSION},
		profileResponseHandlers: make(map[uint32][]chan<- *gccm.CMsgDOTAProfileCard),
	}
	c.buildHandlerMap()
	client.GC.RegisterPacketHandler(c)
	return c
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
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome):           d.handleClientWelcome,
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus):  d.handleConnectionStatus,
		uint32(gcm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse): d.handleGetProfileCardResponse,
	}
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

	var lastState state.Dota2State = d.state
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

// setConnectionStatus sets the connection status, and emits an event.
// NOTE: do not call from inside accessState.
func (d *Dota2) setConnectionStatus(
	connStatus gcsdkm.GCConnectionStatus,
	update *gcsdkm.CMsgConnectionStatus,
) {
	_ = d.accessState(func(ns *state.Dota2State) (changed bool, err error) {
		if ns.ConnectionStatus == connStatus {
			return false, nil
		}

		oldState := ns.ConnectionStatus
		d.le.WithField("old", oldState.String()).
			WithField("new", connStatus.String()).
			Debug("connection status changed")
		d.emit(&devents.GCConnectionStatusChanged{
			OldState: oldState,
			NewState: connStatus,
			Update:   update,
		})

		ns.ClearState() // every time the state changes, we lose the lobbies / etc
		ns.ConnectionStatus = connStatus
		d.connectionCtxMtx.Lock()
		if d.connectionCtxCancel != nil {
			d.connectionCtxCancel()
			d.connectionCtxCancel = nil
			d.connectionCtx = nil
		}
		if connStatus == gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION {
			d.connectionCtx, d.connectionCtxCancel = context.WithCancel(context.Background())
		}
		d.connectionCtxMtx.Unlock()
		return true, nil
	})
}

// SetPlaying informs Steam we are playing / not playing Dota 2.
func (d *Dota2) SetPlaying(playing bool) {
	if playing {
		d.client.GC.SetGamesPlayed(AppID)
	} else {
		d.client.GC.SetGamesPlayed()
		_ = d.accessState(func(ns *state.Dota2State) (changed bool, err error) {
			ns.ClearState()
			return true, nil
		})
	}
}

// SayHello says hello to the Dota2 server, in an attempt to get a session.
func (d *Dota2) SayHello() {
	partnerAccType := gcsdkm.PartnerAccountType_PARTNER_NONE
	engine := gcsdkm.ESourceEngine_k_ESE_Source2
	var clientSessionNeed uint32 = 104
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientHello), &gcsdkm.CMsgClientHello{
		ClientLauncher:    &partnerAccType,
		Engine:            &engine,
		ClientSessionNeed: &clientSessionNeed,
	})
}

// RequestProfileCard sends a request to the DOTA gc for a card and waits for the response.
// If you want to enable timeouts, use a context.WithTimeout
func (d *Dota2) RequestProfileCard(ctx context.Context, accountId uint32) (*gccm.CMsgDOTAProfileCard, error) {
	cctx, err := d.validateConnectionContext()
	if err != nil {
		return nil, err
	}

	handler := make(chan *gccm.CMsgDOTAProfileCard)
	{
		d.profileResponseHandlersMtx.Lock()
		handlerList := d.profileResponseHandlers[accountId]
		handlerList = append(handlerList, handler)
		d.profileResponseHandlers[accountId] = handlerList
		d.profileResponseHandlersMtx.Unlock()
	}

	// When returning, delete the handler from the list
	defer func() {
		d.profileResponseHandlersMtx.Lock()
		defer d.profileResponseHandlersMtx.Unlock()

		handlers := d.profileResponseHandlers[accountId]
		for i, h := range handlers {
			if h == handler {
				handlers[i] = handlers[len(handlers)-1]
				handlers = handlers[:len(handlers)-1]
				break
			}
		}
		if len(handlers) == 0 {
			delete(d.profileResponseHandlers, accountId)
		} else {
			d.profileResponseHandlers[accountId] = handlers
		}
	}()

	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard), &gcclm.CMsgClientToGCGetProfileCard{
		AccountId: &accountId,
	})

	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case <-cctx.Done():
		return nil, NotReadyErr
	case r := <-handler:
		return r, nil
	}
}

// validateConnectionContext checks if the client is ready or not.
func (d *Dota2) validateConnectionContext() (context.Context, error) {
	d.connectionCtxMtx.Lock()
	defer d.connectionCtxMtx.Unlock()

	cctx := d.connectionCtx
	if cctx == nil {
		return nil, NotReadyErr
	}

	select {
	case <-cctx.Done():
		return nil, NotReadyErr
	default:
		return cctx, nil
	}
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

// handleClientWelcome handles an incoming client welcome event.
func (d *Dota2) handleClientWelcome(packet *gamecoordinator.GCPacket) error {
	welcome := &gcsdkm.CMsgClientWelcome{}
	if err := d.unmarshalBody(packet, welcome); err != nil {
		return err
	}

	d.setConnectionStatus(gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil)
	d.emit(&devents.ClientWelcomed{Welcome: welcome})
	return nil
}

// handleConnectionStatus handles the connection status update event.
func (d *Dota2) handleConnectionStatus(packet *gamecoordinator.GCPacket) error {
	stat := &gcsdkm.CMsgConnectionStatus{}
	if err := d.unmarshalBody(packet, stat); err != nil {
		return err
	}

	if stat.Status == nil {
		return nil
	}

	d.setConnectionStatus(*stat.Status, stat)
	return nil
}

// handleGetProfileCardResponse handles the response to the get profile card request
func (d *Dota2) handleGetProfileCardResponse(packet *gamecoordinator.GCPacket) error {
	card := &gccm.CMsgDOTAProfileCard{}

	if err := d.unmarshalBody(packet, card); err != nil {
		return err
	}
	if card.AccountId == nil {
		return errors.New("received a profile card response with nil account id")
	}

	accountID := *card.AccountId
	d.profileResponseHandlersMtx.Lock()
	handlers := d.profileResponseHandlers[accountID]
	d.profileResponseHandlersMtx.Unlock()

	for _, handler := range handlers {
		select {
		case handler <- card:
		default:
		}
	}

	return nil
}

// HandleGCPacket handles an incoming game coordinator packet.
func (d *Dota2) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppID {
		return
	}

	le := d.le.WithField("msgtype", packet.MsgType)
	handler, ok := d.handlers[packet.MsgType]
	if !ok {
		le.Debug("unhandled gc packet")
		d.emit(&devents.UnhandledGCPacket{
			Packet: packet,
		})
		return
	}

	if err := handler(packet); err != nil {
		le.WithError(err).Warn("error handling gc msg")
	}
}
