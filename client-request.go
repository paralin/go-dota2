package dota2

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
)

// pendingRequest pairs an expected message type with its receiving goroutine.
type pendingRequest struct {
	// messageType is the response type accepted for this job.
	messageType uint32
	// packets delivers replies without letting the network reader mutate results.
	packets chan *gamecoordinator.GCPacket
}

// MakeRequest sends a correlated GC request and waits within the current session.
// The coordinator must echo the source job ID in the response's target job ID.
// Cancellation releases the result before return; late replies cannot mutate it.
func (d *Dota2) MakeRequest(
	ctx context.Context,
	reqMsgID uint32,
	request proto.Message,
	respMsgID uint32,
	response proto.Message,
	matchesRequest ...func(proto.Message) bool,
) error {
	// Register the request against one ready session before publishing its job ID.
	if err := ctx.Err(); err != nil {
		return err
	}
	d.mtx.Lock()
	if d.closed || d.connectionCtx == nil {
		d.mtx.Unlock()
		return ErrNotReady
	}
	sessionCtx := d.connectionCtx
	d.nextJobID++
	jobID := d.nextJobID
	packets := make(chan *gamecoordinator.GCPacket, 1)
	d.pending[jobID] = pendingRequest{messageType: respMsgID, packets: packets}
	d.mtx.Unlock()
	defer func() {
		d.mtx.Lock()
		delete(d.pending, jobID)
		d.mtx.Unlock()
	}()

	// Send the correlation identity in the GC envelope, not only in local state.
	message := gamecoordinator.NewGCMsgProtobuf(AppID, reqMsgID, request)
	message.SetSourceJobId(jobID)
	d.coordinator.Write(message)

	// Decode only on the caller's goroutine so cancellation cannot race a writer.
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-sessionCtx.Done():
			return ErrNotReady
		case packet := <-packets:
			if err := ctx.Err(); err != nil {
				return err
			}
			if sessionCtx.Err() != nil {
				return ErrNotReady
			}
			if err := d.unmarshalBody(packet, response); err != nil {
				return err
			}
			matches := true
			for _, match := range matchesRequest {
				if !match(response) {
					matches = false
					break
				}
			}
			if matches {
				return nil
			}
		}
	}
}

// handleResponsePacket delivers only a response addressed to a pending job.
func (d *Dota2) handleResponsePacket(packet *gamecoordinator.GCPacket) bool {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	request, ok := d.pending[packet.TargetJobId]
	if !ok || request.messageType != packet.MsgType {
		return false
	}
	select {
	case request.packets <- packet:
	default:
	}
	return true
}
