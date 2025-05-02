package dota2

import (
	"context"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-steam/protocol/gamecoordinator"

	gcm "github.com/paralin/go-dota2/protocol"
)

// responseHandler returns handled, and any error
type responseHandler func(resp *gamecoordinator.GCPacket) bool

// MakeRequest starts and tracks a request given a context.
func (d *Dota2) MakeRequest(
	ctx context.Context,
	reqMsgID uint32,
	request proto.Message,
	respMsgID uint32,
	response proto.Message,
	matchesRequest ...func(proto.Message) bool,
) (mrErr error) {
	d.le.Debugf("making request: %s", gcm.EDOTAGCMsg(reqMsgID).String())
	cctx, err := d.validateConnectionContext()
	if err != nil {
		return err
	}

	var respMtx sync.Mutex
	respCh := make(chan error, 10)

	d.pendReqMtx.Lock()
	reqID := d.pendReqID
	d.pendReqID++

	defer func() {
		if mrErr == nil {
			return
		}

		d.pendReqMtx.Lock()
		m := d.pendReq[respMsgID]
		if m != nil {
			delete(m, reqID)
		}
		if len(m) == 0 {
			delete(d.pendReq, respMsgID)
		}
		d.pendReqMtx.Unlock()
	}()

	m := d.pendReq[respMsgID]
	if m == nil {
		m = make(map[uint32]responseHandler)
		d.pendReq[respMsgID] = m
	}

	m[reqID] = func(resp *gamecoordinator.GCPacket) bool {
		respMtx.Lock()
		defer respMtx.Unlock()

		if err := d.unmarshalBody(resp, response); err != nil {
			respCh <- err
			return false
		}

		for _, handler := range matchesRequest {
			if !handler(response) {
				return false
			}
		}

		respCh <- nil
		return true
	}
	d.pendReqMtx.Unlock()

	d.write(reqMsgID, request)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-cctx.Done():
		return ErrNotReady
	case rErr := <-respCh:
		return rErr
	}
}

// handleResponsePacket attempts to handle a response packet.
func (d *Dota2) handleResponsePacket(packet *gamecoordinator.GCPacket) (handled bool) {
	d.pendReqMtx.Lock()
	defer d.pendReqMtx.Unlock()

	respHandlers := d.pendReq[packet.MsgType]
	for rid, h := range respHandlers {
		if hnd := h(packet); hnd {
			handled = true
			delete(respHandlers, rid)
		}
	}

	return
}
