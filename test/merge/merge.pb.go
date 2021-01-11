// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merge.proto

package merge

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyMessageNonNull_Color int32

const (
	MyMessageNonNull_RED   MyMessageNonNull_Color = 0
	MyMessageNonNull_GREEN MyMessageNonNull_Color = 1
	MyMessageNonNull_BLUE  MyMessageNonNull_Color = 2
)

var MyMessageNonNull_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var MyMessageNonNull_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x MyMessageNonNull_Color) String() string {
	return proto.EnumName(MyMessageNonNull_Color_name, int32(x))
}
func (MyMessageNonNull_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_merge_9721c1510159f9a5, []int{3, 0}
}

type InnerMessage struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Connected            bool     `protobuf:"varint,3,opt,name=connected,proto3" json:"connected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_merge_9721c1510159f9a5, []int{0}
}
func (m *InnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMessage.Unmarshal(m, b)
}
func (m *InnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMessage.Marshal(b, m, deterministic)
}
func (dst *InnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMessage.Merge(dst, src)
}
func (m *InnerMessage) XXX_Size() int {
	return xxx_messageInfo_InnerMessage.Size(m)
}
func (m *InnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMessage proto.InternalMessageInfo

func (m *InnerMessage) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InnerMessage) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InnerMessage) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

type OtherMessage struct {
	Key                  int64         `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Weight               float32       `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Inner                *InnerMessage `protobuf:"bytes,4,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OtherMessage) Reset()         { *m = OtherMessage{} }
func (m *OtherMessage) String() string { return proto.CompactTextString(m) }
func (*OtherMessage) ProtoMessage()    {}
func (*OtherMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_merge_9721c1510159f9a5, []int{1}
}
func (m *OtherMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherMessage.Unmarshal(m, b)
}
func (m *OtherMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherMessage.Marshal(b, m, deterministic)
}
func (dst *OtherMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherMessage.Merge(dst, src)
}
func (m *OtherMessage) XXX_Size() int {
	return xxx_messageInfo_OtherMessage.Size(m)
}
func (m *OtherMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OtherMessage proto.InternalMessageInfo

func (m *OtherMessage) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *OtherMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OtherMessage) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *OtherMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

