package dota2

import (
	"context"

	"github.com/paralin/go-dota2/cso"
	gccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	gcmm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// JoinLobby attempts to join a lobby by ID.
func (d *Dota2) JoinLobby(lobbyID uint64, passKey string) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoin), &gcccm.CMsgPracticeLobbyJoin{
		LobbyId: &lobbyID,
		PassKey: &passKey,
	})
}

// CreateLobby attempts to create a lobby with details.
func (d *Dota2) CreateLobby(details *gcccm.CMsgPracticeLobbySetDetails) {
	// TODO: investigate SearchKey
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate), &gcccm.CMsgPracticeLobbyCreate{
		PassKey:      details.PassKey,
		LobbyDetails: details,
	})
}

// LaunchLobby launches the current lobby.
func (d *Dota2) LaunchLobby() {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyLaunch), &gcccm.CMsgPracticeLobbyLaunch{})
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
				d.DestroyLobby()
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

// LeaveLobby attempts to leave the current lobby.
func (d *Dota2) LeaveLobby() {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyLeave), &gcccm.CMsgPracticeLobbyLeave{})
}

// AbandonLobby abandons the current lobby.
func (d *Dota2) AbandonLobby() {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame), &gcccm.CMsgAbandonCurrentGame{})
}

// DestroyLobby attempts to destroy the lobby.
func (d *Dota2) DestroyLobby() {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgDestroyLobbyRequest), &gccm.CMsgDOTADestroyLobbyRequest{})
}
