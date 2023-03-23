// Code generated by protoc-gen-go. DO NOT EDIT.
// source: econ_shared_enums.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EGCEconBaseMsg int32

const (
	EGCEconBaseMsg_k_EMsgGCGenericResult EGCEconBaseMsg = 2579
)

var EGCEconBaseMsg_name = map[int32]string{
	2579: "k_EMsgGCGenericResult",
}

var EGCEconBaseMsg_value = map[string]int32{
	"k_EMsgGCGenericResult": 2579,
}

func (x EGCEconBaseMsg) Enum() *EGCEconBaseMsg {
	p := new(EGCEconBaseMsg)
	*p = x
	return p
}

func (x EGCEconBaseMsg) String() string {
	return proto.EnumName(EGCEconBaseMsg_name, int32(x))
}

func (x *EGCEconBaseMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCEconBaseMsg_value, data, "EGCEconBaseMsg")
	if err != nil {
		return err
	}
	*x = EGCEconBaseMsg(value)
	return nil
}

func (EGCEconBaseMsg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df63e448fc286698, []int{0}
}

type EGCMsgResponse int32

const (
	EGCMsgResponse_k_EGCMsgResponseOK           EGCMsgResponse = 0
	EGCMsgResponse_k_EGCMsgResponseDenied       EGCMsgResponse = 1
	EGCMsgResponse_k_EGCMsgResponseServerError  EGCMsgResponse = 2
	EGCMsgResponse_k_EGCMsgResponseTimeout      EGCMsgResponse = 3
	EGCMsgResponse_k_EGCMsgResponseInvalid      EGCMsgResponse = 4
	EGCMsgResponse_k_EGCMsgResponseNoMatch      EGCMsgResponse = 5
	EGCMsgResponse_k_EGCMsgResponseUnknownError EGCMsgResponse = 6
	EGCMsgResponse_k_EGCMsgResponseNotLoggedOn  EGCMsgResponse = 7
	EGCMsgResponse_k_EGCMsgFailedToCreate       EGCMsgResponse = 8
)

var EGCMsgResponse_name = map[int32]string{
	0: "k_EGCMsgResponseOK",
	1: "k_EGCMsgResponseDenied",
	2: "k_EGCMsgResponseServerError",
	3: "k_EGCMsgResponseTimeout",
	4: "k_EGCMsgResponseInvalid",
	5: "k_EGCMsgResponseNoMatch",
	6: "k_EGCMsgResponseUnknownError",
	7: "k_EGCMsgResponseNotLoggedOn",
	8: "k_EGCMsgFailedToCreate",
}

var EGCMsgResponse_value = map[string]int32{
	"k_EGCMsgResponseOK":           0,
	"k_EGCMsgResponseDenied":       1,
	"k_EGCMsgResponseServerError":  2,
	"k_EGCMsgResponseTimeout":      3,
	"k_EGCMsgResponseInvalid":      4,
	"k_EGCMsgResponseNoMatch":      5,
	"k_EGCMsgResponseUnknownError": 6,
	"k_EGCMsgResponseNotLoggedOn":  7,
	"k_EGCMsgFailedToCreate":       8,
}

func (x EGCMsgResponse) Enum() *EGCMsgResponse {
	p := new(EGCMsgResponse)
	*p = x
	return p
}

func (x EGCMsgResponse) String() string {
	return proto.EnumName(EGCMsgResponse_name, int32(x))
}

func (x *EGCMsgResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgResponse_value, data, "EGCMsgResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgResponse(value)
	return nil
}

func (EGCMsgResponse) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df63e448fc286698, []int{1}
}

type EGCMsgUseItemResponse int32

