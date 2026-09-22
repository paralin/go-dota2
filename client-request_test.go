package dota2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	gcsm "github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-steam/protocol/steamlang"
	"github.com/paralin/go-steam/steamid"
	"github.com/sirupsen/logrus"
)

// requestCoordinator captures actual GC envelopes without a Steam login.
type requestCoordinator struct {
	// writes retains outgoing requests until the test responds.
	writes chan gamecoordinator.IGCMsg
}

// Write captures one outgoing envelope.
func (c *requestCoordinator) Write(message gamecoordinator.IGCMsg) { c.writes <- message }

// SetGamesPlayed accepts the test client's game announcements.
func (c *requestCoordinator) SetGamesPlayed(...uint64) {}

// RegisterPacketHandler leaves packet delivery under the test's control.
func (c *requestCoordinator) RegisterPacketHandler(steam.GCPacketHandler) {}

// SteamID supplies a stable identity for the Dota transport contract.
func (c *requestCoordinator) SteamID() steamid.SteamId { return steamid.SteamId(1) }

// newRequestClient admits a client and bounds each test's request lifetime.
func newRequestClient(t *testing.T) (*Dota2, *requestCoordinator, context.Context) {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	t.Cleanup(cancel)
	transport := &requestCoordinator{writes: make(chan gamecoordinator.IGCMsg, 8)}
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	client := NewWithCoordinator(transport, func(any) {}, logger)
	t.Cleanup(client.Close)
	client.setConnectionStatus(gcsm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil)
	return client, transport, ctx
}

// receiveRequest waits for a real envelope without polling shared state.
func receiveRequest(t *testing.T, ctx context.Context, transport *requestCoordinator) gamecoordinator.IGCMsg {
	t.Helper()
	select {
	case message := <-transport.writes:
		return message
	case <-ctx.Done():
		t.Fatal("request was not sent:", ctx.Err())
		return nil
	}
}

// replyPacket encodes a generated message addressed to one request.
func replyPacket(t *testing.T, message gamecoordinator.IGCMsg, version uint32) *gamecoordinator.GCPacket {
	t.Helper()
	body, err := (&gcsm.CMsgClientWelcome{Version: &version}).MarshalVT()
	if err != nil {
		t.Fatal(err)
	}
	return &gamecoordinator.GCPacket{AppId: AppID, MsgType: 900001, TargetJobId: message.GetSourceJobId(), Body: body}
}

// TestRequestCorrelation exercises simultaneous replies of the same message type.
func TestRequestCorrelation(t *testing.T) {
	// Send two requests, then reverse their replies and include an unsolicited one.
	client, transport, ctx := newRequestClient(t)
	first, second := &gcsm.CMsgClientWelcome{}, &gcsm.CMsgClientWelcome{}
	firstDone, secondDone := make(chan error, 1), make(chan error, 1)
	go func() { firstDone <- client.MakeRequest(ctx, 900000, &gcsm.CMsgClientHello{}, 900001, first) }()
	firstMessage := receiveRequest(t, ctx, transport)
	go func() { secondDone <- client.MakeRequest(ctx, 900000, &gcsm.CMsgClientHello{}, 900001, second) }()
	secondMessage := receiveRequest(t, ctx, transport)
	if firstMessage.GetSourceJobId() == secondMessage.GetSourceJobId() {
		t.Fatal("concurrent requests reused a job ID")
	}
	unsolicited := replyPacket(t, firstMessage, 99)
	unsolicited.TargetJobId = 0
	client.HandleGCPacket(unsolicited)
	client.HandleGCPacket(replyPacket(t, secondMessage, 22))
	client.HandleGCPacket(replyPacket(t, firstMessage, 11))

	// Each caller must receive its own reply, independent of arrival order.
	if err := <-firstDone; err != nil {
		t.Fatal(err)
	}
	if err := <-secondDone; err != nil {
		t.Fatal(err)
	}
	if first.GetVersion() != 11 || second.GetVersion() != 22 {
		t.Fatalf("replies crossed: first=%d second=%d", first.GetVersion(), second.GetVersion())
	}
}

// TestRequestCancellation prevents late replies from mutating a returned result.
func TestRequestCancellation(t *testing.T) {
	// Cancel after the request reaches the transport.
	client, transport, ctx := newRequestClient(t)
	requestCtx, cancel := context.WithCancel(ctx)
	response := &gcsm.CMsgClientWelcome{}
	done := make(chan error, 1)
	go func() { done <- client.MakeRequest(requestCtx, 900000, &gcsm.CMsgClientHello{}, 900001, response) }()
	message := receiveRequest(t, ctx, transport)
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation returned %v", err)
	}

	// Repeated late responses must neither block dispatch nor write to response.
	for range 20 {
		client.HandleGCPacket(replyPacket(t, message, 33))
	}
	if response.GetVersion() != 0 {
		t.Fatal("late response changed the cancelled result")
	}
}

// TestSessionStop interrupts pending work and rejects a late welcome.
func TestSessionStop(t *testing.T) {
	for _, closeClient := range []bool{false, true} {
		t.Run(map[bool]string{false: "stop", true: "close"}[closeClient], func(t *testing.T) {
			// End the session with a request in flight.
			client, transport, ctx := newRequestClient(t)
			done := make(chan error, 1)
			go func() {
				done <- client.MakeRequest(ctx, 900000, &gcsm.CMsgClientHello{}, 900001, &gcsm.CMsgClientWelcome{})
			}()
			receiveRequest(t, ctx, transport)
			if closeClient {
				client.Close()
			} else {
				client.SetPlaying(false)
			}
			if err := <-done; !errors.Is(err, ErrNotReady) {
				t.Fatalf("session interruption returned %v", err)
			}

			// An old welcome cannot reopen admission after stopping or closing.
			client.HandleGCPacket(&gamecoordinator.GCPacket{AppId: AppID, MsgType: uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome)})
			if err := client.MakeRequest(ctx, 900000, &gcsm.CMsgClientHello{}, 900001, &gcsm.CMsgClientWelcome{}); !errors.Is(err, ErrNotReady) {
				t.Fatalf("late welcome reopened admission: %v", err)
			}
		})
	}
}

// TestRequestEnvelope verifies the wire job ID and generated payload together.
func TestRequestEnvelope(t *testing.T) {
	// Serialize through the same Steam envelope implementation used in production.
	message := gamecoordinator.NewGCMsgProtobuf(AppID, 900000, &gcsm.CMsgClientHello{
		Version: proto.Uint32(6689), Engine: gcsm.ESourceEngine_k_ESE_Source2.Enum(),
	})
	message.SetSourceJobId(73)
	var buffer bytes.Buffer
	if err := message.Serialize(&buffer); err != nil {
		t.Fatal(err)
	}

	// Decode the envelope and body through their protocol implementations.
	header := steamlang.NewMsgGCHdrProtoBuf()
	if err := header.Deserialize(&buffer); err != nil {
		t.Fatal(err)
	}
	if header.Proto.GetJobidSource() != 73 {
		t.Fatal("source job ID was lost on the wire")
	}
	var hello gcsm.CMsgClientHello
	if err := hello.UnmarshalVT(buffer.Bytes()); err != nil {
		t.Fatal(err)
	}
	if hello.GetVersion() != 6689 || hello.GetEngine() != gcsm.ESourceEngine_k_ESE_Source2 {
		t.Fatal("generated payload was lost on the wire")
	}
}
