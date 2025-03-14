# Documentation for go-dota2
### Updated: 14.03.25
Note: Since the library (in theory) has all of the dota 2 GC functions integrated (thanks to [apigen](https://github.com/voltageeee/go-dota2/tree/master/apigen)), this documentation isn't complete and never will be complete. As of now, it provides basic information on how to update generated code and how to generate protobufs.
Most of the client-to-gc functions can be found in [client_generated.go](https://github.com/voltageeee/go-dota2/blob/master/client_generated.go) file.

### Updating protobufs
Protobuf (protocol buffers) is a protocol for exchanging serialized and structured data types. Let's take a look on a CSODOTALobby object in protobuf format:

    message CSODOTALobby {
	message CExtraMsg {
		optional uint32 id = 1;
		optional bytes contents = 2;
	}

	enum State {
		UI = 0;
		READYUP = 4;
		// lots of other states...
	}

	enum LobbyType {
		INVALID = -1;
		CASUAL_MATCH = 0;
		PRACTICE = 1;
		// lots of other types...
	}

	optional uint64 lobby_id = 1 [(key_field) = true];
	repeated CSODOTALobbyMember all_members = 120;
	repeated uint32 member_indices = 121;
	repeated uint32 left_member_indices = 122;
	repeated uint32 free_member_indices = 123;
	optional fixed64 leader_id = 11;
	optional fixed64 server_id = 6 [default = 0];
	optional uint32 game_mode = 3;
	// lots of other fields