const (
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed                    EGCMsgUseItemResponse = 0
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_GiftNoOtherPlayers          EGCMsgUseItemResponse = 1
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ServerError                 EGCMsgUseItemResponse = 2
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MiniGameAlreadyStarted      EGCMsgUseItemResponse = 3
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted       EGCMsgUseItemResponse = 4
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted EGCMsgUseItemResponse = 5
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotInLowPriorityPool        EGCMsgUseItemResponse = 6
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotHighEnoughLevel          EGCMsgUseItemResponse = 7
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EventNotActive              EGCMsgUseItemResponse = 8
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted EGCMsgUseItemResponse = 9
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MissingRequirement          EGCMsgUseItemResponse = 10
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew        EGCMsgUseItemResponse = 11
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_Complete     EGCMsgUseItemResponse = 12
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_Compendium         EGCMsgUseItemResponse = 13
)

var EGCMsgUseItemResponse_name = map[int32]string{
	0:  "k_EGCMsgUseItemResponse_ItemUsed",
	1:  "k_EGCMsgUseItemResponse_GiftNoOtherPlayers",
	2:  "k_EGCMsgUseItemResponse_ServerError",
	3:  "k_EGCMsgUseItemResponse_MiniGameAlreadyStarted",
	4:  "k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted",
	5:  "k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted",
	6:  "k_EGCMsgUseItemResponse_NotInLowPriorityPool",
	7:  "k_EGCMsgUseItemResponse_NotHighEnoughLevel",
	8:  "k_EGCMsgUseItemResponse_EventNotActive",
	9:  "k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted",
	10: "k_EGCMsgUseItemResponse_MissingRequirement",
	11: "k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew",
	12: "k_EGCMsgUseItemResponse_EmoticonUnlock_Complete",
	13: "k_EGCMsgUseItemResponse_ItemUsed_Compendium",
}

var EGCMsgUseItemResponse_value = map[string]int32{
	"k_EGCMsgUseItemResponse_ItemUsed":                    0,
	"k_EGCMsgUseItemResponse_GiftNoOtherPlayers":          1,
	"k_EGCMsgUseItemResponse_ServerError":                 2,
	"k_EGCMsgUseItemResponse_MiniGameAlreadyStarted":      3,
	"k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted":       4,
	"k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted": 5,
	"k_EGCMsgUseItemResponse_NotInLowPriorityPool":        6,
	"k_EGCMsgUseItemResponse_NotHighEnoughLevel":          7,
	"k_EGCMsgUseItemResponse_EventNotActive":              8,
	"k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted": 9,
	"k_EGCMsgUseItemResponse_MissingRequirement":          10,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew":        11,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_Complete":     12,
	"k_EGCMsgUseItemResponse_ItemUsed_Compendium":         13,
}

func (x EGCMsgUseItemResponse) Enum() *EGCMsgUseItemResponse {
	p := new(EGCMsgUseItemResponse)
	*p = x
	return p
}

func (x EGCMsgUseItemResponse) String() string {
	return proto.EnumName(EGCMsgUseItemResponse_name, int32(x))
}

func (x *EGCMsgUseItemResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgUseItemResponse_value, data, "EGCMsgUseItemResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgUseItemResponse(value)
	return nil
}

func (EGCMsgUseItemResponse) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df63e448fc286698, []int{2}
}

type CMsgGenericResult struct {
	Eresult              *uint32  `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	DebugMessage         *string  `protobuf:"bytes,2,opt,name=debug_message,json=debugMessage" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgGenericResult) Reset()         { *m = CMsgGenericResult{} }
func (m *CMsgGenericResult) String() string { return proto.CompactTextString(m) }
func (*CMsgGenericResult) ProtoMessage()    {}
func (*CMsgGenericResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_df63e448fc286698, []int{0}
}

func (m *CMsgGenericResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgGenericResult.Unmarshal(m, b)
}
func (m *CMsgGenericResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgGenericResult.Marshal(b, m, deterministic)
}
func (m *CMsgGenericResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgGenericResult.Merge(m, src)
}
func (m *CMsgGenericResult) XXX_Size() int {
	return xxx_messageInfo_CMsgGenericResult.Size(m)
}
func (m *CMsgGenericResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgGenericResult.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgGenericResult proto.InternalMessageInfo

const Default_CMsgGenericResult_Eresult uint32 = 2

func (m *CMsgGenericResult) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgGenericResult_Eresult
}