type RequiredInnerMessageNonNull struct {
	LeoFinallyWonAnOscar InnerMessage `protobuf:"bytes,1,opt,name=leo_finally_won_an_oscar,json=leoFinallyWonAnOscar,proto3" json:"leo_finally_won_an_oscar"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RequiredInnerMessageNonNull) Reset()         { *m = RequiredInnerMessageNonNull{} }
func (m *RequiredInnerMessageNonNull) String() string { return proto.CompactTextString(m) }
func (*RequiredInnerMessageNonNull) ProtoMessage()    {}
func (*RequiredInnerMessageNonNull) Descriptor() ([]byte, []int) {
	return fileDescriptor_merge_9721c1510159f9a5, []int{2}
}
func (m *RequiredInnerMessageNonNull) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequiredInnerMessageNonNull.Unmarshal(m, b)
}
func (m *RequiredInnerMessageNonNull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequiredInnerMessageNonNull.Marshal(b, m, deterministic)
}
func (dst *RequiredInnerMessageNonNull) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredInnerMessageNonNull.Merge(dst, src)
}
func (m *RequiredInnerMessageNonNull) XXX_Size() int {
	return xxx_messageInfo_RequiredInnerMessageNonNull.Size(m)
}
func (m *RequiredInnerMessageNonNull) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredInnerMessageNonNull.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredInnerMessageNonNull proto.InternalMessageInfo

func (m *RequiredInnerMessageNonNull) GetLeoFinallyWonAnOscar() InnerMessage {
	if m != nil {
		return m.LeoFinallyWonAnOscar
	}
	return InnerMessage{}
}

type MyMessageNonNull struct {
	Count          int32                       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Name           string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quote          string                      `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	Pet            []string                    `protobuf:"bytes,4,rep,name=pet,proto3" json:"pet,omitempty"`
	Inner          InnerMessage                `protobuf:"bytes,5,opt,name=inner,proto3" json:"inner"`
	Others         []OtherMessage              `protobuf:"bytes,6,rep,name=others,proto3" json:"others"`
	WeMustGoDeeper RequiredInnerMessageNonNull `protobuf:"bytes,13,opt,name=we_must_go_deeper,json=weMustGoDeeper,proto3" json:"we_must_go_deeper"`
	RepInner       []InnerMessage              `protobuf:"bytes,12,rep,name=rep_inner,json=repInner,proto3" json:"rep_inner"`
	Bikeshed       MyMessageNonNull_Color      `protobuf:"varint,7,opt,name=bikeshed,proto3,enum=merge.MyMessageNonNull_Color" json:"bikeshed,omitempty"`
	// This field becomes [][]byte in the generated code.
	RepBytes             [][]byte `protobuf:"bytes,10,rep,name=rep_bytes,json=repBytes,proto3" json:"rep_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMessageNonNull) Reset()         { *m = MyMessageNonNull{} }
func (m *MyMessageNonNull) String() string { return proto.CompactTextString(m) }
func (*MyMessageNonNull) ProtoMessage()    {}
func (*MyMessageNonNull) Descriptor() ([]byte, []int) {
	return fileDescriptor_merge_9721c1510159f9a5, []int{3}
}
func (m *MyMessageNonNull) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessageNonNull.Unmarshal(m, b)
}
func (m *MyMessageNonNull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessageNonNull.Marshal(b, m, deterministic)
}
func (dst *MyMessageNonNull) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessageNonNull.Merge(dst, src)
}
func (m *MyMessageNonNull) XXX_Size() int {
	return xxx_messageInfo_MyMessageNonNull.Size(m)
}
func (m *MyMessageNonNull) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessageNonNull.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessageNonNull proto.InternalMessageInfo

func (m *MyMessageNonNull) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MyMessageNonNull) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MyMessageNonNull) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *MyMessageNonNull) GetPet() []string {
	if m != nil {
		return m.Pet
	}
	return nil
}

func (m *MyMessageNonNull) GetInner() InnerMessage {
	if m != nil {
		return m.Inner
	}
	return InnerMessage{}
}

func (m *MyMessageNonNull) GetOthers() []OtherMessage {
	if m != nil {
		return m.Others
	}
	return nil
}

func (m *MyMessageNonNull) GetWeMustGoDeeper() RequiredInnerMessageNonNull {
	if m != nil {
		return m.WeMustGoDeeper
	}
	return RequiredInnerMessageNonNull{}
}

func (m *MyMessageNonNull) GetRepInner() []InnerMessage {
	if m != nil {
		return m.RepInner
	}
	return nil
}

func (m *MyMessageNonNull) GetBikeshed() MyMessageNonNull_Color {
	if m != nil {
		return m.Bikeshed
	}
	return MyMessageNonNull_RED
}

func (m *MyMessageNonNull) GetRepBytes() [][]byte {
	if m != nil {
		return m.RepBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*InnerMessage)(nil), "merge.InnerMessage")
	proto.RegisterType((*OtherMessage)(nil), "merge.OtherMessage")
	proto.RegisterType((*RequiredInnerMessageNonNull)(nil), "merge.RequiredInnerMessageNonNull")
	proto.RegisterType((*MyMessageNonNull)(nil), "merge.MyMessageNonNull")
	proto.RegisterEnum("merge.MyMessageNonNull_Color", MyMessageNonNull_Color_name, MyMessageNonNull_Color_value)
}

func init() { proto.RegisterFile("merge.proto", fileDescriptor_merge_9721c1510159f9a5) }

var fileDescriptor_merge_9721c1510159f9a5 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x8b, 0xd3, 0x4c,
	0x14, 0xdd, 0x34, 0x4d, 0xb7, 0xb9, 0xdb, 0x6f, 0xe9, 0x37, 0x16, 0x09, 0xae, 0x62, 0x08, 0x08,
	0xf1, 0xc1, 0x16, 0x57, 0x10, 0x7c, 0xb4, 0x6e, 0x5d, 0x04, 0xdb, 0xc5, 0x51, 0xf1, 0x31, 0xa4,
	0xed, 0xdd, 0x34, 0x6c, 0x3a, 0x37, 0x3b, 0x99, 0x58, 0xf2, 0x0f, 0xfd, 0x01, 0x3e, 0xfb, 0x5b,
	0x64, 0x66, 0xa2, 0x5b, 0x97, 0x45, 0xdf, 0xee, 0x99, 0x9c, 0x73, 0xe6, 0xdc, 0x33, 0x04, 0x8e,
	0xb6, 0x28, 0x33, 0x1c, 0x97, 0x92, 0x14, 0x31, 0xcf, 0x80, 0x07, 0xcf, 0xb2, 0x5c, 0x6d, 0xea,
	0xe5, 0x78, 0x45, 0xdb, 0x49, 0x46, 0x19, 0x4d, 0xcc, 0xd7, 0x65, 0x7d, 0x69, 0x90, 0x01, 0x66,
	0xb2, 0xaa, 0xe8, 0x13, 0x0c, 0xde, 0x09, 0x81, 0x72, 0x8e, 0x55, 0x95, 0x66, 0xc8, 0x18, 0x74,
	0x37, 0x54, 0xa9, 0xc0, 0x09, 0x9d, 0xd8, 0xe7, 0x66, 0xd6, 0x67, 0x25, 0x49, 0x15, 0x74, 0x42,
	0x27, 0xf6, 0xb8, 0x99, 0xd9, 0x43, 0xf0, 0x57, 0x24, 0x04, 0xae, 0x14, 0xae, 0x03, 0x37, 0x74,
	0xe2, 0x3e, 0xbf, 0x39, 0x88, 0x1a, 0x18, 0x5c, 0xa8, 0xcd, 0x8d, 0xeb, 0x10, 0xdc, 0x2b, 0x6c,
	0x8c, 0xa9, 0xcb, 0xf5, 0xc8, 0x46, 0xe0, 0x7d, 0x4d, 0x8b, 0x1a, 0x8d, 0xe9, 0x80, 0x5b, 0xc0,
	0xee, 0x43, 0x6f, 0x87, 0x79, 0xb6, 0x51, 0xc6, 0xb2, 0xc3, 0x5b, 0xc4, 0x9e, 0x82, 0x97, 0xeb,
	0x94, 0x41, 0x37, 0x74, 0xe2, 0xa3, 0xd3, 0x7b, 0x63, 0xbb, 0xf8, 0x7e, 0x72, 0x6e, 0x19, 0x51,
	0x09, 0x27, 0x1c, 0xaf, 0xeb, 0x5c, 0xe2, 0x7a, 0xff, 0xf3, 0x82, 0xc4, 0xa2, 0x2e, 0x0a, 0xf6,
	0x01, 0x82, 0x02, 0x29, 0xb9, 0xcc, 0x45, 0x5a, 0x14, 0x4d, 0xb2, 0x23, 0x91, 0xa4, 0x22, 0xa1,
	0x6a, 0x95, 0x4a, 0x13, 0xef, 0x6e, 0xf3, 0x69, 0xf7, 0xdb, 0x8f, 0xc7, 0x07, 0x7c, 0x54, 0x20,
	0xbd, 0xb5, 0xca, 0x2f, 0x24, 0x5e, 0x8b, 0x0b, 0x2d, 0x8b, 0xbe, 0xbb, 0x30, 0x9c, 0x37, 0xb7,
	0xee, 0x19, 0x81, 0xb7, 0xa2, 0x5a, 0xd8, 0x22, 0x3d, 0x6e, 0x81, 0x6e, 0x52, 0xa4, 0x5b, 0xbb,
	0xb4, 0xcf, 0xcd, 0xac, 0x99, 0xd7, 0x35, 0x29, 0x34, 0x2b, 0xfb, 0xdc, 0x02, 0xdd, 0x58, 0x89,
	0x2a, 0xe8, 0x86, 0x6e, 0xec, 0x73, 0x3d, 0xb2, 0xc9, 0xaf, 0x0e, 0xbc, 0x7f, 0xc5, 0xb4, 0x3c,
	0xf6, 0x1c, 0x7a, 0xa4, 0x1f, 0xa1, 0x0a, 0x7a, 0xa1, 0xbb, 0xa7, 0xd8, 0x7f, 0x99, 0x56, 0xd1,
	0x12, 0xd9, 0x47, 0xf8, 0x7f, 0x87, 0xc9, 0xb6, 0xae, 0x54, 0x92, 0x51, 0xb2, 0x46, 0x2c, 0x51,
	0x06, 0xff, 0x99, 0xfb, 0xa2, 0x56, 0xfd, 0x97, 0x72, 0x5b, 0xb3, 0xe3, 0x1d, 0xce, 0xeb, 0x4a,
	0x9d, 0xd3, 0x99, 0xd1, 0xb3, 0x97, 0xe0, 0x4b, 0x2c, 0x13, 0x1b, 0x7e, 0xf0, 0x47, 0x94, 0x3b,
	0xc2, 0xf7, 0x25, 0x96, 0xe6, 0x98, 0xbd, 0x82, 0xfe, 0x32, 0xbf, 0xc2, 0x6a, 0x83, 0xeb, 0xe0,
	0x30, 0x74, 0xe2, 0xe3, 0xd3, 0x47, 0xad, 0xec, 0x76, 0xdb, 0xe3, 0x37, 0x54, 0x90, 0xe4, 0xbf,
	0xe9, 0xec, 0xc4, 0x5e, 0xb9, 0x6c, 0x14, 0x56, 0x01, 0x84, 0x6e, 0x3c, 0x30, 0xbe, 0x53, 0x8d,
	0xa3, 0x27, 0xe0, 0x19, 0x3e, 0x3b, 0x04, 0x97, 0xcf, 0xce, 0x86, 0x07, 0xcc, 0x07, 0xef, 0x9c,
	0xcf, 0x66, 0x8b, 0xa1, 0xc3, 0xfa, 0xd0, 0x9d, 0xbe, 0xff, 0x3c, 0x1b, 0x76, 0x96, 0x3d, 0xf3,
	0x83, 0xbc, 0xf8, 0x19, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xa1, 0xc9, 0x5c, 0x65, 0x03, 0x00, 0x00,
}
