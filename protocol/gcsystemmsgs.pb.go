// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: gcsystemmsgs.proto

package protocol

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ESOMsg int32

const (
	ESOMsg_k_ESOMsg_Create                   ESOMsg = 21
	ESOMsg_k_ESOMsg_Update                   ESOMsg = 22
	ESOMsg_k_ESOMsg_Destroy                  ESOMsg = 23
	ESOMsg_k_ESOMsg_CacheSubscribed          ESOMsg = 24
	ESOMsg_k_ESOMsg_CacheUnsubscribed        ESOMsg = 25
	ESOMsg_k_ESOMsg_UpdateMultiple           ESOMsg = 26
	ESOMsg_k_ESOMsg_CacheSubscriptionRefresh ESOMsg = 28
	ESOMsg_k_ESOMsg_CacheSubscribedUpToDate  ESOMsg = 29
)

// Enum value maps for ESOMsg.
var (
	ESOMsg_name = map[int32]string{
		21: "k_ESOMsg_Create",
		22: "k_ESOMsg_Update",
		23: "k_ESOMsg_Destroy",
		24: "k_ESOMsg_CacheSubscribed",
		25: "k_ESOMsg_CacheUnsubscribed",
		26: "k_ESOMsg_UpdateMultiple",
		28: "k_ESOMsg_CacheSubscriptionRefresh",
		29: "k_ESOMsg_CacheSubscribedUpToDate",
	}
	ESOMsg_value = map[string]int32{
		"k_ESOMsg_Create":                   21,
		"k_ESOMsg_Update":                   22,
		"k_ESOMsg_Destroy":                  23,
		"k_ESOMsg_CacheSubscribed":          24,
		"k_ESOMsg_CacheUnsubscribed":        25,
		"k_ESOMsg_UpdateMultiple":           26,
		"k_ESOMsg_CacheSubscriptionRefresh": 28,
		"k_ESOMsg_CacheSubscribedUpToDate":  29,
	}
)

func (x ESOMsg) Enum() *ESOMsg {
	p := new(ESOMsg)
	*p = x
	return p
}

func (x ESOMsg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ESOMsg) Descriptor() protoreflect.EnumDescriptor {
	return file_gcsystemmsgs_proto_enumTypes[0].Descriptor()
}

func (ESOMsg) Type() protoreflect.EnumType {
	return &file_gcsystemmsgs_proto_enumTypes[0]
}

func (x ESOMsg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ESOMsg) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ESOMsg(num)
	return nil
}

// Deprecated: Use ESOMsg.Descriptor instead.
func (ESOMsg) EnumDescriptor() ([]byte, []int) {
	return file_gcsystemmsgs_proto_rawDescGZIP(), []int{0}
}

type EGCBaseClientMsg int32

const (
	EGCBaseClientMsg_k_EMsgGCPingRequest                  EGCBaseClientMsg = 3001
	EGCBaseClientMsg_k_EMsgGCPingResponse                 EGCBaseClientMsg = 3002
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarRequest    EGCBaseClientMsg = 3003
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarResponse   EGCBaseClientMsg = 3004
	EGCBaseClientMsg_k_EMsgGCCompressedMsgToClient        EGCBaseClientMsg = 3005
	EGCBaseClientMsg_k_EMsgGCCompressedMsgToClient_Legacy EGCBaseClientMsg = 523
	EGCBaseClientMsg_k_EMsgGCToClientRequestDropped       EGCBaseClientMsg = 3006
	EGCBaseClientMsg_k_EMsgGCClientWelcome                EGCBaseClientMsg = 4004
	EGCBaseClientMsg_k_EMsgGCServerWelcome                EGCBaseClientMsg = 4005
	EGCBaseClientMsg_k_EMsgGCClientHello                  EGCBaseClientMsg = 4006
	EGCBaseClientMsg_k_EMsgGCServerHello                  EGCBaseClientMsg = 4007
	EGCBaseClientMsg_k_EMsgGCClientConnectionStatus       EGCBaseClientMsg = 4009
	EGCBaseClientMsg_k_EMsgGCServerConnectionStatus       EGCBaseClientMsg = 4010
)

// Enum value maps for EGCBaseClientMsg.
var (
	EGCBaseClientMsg_name = map[int32]string{
		3001: "k_EMsgGCPingRequest",
		3002: "k_EMsgGCPingResponse",
		3003: "k_EMsgGCToClientPollConvarRequest",
		3004: "k_EMsgGCToClientPollConvarResponse",
		3005: "k_EMsgGCCompressedMsgToClient",
		523:  "k_EMsgGCCompressedMsgToClient_Legacy",
		3006: "k_EMsgGCToClientRequestDropped",
		4004: "k_EMsgGCClientWelcome",
		4005: "k_EMsgGCServerWelcome",
		4006: "k_EMsgGCClientHello",
		4007: "k_EMsgGCServerHello",
		4009: "k_EMsgGCClientConnectionStatus",
		4010: "k_EMsgGCServerConnectionStatus",
	}
	EGCBaseClientMsg_value = map[string]int32{
		"k_EMsgGCPingRequest":                  3001,
		"k_EMsgGCPingResponse":                 3002,
		"k_EMsgGCToClientPollConvarRequest":    3003,
		"k_EMsgGCToClientPollConvarResponse":   3004,
		"k_EMsgGCCompressedMsgToClient":        3005,
		"k_EMsgGCCompressedMsgToClient_Legacy": 523,
		"k_EMsgGCToClientRequestDropped":       3006,
		"k_EMsgGCClientWelcome":                4004,
		"k_EMsgGCServerWelcome":                4005,
		"k_EMsgGCClientHello":                  4006,
		"k_EMsgGCServerHello":                  4007,
		"k_EMsgGCClientConnectionStatus":       4009,
		"k_EMsgGCServerConnectionStatus":       4010,
	}
)