func (m *CMsgGenericResult) GetDebugMessage() string {
	if m != nil && m.DebugMessage != nil {
		return *m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.EGCEconBaseMsg", EGCEconBaseMsg_name, EGCEconBaseMsg_value)
	proto.RegisterEnum("protocol.EGCMsgResponse", EGCMsgResponse_name, EGCMsgResponse_value)
	proto.RegisterEnum("protocol.EGCMsgUseItemResponse", EGCMsgUseItemResponse_name, EGCMsgUseItemResponse_value)
	proto.RegisterType((*CMsgGenericResult)(nil), "protocol.CMsgGenericResult")
}

func init() {
	proto.RegisterFile("econ_shared_enums.proto", fileDescriptor_df63e448fc286698)
}

var fileDescriptor_df63e448fc286698 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x5d, 0x6f, 0x12, 0x4d,
	0x14, 0xee, 0xb6, 0xa5, 0xd0, 0x79, 0xcb, 0x9b, 0x71, 0x62, 0x3f, 0x52, 0x4c, 0x24, 0xd6, 0x68,
	0x83, 0x95, 0x6a, 0x7b, 0x61, 0xe2, 0x5d, 0x4b, 0x57, 0x24, 0xb2, 0x0b, 0xd9, 0x96, 0xeb, 0xcd,
	0xb8, 0x7b, 0x5c, 0x26, 0xec, 0xce, 0xc1, 0x99, 0x59, 0x08, 0xbf, 0xc3, 0x4b, 0x7f, 0x8f, 0xff,
	0xcb, 0xb0, 0x08, 0x11, 0xc2, 0x52, 0xaf, 0xe6, 0xcc, 0xf3, 0x91, 0xe7, 0x3c, 0x87, 0x1c, 0x43,
	0x80, 0xd2, 0xd7, 0x7d, 0xae, 0x20, 0xf4, 0x41, 0xa6, 0x89, 0xae, 0x0f, 0x15, 0x1a, 0x64, 0xa5,
	0xec, 0x09, 0x30, 0x7e, 0xd1, 0x23, 0x4f, 0x1a, 0x8e, 0x8e, 0x9a, 0x20, 0x41, 0x89, 0xc0, 0x03,
	0x9d, 0xc6, 0x86, 0x55, 0x48, 0x11, 0x54, 0x36, 0x9e, 0x58, 0x55, 0xeb, 0xbc, 0xfc, 0xd1, 0xba,
	0xf2, 0xe6, 0x08, 0x3b, 0x23, 0xe5, 0x10, 0xbe, 0xa6, 0x91, 0x9f, 0x80, 0xd6, 0x3c, 0x82, 0x93,
	0xed, 0xaa, 0x75, 0xbe, 0xef, 0x1d, 0x64, 0xa0, 0x33, 0xc3, 0x6a, 0x17, 0xe4, 0x7f, 0xbb, 0xd9,
	0xb0, 0x03, 0x94, 0xb7, 0x5c, 0x83, 0xa3, 0x23, 0x76, 0x4a, 0x0e, 0x07, 0xbe, 0x3d, 0x8d, 0x6a,
	0x2c, 0x85, 0xd1, 0x1f, 0x4f, 0x6b, 0x3f, 0xb7, 0x33, 0xb9, 0xa3, 0x23, 0x0f, 0xf4, 0x10, 0xa5,
	0x06, 0x76, 0x44, 0xd8, 0xc0, 0x5f, 0xc6, 0x3a, 0x5f, 0xe8, 0x16, 0x3b, 0x25, 0x47, 0xab, 0xf8,
	0x1d, 0x48, 0x01, 0x21, 0xb5, 0xd8, 0x73, 0x52, 0x59, 0xe5, 0xee, 0x41, 0x8d, 0x40, 0xd9, 0x4a,
	0xa1, 0xa2, 0xdb, 0xac, 0x42, 0x8e, 0x57, 0x05, 0x0f, 0x22, 0x01, 0x4c, 0x0d, 0xdd, 0x59, 0x47,
	0xb6, 0xe4, 0x88, 0xc7, 0x22, 0xa4, 0xbb, 0xeb, 0x48, 0x17, 0x1d, 0x6e, 0x82, 0x3e, 0x2d, 0xb0,
	0x2a, 0x79, 0xb6, 0x4a, 0xf6, 0xe4, 0x40, 0xe2, 0x58, 0xce, 0x82, 0xf7, 0xd6, 0x6d, 0xe6, 0xa2,
	0x69, 0x63, 0x14, 0x41, 0xd8, 0x91, 0xb4, 0xf8, 0x77, 0xad, 0x4f, 0x5c, 0xc4, 0x10, 0x3e, 0x60,
	0x43, 0x01, 0x37, 0x40, 0x4b, 0xb5, 0x5f, 0x05, 0x72, 0x38, 0xa3, 0x7a, 0x1a, 0x5a, 0x06, 0x92,
	0xc5, 0x91, 0x5e, 0x92, 0xea, 0xdc, 0xb5, 0x42, 0xf9, 0xd3, 0x4f, 0x4f, 0x43, 0x48, 0xb7, 0x58,
	0x9d, 0xd4, 0xf2, 0x54, 0x4d, 0xf1, 0xcd, 0xb8, 0xd8, 0x31, 0x7d, 0x50, 0xdd, 0x98, 0x4f, 0x40,
	0x69, 0x6a, 0xb1, 0xd7, 0xe4, 0x2c, 0x4f, 0xbf, 0x7c, 0xce, 0x2b, 0x52, 0xcf, 0x13, 0x3a, 0x42,
	0x8a, 0x26, 0x4f, 0xe0, 0x26, 0x56, 0xc0, 0xc3, 0xc9, 0xbd, 0xe1, 0xca, 0x40, 0x48, 0x77, 0xd8,
	0x7b, 0xf2, 0xf6, 0xb1, 0x95, 0xb3, 0x41, 0x37, 0x15, 0x97, 0x53, 0xcb, 0x2e, 0xfb, 0x40, 0xae,
	0xf3, 0x2c, 0x77, 0x0a, 0x87, 0x1e, 0x37, 0x70, 0x8b, 0x32, 0xd5, 0x7f, 0xb2, 0xe6, 0xc6, 0x02,
	0x7b, 0x47, 0x2e, 0xf2, 0x8c, 0x2e, 0x9a, 0x96, 0x6c, 0xe3, 0xb8, 0xab, 0x04, 0x2a, 0x61, 0x26,
	0x5d, 0xc4, 0x98, 0xee, 0x6d, 0x3a, 0x95, 0x8b, 0xe6, 0xb3, 0x88, 0xfa, 0xb6, 0xc4, 0x34, 0xea,
	0xb7, 0x61, 0x04, 0x31, 0x2d, 0xb2, 0x1a, 0x79, 0x95, 0xa7, 0xb7, 0x47, 0x20, 0x8d, 0x8b, 0xe6,
	0x26, 0x30, 0x62, 0x04, 0xb4, 0xb4, 0xa9, 0xc6, 0xa2, 0x79, 0x66, 0xea, 0xa2, 0x90, 0x66, 0xd1,
	0x7f, 0x7f, 0xd3, 0x52, 0x8e, 0xd0, 0x5a, 0xc8, 0xc8, 0x83, 0xef, 0xa9, 0x50, 0x90, 0x80, 0x34,
	0x94, 0x6c, 0xaa, 0x6d, 0x27, 0x68, 0x44, 0x80, 0xb2, 0x27, 0x63, 0x0c, 0x06, 0xbe, 0x8b, 0x2e,
	0x8c, 0xe9, 0x7f, 0xec, 0x9a, 0x5c, 0xfe, 0xa3, 0xa3, 0x81, 0xc9, 0x30, 0x06, 0x03, 0xf4, 0x80,
	0x5d, 0x92, 0x37, 0x8f, 0xf6, 0x99, 0xca, 0x41, 0x86, 0x22, 0x4d, 0x68, 0xf9, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0xfa, 0x3c, 0xee, 0x8e, 0x04, 0x00, 0x00,
}
