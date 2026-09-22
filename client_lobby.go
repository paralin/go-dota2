package dota2

import (
	"context"

	"github.com/paralin/go-dota2/cso"
	gcm "github.com/paralin/go-dota2/protocol"
)

// CreateLobby attempts to create a lobby with details.
func (d *Dota2) CreateLobby(details *gcm.CMsgPracticeLobbySetDetails) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate), &gcm.CMsgPracticeLobbyCreate{
		PassKey:      details.PassKey,
		LobbyDetails: details,
	})
}

// LeaveCreateLobby attempts to leave any current lobby and creates a new one.
func (d *Dota2) LeaveCreateLobby(ctx context.Context, details *gcm.CMsgPracticeLobbySetDetails, destroyOldLobby bool) error {
	// Subscribe before observing the lobby so a transition cannot be missed.
	cacheCtr, err := d.cache.GetContainerForTypeID(uint32(cso.Lobby))
	if err != nil {
		return err
	}
	eventCh, eventCancel, err := cacheCtr.Subscribe()
	if err != nil {
		return err
	}
	defer eventCancel()

	// Leave the current lobby, then wait until a newly created lobby appears.
	var wasInNoLobby bool
	for {
		lobbyObj := cacheCtr.GetOne()
		if lobbyObj != nil {
			lob := lobbyObj.(*gcm.CSODOTALobby)
			le := d.le.WithField("lobby-id", lob.GetLobbyId())
			if wasInNoLobby {
				le.Debug("successfully created lobby")
				return nil
			}

			// Only the current authenticated leader can destroy the old lobby.
			le.Debug("attempting to leave lobby")
			if destroyOldLobby && lob.GetLeaderId() == d.coordinator.SteamID().ToUint64() {
				resp, err := d.DestroyLobby(ctx)
				if err != nil {
					return err
				}
				le.WithField("result", resp.GetResult().String()).Debug("destroy lobby result")
			}
			if lob.GetState() != gcm.CSODOTALobby_UI {
				d.AbandonLobby()
			}
			d.LeaveLobby()
		} else {
			wasInNoLobby = true
			d.le.Debug("creating lobby")
			d.CreateLobby(details)
		}

		// Recheck the authoritative cache after its next update.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-eventCh:
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

// RespondLobbyInvite responds to a lobby invite.
func (d *Dota2) RespondLobbyInvite(lobbyId uint64, accept bool) {
	d.write(uint32(gcm.EGCBaseMsg_k_EMsgGCLobbyInviteResponse), &gcm.CMsgLobbyInviteResponse{
		LobbyId: &lobbyId,
		Accept:  &accept,
	})
}
