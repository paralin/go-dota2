package dota2

import (
	"context"

	"github.com/paralin/go-dota2/cso"
	gcccm "github.com/paralin/go-dota2/protocol"
	gcm "github.com/paralin/go-dota2/protocol"
	gcmm "github.com/paralin/go-dota2/protocol"
)

// CreateLobby attempts to create a lobby with details.
func (d *Dota2) CreateLobby(details *gcccm.CMsgPracticeLobbySetDetails) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate), &gcccm.CMsgPracticeLobbyCreate{
		PassKey:      details.PassKey,
		LobbyDetails: details,
	})
}

// LeaveCreateLobby attempts to leave any current lobby and creates a new one.
func (d *Dota2) LeaveCreateLobby(ctx context.Context, details *gcccm.CMsgPracticeLobbySetDetails, destroyOldLobby bool) error {
	cacheCtr, err := d.cache.GetContainerForTypeID(uint32(cso.Lobby))
	if err != nil {
		return err
	}

	eventCh, eventCancel, err := cacheCtr.Subscribe()
	if err != nil {
		return err
	}
	defer eventCancel()

	var wasInNoLobby bool
	for {
		lobbyObj := cacheCtr.GetOne()
		if lobbyObj != nil {
			lob := lobbyObj.(*gcmm.CSODOTALobby)
			le := d.le.WithField("lobby-id", lob.GetLobbyId())
			if wasInNoLobby {
				le.Debug("successfully created lobby")
				return nil
			}

			le.Debug("attempting to leave lobby")
			if destroyOldLobby && lob.GetLeaderId() == d.client.SteamId().ToUint64() {
				resp, err := d.DestroyLobby(ctx)
				if err != nil {
					return err
				}
				le.WithField("result", resp.GetResult().String()).Debug("destroy lobby result")
			}
			if lob.GetState() != gcmm.CSODOTALobby_UI {
				d.AbandonLobby()
			}
			d.LeaveLobby()
		} else {
			wasInNoLobby = true
			d.le.Debug("creating lobby")
			d.CreateLobby(details)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-eventCh:
			_ = event
		}
	}
}

// ClearTeamFromLobby clears the team from a practice lobby.
func (d *Dota2) ClearTeamFromLobby() {
	d.write(
		uint32(gcm.EDOTAGCMsg_k_EMsgGCClearPracticeLobbyTeam),
		// unknown proto type: send empty of this one
		&gcm.CMsgFlipLobbyTeams{},
	)
}
