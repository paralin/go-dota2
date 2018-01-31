package dota2

import (
	"context"

	"github.com/golang/protobuf/proto"
	// gcmm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	gcclm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	gccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// RequestProfileCard sends a request to the DOTA gc for a card and waits for the response.
// If you want to enable timeouts, use a context.WithTimeout
func (d *Dota2) RequestProfileCard(ctx context.Context, accountID uint32) (*gccm.CMsgDOTAProfileCard, error) {
	resp := &gccm.CMsgDOTAProfileCard{}
	req := &gcclm.CMsgClientToGCGetProfileCard{AccountId: &accountID}

	return resp, d.MakeRequest(
		ctx,
		uint32(gcm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard),
		req,
		uint32(gcm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse),
		resp,
		func(m proto.Message) bool {
			potentialResp := m.(*gccm.CMsgDOTAProfileCard)
			return potentialResp.GetAccountId() == accountID
		},
	)
}
