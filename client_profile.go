package dota2

import (
	"context"
	"errors"

	"github.com/faceit/go-steam/protocol/gamecoordinator"
	// gcmm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	gcclm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	gccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

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