func (x EGCBaseClientMsg) Enum() *EGCBaseClientMsg {
	p := new(EGCBaseClientMsg)
	*p = x
	return p
}

func (x EGCBaseClientMsg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EGCBaseClientMsg) Descriptor() protoreflect.EnumDescriptor {
	return file_gcsystemmsgs_proto_enumTypes[1].Descriptor()
}

func (EGCBaseClientMsg) Type() protoreflect.EnumType {
	return &file_gcsystemmsgs_proto_enumTypes[1]
}

func (x EGCBaseClientMsg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EGCBaseClientMsg) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EGCBaseClientMsg(num)
	return nil
}

// Deprecated: Use EGCBaseClientMsg.Descriptor instead.
func (EGCBaseClientMsg) EnumDescriptor() ([]byte, []int) {
	return file_gcsystemmsgs_proto_rawDescGZIP(), []int{1}
}

var File_gcsystemmsgs_proto protoreflect.FileDescriptor

const file_gcsystemmsgs_proto_rawDesc = "" +
	"\n" +
	"\x12gcsystemmsgs.proto\x12\bprotocol*\xf0\x01\n" +
	"\x06ESOMsg\x12\x13\n" +
	"\x0fk_ESOMsg_Create\x10\x15\x12\x13\n" +
	"\x0fk_ESOMsg_Update\x10\x16\x12\x14\n" +
	"\x10k_ESOMsg_Destroy\x10\x17\x12\x1c\n" +
	"\x18k_ESOMsg_CacheSubscribed\x10\x18\x12\x1e\n" +
	"\x1ak_ESOMsg_CacheUnsubscribed\x10\x19\x12\x1b\n" +
	"\x17k_ESOMsg_UpdateMultiple\x10\x1a\x12%\n" +
	"!k_ESOMsg_CacheSubscriptionRefresh\x10\x1c\x12$\n" +
	" k_ESOMsg_CacheSubscribedUpToDate\x10\x1d*\xc2\x03\n" +
	"\x10EGCBaseClientMsg\x12\x18\n" +
	"\x13k_EMsgGCPingRequest\x10\xb9\x17\x12\x19\n" +
	"\x14k_EMsgGCPingResponse\x10\xba\x17\x12&\n" +
	"!k_EMsgGCToClientPollConvarRequest\x10\xbb\x17\x12'\n" +
	"\"k_EMsgGCToClientPollConvarResponse\x10\xbc\x17\x12\"\n" +
	"\x1dk_EMsgGCCompressedMsgToClient\x10\xbd\x17\x12)\n" +
	"$k_EMsgGCCompressedMsgToClient_Legacy\x10\x8b\x04\x12#\n" +
	"\x1ek_EMsgGCToClientRequestDropped\x10\xbe\x17\x12\x1a\n" +
	"\x15k_EMsgGCClientWelcome\x10\xa4\x1f\x12\x1a\n" +
	"\x15k_EMsgGCServerWelcome\x10\xa5\x1f\x12\x18\n" +
	"\x13k_EMsgGCClientHello\x10\xa6\x1f\x12\x18\n" +
	"\x13k_EMsgGCServerHello\x10\xa7\x1f\x12#\n" +
	"\x1ek_EMsgGCClientConnectionStatus\x10\xa9\x1f\x12#\n" +
	"\x1ek_EMsgGCServerConnectionStatus\x10\xaa\x1f"

var (
	file_gcsystemmsgs_proto_rawDescOnce sync.Once
	file_gcsystemmsgs_proto_rawDescData []byte
)

func file_gcsystemmsgs_proto_rawDescGZIP() []byte {
	file_gcsystemmsgs_proto_rawDescOnce.Do(func() {
		file_gcsystemmsgs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gcsystemmsgs_proto_rawDesc), len(file_gcsystemmsgs_proto_rawDesc)))
	})
	return file_gcsystemmsgs_proto_rawDescData
}

var file_gcsystemmsgs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gcsystemmsgs_proto_goTypes = []any{
	(ESOMsg)(0),           // 0: protocol.ESOMsg
	(EGCBaseClientMsg)(0), // 1: protocol.EGCBaseClientMsg
}
var file_gcsystemmsgs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gcsystemmsgs_proto_init() }
func file_gcsystemmsgs_proto_init() {
	if File_gcsystemmsgs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gcsystemmsgs_proto_rawDesc), len(file_gcsystemmsgs_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gcsystemmsgs_proto_goTypes,
		DependencyIndexes: file_gcsystemmsgs_proto_depIdxs,
		EnumInfos:         file_gcsystemmsgs_proto_enumTypes,
	}.Build()
	File_gcsystemmsgs_proto = out.File
	file_gcsystemmsgs_proto_goTypes = nil
	file_gcsystemmsgs_proto_depIdxs = nil
}
