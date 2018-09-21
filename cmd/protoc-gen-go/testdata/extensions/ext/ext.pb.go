// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/ext/ext.proto

package ext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	base "google.golang.org/proto/cmd/protoc-gen-go/testdata/extensions/base"
	extra "google.golang.org/proto/cmd/protoc-gen-go/testdata/extensions/extra"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
)

var Enum_name = map[int32]string{
	0: "ZERO",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (x *Enum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum_value, data, "Enum")
	if err != nil {
		return err
	}
	*x = Enum(value)
	return nil
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{0}
}

type Message struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message_M struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_M) Reset()         { *m = Message_M{} }
func (m *Message_M) String() string { return proto.CompactTextString(m) }
func (*Message_M) ProtoMessage()    {}
func (*Message_M) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{0, 0}
}

func (m *Message_M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_M.Unmarshal(m, b)
}
func (m *Message_M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_M.Marshal(b, m, deterministic)
}
func (m *Message_M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_M.Merge(m, src)
}
func (m *Message_M) XXX_Size() int {
	return xxx_messageInfo_Message_M.Size(m)
}
func (m *Message_M) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_M.DiscardUnknown(m)
}

var xxx_messageInfo_Message_M proto.InternalMessageInfo

type ExtensionGroup struct {
	ExtensionGroup       *string  `protobuf:"bytes,120,opt,name=extension_group,json=extensionGroup" json:"extension_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionGroup) Reset()         { *m = ExtensionGroup{} }
func (m *ExtensionGroup) String() string { return proto.CompactTextString(m) }
func (*ExtensionGroup) ProtoMessage()    {}
func (*ExtensionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{1}
}

func (m *ExtensionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionGroup.Unmarshal(m, b)
}
func (m *ExtensionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionGroup.Marshal(b, m, deterministic)
}
func (m *ExtensionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionGroup.Merge(m, src)
}
func (m *ExtensionGroup) XXX_Size() int {
	return xxx_messageInfo_ExtensionGroup.Size(m)
}
func (m *ExtensionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionGroup proto.InternalMessageInfo

func (m *ExtensionGroup) GetExtensionGroup() string {
	if m != nil && m.ExtensionGroup != nil {
		return *m.ExtensionGroup
	}
	return ""
}

// Extend in the scope of another type.
type ExtendingMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtendingMessage) Reset()         { *m = ExtendingMessage{} }
func (m *ExtendingMessage) String() string { return proto.CompactTextString(m) }
func (*ExtendingMessage) ProtoMessage()    {}
func (*ExtendingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{2}
}

func (m *ExtendingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendingMessage.Unmarshal(m, b)
}
func (m *ExtendingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendingMessage.Marshal(b, m, deterministic)
}
func (m *ExtendingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendingMessage.Merge(m, src)
}
func (m *ExtendingMessage) XXX_Size() int {
	return xxx_messageInfo_ExtendingMessage.Size(m)
}
func (m *ExtendingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendingMessage proto.InternalMessageInfo

var E_ExtendingMessage_ExtendingMessageString = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         200,
	Name:          "goproto.protoc.extension.ext.ExtendingMessage.extending_message_string",
	Tag:           "bytes,200,opt,name=extending_message_string",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtendingMessage_ExtendingMessageSubmessage = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*ExtendingMessage_ExtendingMessageSubmessage)(nil),
	Field:         201,
	Name:          "goproto.protoc.extension.ext.ExtendingMessage.extending_message_submessage",
	Tag:           "bytes,201,opt,name=extending_message_submessage",
	Filename:      "extensions/ext/ext.proto",
}

type ExtendingMessage_ExtendingMessageSubmessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtendingMessage_ExtendingMessageSubmessage) Reset() {
	*m = ExtendingMessage_ExtendingMessageSubmessage{}
}
func (m *ExtendingMessage_ExtendingMessageSubmessage) String() string {
	return proto.CompactTextString(m)
}
func (*ExtendingMessage_ExtendingMessageSubmessage) ProtoMessage() {}
func (*ExtendingMessage_ExtendingMessageSubmessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{2, 0}
}

func (m *ExtendingMessage_ExtendingMessageSubmessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage.Unmarshal(m, b)
}
func (m *ExtendingMessage_ExtendingMessageSubmessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage.Marshal(b, m, deterministic)
}
func (m *ExtendingMessage_ExtendingMessageSubmessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage.Merge(m, src)
}
func (m *ExtendingMessage_ExtendingMessageSubmessage) XXX_Size() int {
	return xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage.Size(m)
}
func (m *ExtendingMessage_ExtendingMessageSubmessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendingMessage_ExtendingMessageSubmessage proto.InternalMessageInfo

type RepeatedGroup struct {
	RepeatedXGroup       []string `protobuf:"bytes,319,rep,name=repeated_x_group,json=repeatedXGroup" json:"repeated_x_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatedGroup) Reset()         { *m = RepeatedGroup{} }
func (m *RepeatedGroup) String() string { return proto.CompactTextString(m) }
func (*RepeatedGroup) ProtoMessage()    {}
func (*RepeatedGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{3}
}

func (m *RepeatedGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedGroup.Unmarshal(m, b)
}
func (m *RepeatedGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedGroup.Marshal(b, m, deterministic)
}
func (m *RepeatedGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedGroup.Merge(m, src)
}
func (m *RepeatedGroup) XXX_Size() int {
	return xxx_messageInfo_RepeatedGroup.Size(m)
}
func (m *RepeatedGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedGroup.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedGroup proto.InternalMessageInfo

func (m *RepeatedGroup) GetRepeatedXGroup() []string {
	if m != nil {
		return m.RepeatedXGroup
	}
	return nil
}

// An extension of an extension.
type Extendable struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Extendable) Reset()         { *m = Extendable{} }
func (m *Extendable) String() string { return proto.CompactTextString(m) }
func (*Extendable) ProtoMessage()    {}
func (*Extendable) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{4}
}

var extRange_Extendable = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*Extendable) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Extendable
}

func (m *Extendable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extendable.Unmarshal(m, b)
}
func (m *Extendable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extendable.Marshal(b, m, deterministic)
}
func (m *Extendable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extendable.Merge(m, src)
}
func (m *Extendable) XXX_Size() int {
	return xxx_messageInfo_Extendable.Size(m)
}
func (m *Extendable) XXX_DiscardUnknown() {
	xxx_messageInfo_Extendable.DiscardUnknown(m)
}

var xxx_messageInfo_Extendable proto.InternalMessageInfo

// Message set wire format.
type MessageSetWireFormatExtension struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageSetWireFormatExtension) Reset()         { *m = MessageSetWireFormatExtension{} }
func (m *MessageSetWireFormatExtension) String() string { return proto.CompactTextString(m) }
func (*MessageSetWireFormatExtension) ProtoMessage()    {}
func (*MessageSetWireFormatExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf470ef4907b23cb, []int{5}
}

func (m *MessageSetWireFormatExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSetWireFormatExtension.Unmarshal(m, b)
}
func (m *MessageSetWireFormatExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSetWireFormatExtension.Marshal(b, m, deterministic)
}
func (m *MessageSetWireFormatExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSetWireFormatExtension.Merge(m, src)
}
func (m *MessageSetWireFormatExtension) XXX_Size() int {
	return xxx_messageInfo_MessageSetWireFormatExtension.Size(m)
}
func (m *MessageSetWireFormatExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSetWireFormatExtension.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSetWireFormatExtension proto.InternalMessageInfo

var E_MessageSetWireFormatExtension_MessageSetField = &proto.ExtensionDesc{
	ExtendedType:  (*base.MessageSetWireFormatMessage)(nil),
	ExtensionType: (*MessageSetWireFormatExtension)(nil),
	Field:         100,
	Name:          "goproto.protoc.extension.ext.MessageSetWireFormatExtension.message_set_field",
	Tag:           "bytes,100,opt,name=message_set_field",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionBool = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*bool)(nil),
	Field:         101,
	Name:          "goproto.protoc.extension.ext.extension_bool",
	Tag:           "varint,101,opt,name=extension_bool",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionEnum = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*Enum)(nil),
	Field:         102,
	Name:          "goproto.protoc.extension.ext.extension_enum",
	Tag:           "varint,102,opt,name=extension_enum,enum=goproto.protoc.extension.ext.Enum",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         103,
	Name:          "goproto.protoc.extension.ext.extension_int32",
	Tag:           "varint,103,opt,name=extension_int32",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         104,
	Name:          "goproto.protoc.extension.ext.extension_sint32",
	Tag:           "zigzag32,104,opt,name=extension_sint32",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         105,
	Name:          "goproto.protoc.extension.ext.extension_uint32",
	Tag:           "varint,105,opt,name=extension_uint32",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         106,
	Name:          "goproto.protoc.extension.ext.extension_int64",
	Tag:           "varint,106,opt,name=extension_int64",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         107,
	Name:          "goproto.protoc.extension.ext.extension_sint64",
	Tag:           "zigzag64,107,opt,name=extension_sint64",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         108,
	Name:          "goproto.protoc.extension.ext.extension_uint64",
	Tag:           "varint,108,opt,name=extension_uint64",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         109,
	Name:          "goproto.protoc.extension.ext.extension_sfixed32",
	Tag:           "fixed32,109,opt,name=extension_sfixed32",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         110,
	Name:          "goproto.protoc.extension.ext.extension_fixed32",
	Tag:           "fixed32,110,opt,name=extension_fixed32",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionFloat = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*float32)(nil),
	Field:         111,
	Name:          "goproto.protoc.extension.ext.extension_float",
	Tag:           "fixed32,111,opt,name=extension_float",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         112,
	Name:          "goproto.protoc.extension.ext.extension_sfixed64",
	Tag:           "fixed64,112,opt,name=extension_sfixed64",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         113,
	Name:          "goproto.protoc.extension.ext.extension_fixed64",
	Tag:           "fixed64,113,opt,name=extension_fixed64",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionDouble = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*float64)(nil),
	Field:         114,
	Name:          "goproto.protoc.extension.ext.extension_double",
	Tag:           "fixed64,114,opt,name=extension_double",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionString = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         115,
	Name:          "goproto.protoc.extension.ext.extension_string",
	Tag:           "bytes,115,opt,name=extension_string",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtensionBytes = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]byte)(nil),
	Field:         116,
	Name:          "goproto.protoc.extension.ext.extension_bytes",
	Tag:           "bytes,116,opt,name=extension_bytes",
	Filename:      "extensions/ext/ext.proto",
}

var E_Extension_Message = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*Message)(nil),
	Field:         117,
	Name:          "goproto.protoc.extension.ext.extension_Message",
	Tag:           "bytes,117,opt,name=extension_Message",
	Filename:      "extensions/ext/ext.proto",
}

var E_Extension_MessageM = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*Message_M)(nil),
	Field:         118,
	Name:          "goproto.protoc.extension.ext.extension_MessageM",
	Tag:           "bytes,118,opt,name=extension_MessageM",
	Filename:      "extensions/ext/ext.proto",
}

var E_Extensiongroup = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*ExtensionGroup)(nil),
	Field:         119,
	Name:          "goproto.protoc.extension.ext.extensiongroup",
	Tag:           "group,119,opt,name=ExtensionGroup",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtraMessage = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*extra.ExtraMessage)(nil),
	Field:         9,
	Name:          "goproto.protoc.extension.ext.extra_message",
	Tag:           "bytes,9,opt,name=extra_message",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXBool = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]bool)(nil),
	Field:         301,
	Name:          "goproto.protoc.extension.ext.repeated_x_bool",
	Tag:           "varint,301,rep,name=repeated_x_bool",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXEnum = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]Enum)(nil),
	Field:         302,
	Name:          "goproto.protoc.extension.ext.repeated_x_enum",
	Tag:           "varint,302,rep,name=repeated_x_enum,enum=goproto.protoc.extension.ext.Enum",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         303,
	Name:          "goproto.protoc.extension.ext.repeated_x_int32",
	Tag:           "varint,303,rep,name=repeated_x_int32",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         304,
	Name:          "goproto.protoc.extension.ext.repeated_x_sint32",
	Tag:           "zigzag32,304,rep,name=repeated_x_sint32",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]uint32)(nil),
	Field:         305,
	Name:          "goproto.protoc.extension.ext.repeated_x_uint32",
	Tag:           "varint,305,rep,name=repeated_x_uint32",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         306,
	Name:          "goproto.protoc.extension.ext.repeated_x_int64",
	Tag:           "varint,306,rep,name=repeated_x_int64",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         307,
	Name:          "goproto.protoc.extension.ext.repeated_x_sint64",
	Tag:           "zigzag64,307,rep,name=repeated_x_sint64",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]uint64)(nil),
	Field:         308,
	Name:          "goproto.protoc.extension.ext.repeated_x_uint64",
	Tag:           "varint,308,rep,name=repeated_x_uint64",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int32)(nil),
	Field:         309,
	Name:          "goproto.protoc.extension.ext.repeated_x_sfixed32",
	Tag:           "fixed32,309,rep,name=repeated_x_sfixed32",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]uint32)(nil),
	Field:         310,
	Name:          "goproto.protoc.extension.ext.repeated_x_fixed32",
	Tag:           "fixed32,310,rep,name=repeated_x_fixed32",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXFloat = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]float32)(nil),
	Field:         311,
	Name:          "goproto.protoc.extension.ext.repeated_x_float",
	Tag:           "fixed32,311,rep,name=repeated_x_float",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]int64)(nil),
	Field:         312,
	Name:          "goproto.protoc.extension.ext.repeated_x_sfixed64",
	Tag:           "fixed64,312,rep,name=repeated_x_sfixed64",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]uint64)(nil),
	Field:         313,
	Name:          "goproto.protoc.extension.ext.repeated_x_fixed64",
	Tag:           "fixed64,313,rep,name=repeated_x_fixed64",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXDouble = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]float64)(nil),
	Field:         314,
	Name:          "goproto.protoc.extension.ext.repeated_x_double",
	Tag:           "fixed64,314,rep,name=repeated_x_double",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXString = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]string)(nil),
	Field:         315,
	Name:          "goproto.protoc.extension.ext.repeated_x_string",
	Tag:           "bytes,315,rep,name=repeated_x_string",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedXBytes = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([][]byte)(nil),
	Field:         316,
	Name:          "goproto.protoc.extension.ext.repeated_x_bytes",
	Tag:           "bytes,316,rep,name=repeated_x_bytes",
	Filename:      "extensions/ext/ext.proto",
}

var E_RepeatedX_Message = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]*Message)(nil),
	Field:         317,
	Name:          "goproto.protoc.extension.ext.repeated_x_Message",
	Tag:           "bytes,317,rep,name=repeated_x_Message",
	Filename:      "extensions/ext/ext.proto",
}

var E_Repeatedgroup = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: ([]*RepeatedGroup)(nil),
	Field:         318,
	Name:          "goproto.protoc.extension.ext.repeatedgroup",
	Tag:           "group,318,rep,name=RepeatedGroup",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtendableField = &proto.ExtensionDesc{
	ExtendedType:  (*base.BaseMessage)(nil),
	ExtensionType: (*Extendable)(nil),
	Field:         400,
	Name:          "goproto.protoc.extension.ext.extendable_field",
	Tag:           "bytes,400,opt,name=extendable_field",
	Filename:      "extensions/ext/ext.proto",
}

var E_ExtendableStringField = &proto.ExtensionDesc{
	ExtendedType:  (*Extendable)(nil),
	ExtensionType: (*string)(nil),
	Field:         1,
	Name:          "goproto.protoc.extension.ext.extendable_string_field",
	Tag:           "bytes,1,opt,name=extendable_string_field",
	Filename:      "extensions/ext/ext.proto",
}

func init() {
	proto.RegisterType((*Message)(nil), "goproto.protoc.extension.ext.Message")
	proto.RegisterType((*Message_M)(nil), "goproto.protoc.extension.ext.Message.M")
	proto.RegisterType((*ExtensionGroup)(nil), "goproto.protoc.extension.ext.ExtensionGroup")
	proto.RegisterType((*ExtendingMessage)(nil), "goproto.protoc.extension.ext.ExtendingMessage")
	proto.RegisterType((*ExtendingMessage_ExtendingMessageSubmessage)(nil), "goproto.protoc.extension.ext.ExtendingMessage.ExtendingMessageSubmessage")
	proto.RegisterType((*RepeatedGroup)(nil), "goproto.protoc.extension.ext.RepeatedGroup")
	proto.RegisterType((*Extendable)(nil), "goproto.protoc.extension.ext.Extendable")
	proto.RegisterType((*MessageSetWireFormatExtension)(nil), "goproto.protoc.extension.ext.MessageSetWireFormatExtension")
	proto.RegisterEnum("goproto.protoc.extension.ext.Enum", Enum_name, Enum_value)
	proto.RegisterExtension(E_ExtendingMessage_ExtendingMessageString)
	proto.RegisterExtension(E_ExtendingMessage_ExtendingMessageSubmessage)
	proto.RegisterExtension(E_MessageSetWireFormatExtension_MessageSetField)
	proto.RegisterExtension(E_ExtensionBool)
	proto.RegisterExtension(E_ExtensionEnum)
	proto.RegisterExtension(E_ExtensionInt32)
	proto.RegisterExtension(E_ExtensionSint32)
	proto.RegisterExtension(E_ExtensionUint32)
	proto.RegisterExtension(E_ExtensionInt64)
	proto.RegisterExtension(E_ExtensionSint64)
	proto.RegisterExtension(E_ExtensionUint64)
	proto.RegisterExtension(E_ExtensionSfixed32)
	proto.RegisterExtension(E_ExtensionFixed32)
	proto.RegisterExtension(E_ExtensionFloat)
	proto.RegisterExtension(E_ExtensionSfixed64)
	proto.RegisterExtension(E_ExtensionFixed64)
	proto.RegisterExtension(E_ExtensionDouble)
	proto.RegisterExtension(E_ExtensionString)
	proto.RegisterExtension(E_ExtensionBytes)
	proto.RegisterExtension(E_Extension_Message)
	proto.RegisterExtension(E_Extension_MessageM)
	proto.RegisterExtension(E_Extensiongroup)
	proto.RegisterExtension(E_ExtraMessage)
	proto.RegisterExtension(E_RepeatedXBool)
	proto.RegisterExtension(E_RepeatedXEnum)
	proto.RegisterExtension(E_RepeatedXInt32)
	proto.RegisterExtension(E_RepeatedXSint32)
	proto.RegisterExtension(E_RepeatedXUint32)
	proto.RegisterExtension(E_RepeatedXInt64)
	proto.RegisterExtension(E_RepeatedXSint64)
	proto.RegisterExtension(E_RepeatedXUint64)
	proto.RegisterExtension(E_RepeatedXSfixed32)
	proto.RegisterExtension(E_RepeatedXFixed32)
	proto.RegisterExtension(E_RepeatedXFloat)
	proto.RegisterExtension(E_RepeatedXSfixed64)
	proto.RegisterExtension(E_RepeatedXFixed64)
	proto.RegisterExtension(E_RepeatedXDouble)
	proto.RegisterExtension(E_RepeatedXString)
	proto.RegisterExtension(E_RepeatedXBytes)
	proto.RegisterExtension(E_RepeatedX_Message)
	proto.RegisterExtension(E_Repeatedgroup)
	proto.RegisterExtension(E_ExtendableField)
	proto.RegisterExtension(E_ExtendableStringField)
}

func init() { proto.RegisterFile("extensions/ext/ext.proto", fileDescriptor_bf470ef4907b23cb) }

var fileDescriptor_bf470ef4907b23cb = []byte{
	// 1105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x98, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x99, 0x24, 0xbb, 0xcd, 0x0e, 0x4d, 0x9b, 0x18, 0xb1, 0x44, 0x51, 0x91, 0x46, 0x91,
	0xd0, 0x0e, 0x85, 0x75, 0xa5, 0x10, 0x45, 0xc2, 0x3c, 0x51, 0x68, 0xd1, 0x3e, 0x44, 0x08, 0x57,
	0x55, 0xbb, 0xcb, 0x43, 0x70, 0xea, 0xa9, 0x31, 0x38, 0x9e, 0xe2, 0x0b, 0x1b, 0x1e, 0x58, 0x45,
	0x7c, 0x02, 0xbe, 0x03, 0xcf, 0xdc, 0x6f, 0xcb, 0xfd, 0x09, 0x09, 0x5e, 0xf9, 0x44, 0x68, 0x66,
	0x7c, 0x19, 0x3b, 0x69, 0x36, 0xe3, 0x87, 0x56, 0x93, 0xb1, 0xf3, 0x3b, 0xc7, 0xff, 0x73, 0xe2,
	0xff, 0xb1, 0x61, 0x97, 0xcc, 0x23, 0xe2, 0x87, 0x2e, 0xf5, 0xc3, 0x03, 0x32, 0x8f, 0xd8, 0x9f,
	0x7e, 0x15, 0xd0, 0x88, 0x6a, 0x7b, 0x0e, 0xe5, 0x0b, 0xf1, 0xf1, 0x42, 0xcf, 0x4e, 0x64, 0xab,
	0x5e, 0x4f, 0xfa, 0xde, 0xd4, 0x0a, 0x09, 0xff, 0x27, 0x4e, 0xed, 0xed, 0x15, 0x99, 0x81, 0x25,
	0xfe, 0x8b, 0xa3, 0xfd, 0x3e, 0xdc, 0x1a, 0x93, 0x30, 0xb4, 0x1c, 0xa2, 0x69, 0xb0, 0x61, 0x5b,
	0x91, 0xd5, 0x05, 0x08, 0xe0, 0x6d, 0x93, 0xaf, 0x7b, 0x75, 0x08, 0xc6, 0xfd, 0x57, 0xe1, 0xce,
	0x51, 0xca, 0x78, 0x2b, 0xa0, 0xf1, 0x95, 0x76, 0x07, 0xee, 0x66, 0xd4, 0x89, 0xc3, 0xb6, 0xba,
	0x73, 0x04, 0xf0, 0x2d, 0x73, 0x87, 0x14, 0x4e, 0xec, 0xff, 0x57, 0x83, 0x6d, 0xfe, 0x5d, 0xdb,
	0xf5, 0x9d, 0x24, 0x50, 0x6f, 0x0f, 0xf6, 0xca, 0x7b, 0x27, 0xf1, 0x74, 0x26, 0x56, 0x03, 0x92,
	0xa8, 0xc0, 0x8e, 0x4e, 0x92, 0xcd, 0x49, 0x18, 0x05, 0xae, 0xef, 0x68, 0xfb, 0xfa, 0xb5, 0x32,
	0xf0, 0x2b, 0x3e, 0xb4, 0x42, 0x92, 0x10, 0xbb, 0xff, 0x00, 0x9e, 0xd1, 0x6d, 0x52, 0x0e, 0xc5,
	0x51, 0x83, 0xc7, 0x00, 0xee, 0xad, 0x88, 0x93, 0xe5, 0xa1, 0x14, 0xeb, 0x5f, 0x16, 0xeb, 0xe9,
	0xc1, 0x3d, 0x7d, 0x5d, 0x95, 0xf4, 0xf2, 0x45, 0xeb, 0xd7, 0xab, 0x60, 0xf6, 0xc8, 0xb5, 0xc7,
	0xfa, 0x06, 0x6c, 0x99, 0xe4, 0x8a, 0x58, 0x11, 0xb1, 0x45, 0x39, 0x5e, 0x84, 0xed, 0x20, 0xd9,
	0x98, 0xcc, 0x93, 0x7a, 0xfc, 0x55, 0x43, 0x75, 0x56, 0x90, 0xf4, 0xc0, 0xb9, 0x28, 0xc8, 0x6d,
	0x08, 0x45, 0x54, 0x6b, 0xea, 0x91, 0xfd, 0x66, 0x13, 0xb4, 0x17, 0x8b, 0xc5, 0xa2, 0xd6, 0xff,
	0x1b, 0xc0, 0xe7, 0xd3, 0x48, 0x24, 0x3a, 0x73, 0x03, 0x72, 0x4c, 0x83, 0x99, 0x15, 0x65, 0x85,
	0x1f, 0x7c, 0x01, 0x60, 0x27, 0x93, 0x89, 0x44, 0x93, 0x4b, 0x97, 0x78, 0xb6, 0x66, 0x3c, 0x41,
	0xa5, 0x55, 0xcc, 0x54, 0x35, 0x9b, 0x8b, 0xf6, 0xda, 0x7a, 0xd1, 0xd6, 0x26, 0x65, 0xee, 0xce,
	0xb2, 0xc3, 0xc7, 0x2c, 0x9f, 0xfd, 0x36, 0x6c, 0x1c, 0xf9, 0xf1, 0x4c, 0x6b, 0xc2, 0xc6, 0x83,
	0x23, 0xf3, 0xed, 0xf6, 0x53, 0xc6, 0x3b, 0x30, 0x6f, 0xca, 0xc9, 0x94, 0x52, 0x4f, 0xa9, 0xb2,
	0x04, 0x01, 0xdc, 0x34, 0x5b, 0xd9, 0x19, 0x87, 0x94, 0x7a, 0x46, 0x2c, 0x23, 0x09, 0x0b, 0xa7,
	0x82, 0xbc, 0x44, 0x00, 0xef, 0x0c, 0xfa, 0x4f, 0xe8, 0x15, 0x3f, 0x9e, 0x49, 0x61, 0xd9, 0x47,
	0xe3, 0x44, 0xfe, 0xd5, 0xb9, 0x7e, 0xf4, 0xca, 0x40, 0x29, 0xae, 0x83, 0x00, 0xbe, 0x21, 0xfd,
	0x42, 0xef, 0x31, 0x82, 0x71, 0x0a, 0xdb, 0x39, 0x34, 0x54, 0xa7, 0xbe, 0x8f, 0x00, 0xee, 0x98,
	0x79, 0x62, 0x27, 0xee, 0x32, 0x36, 0x56, 0xc7, 0xba, 0x08, 0xe0, 0x96, 0x84, 0x3d, 0x15, 0xd8,
	0xb2, 0x04, 0xa3, 0xa1, 0x12, 0xf5, 0x03, 0x04, 0x70, 0xbd, 0x28, 0xc1, 0x68, 0xb8, 0x2c, 0x81,
	0x22, 0xf5, 0x43, 0x04, 0xb0, 0x56, 0x92, 0xa0, 0x8c, 0x8d, 0xd5, 0xb1, 0x1e, 0x02, 0xb8, 0x51,
	0x92, 0x60, 0x34, 0x34, 0xee, 0x43, 0x4d, 0xca, 0xf6, 0xd2, 0x9d, 0x13, 0x5b, 0x51, 0xdb, 0x19,
	0x02, 0x78, 0xd7, 0xec, 0xe4, 0xf9, 0x26, 0x10, 0xe3, 0x0c, 0xe6, 0x9b, 0x93, 0x2a, 0x64, 0x1f,
	0x01, 0xbc, 0x65, 0xe6, 0x97, 0x7d, 0x9c, 0x80, 0x0b, 0x65, 0xbb, 0xf4, 0xa8, 0x15, 0x29, 0x61,
	0x29, 0x02, 0xb8, 0x26, 0x95, 0xed, 0x98, 0x11, 0x56, 0x09, 0xa1, 0xa8, 0xf0, 0x15, 0x02, 0xb8,
	0xbd, 0x24, 0xc4, 0x68, 0xb8, 0x42, 0x08, 0x45, 0xf2, 0x47, 0x08, 0xe0, 0x9b, 0x65, 0x21, 0xca,
	0x3d, 0x61, 0xd3, 0x78, 0xea, 0xa9, 0x19, 0x4d, 0x80, 0x00, 0x06, 0x52, 0x4f, 0xbc, 0xc9, 0x11,
	0xa5, 0x0e, 0x56, 0xf7, 0xca, 0x90, 0x5b, 0xa5, 0xd4, 0xc1, 0x1c, 0x51, 0x2c, 0xdb, 0xf4, 0x93,
	0x88, 0x84, 0x4a, 0xd4, 0x88, 0x0f, 0x12, 0x79, 0xd9, 0x0e, 0x19, 0xc1, 0x78, 0x24, 0x6b, 0x3b,
	0xae, 0x60, 0xb6, 0x31, 0xb7, 0x8d, 0x17, 0x36, 0xb2, 0x0d, 0xa9, 0x04, 0xc9, 0x8e, 0xf1, 0x19,
	0x90, 0xfb, 0x26, 0xd9, 0x1d, 0x2b, 0x65, 0xf0, 0x31, 0xcf, 0xe0, 0xce, 0x46, 0x19, 0xe8, 0x63,
	0xa9, 0xc1, 0xd2, 0x68, 0xc6, 0x02, 0x48, 0x16, 0xc2, 0x0d, 0x5b, 0x29, 0x81, 0x87, 0x08, 0x60,
	0x38, 0x78, 0x79, 0x83, 0x71, 0x23, 0x1b, 0xc5, 0xcc, 0x52, 0x3c, 0xe3, 0x53, 0xd8, 0xe2, 0x83,
	0xe0, 0xa4, 0xca, 0xc0, 0x73, 0x8b, 0x2b, 0xb0, 0x3e, 0x81, 0xc0, 0x62, 0x29, 0x04, 0x56, 0x5a,
	0x8a, 0x6d, 0x22, 0x7d, 0x62, 0xbd, 0x25, 0xcd, 0x2c, 0xca, 0xbe, 0xfc, 0x25, 0x1b, 0x6f, 0x9a,
	0x66, 0x2b, 0x1b, 0x6f, 0xb8, 0x31, 0x3f, 0x2c, 0x40, 0x95, 0x9d, 0xf9, 0x2b, 0x06, 0xdd, 0xd0,
	0x9a, 0xb3, 0xc0, 0xdc, 0x9a, 0x4f, 0x0b, 0x13, 0x98, 0xba, 0xdd, 0x7d, 0xcd, 0x22, 0xdf, 0x90,
	0xa6, 0x35, 0x61, 0xce, 0x67, 0xb0, 0x23, 0x61, 0x2b, 0xb8, 0xf3, 0x37, 0x8c, 0xdb, 0x31, 0x33,
	0x55, 0xce, 0x13, 0x7b, 0x2e, 0x82, 0x2b, 0xf8, 0xf3, 0xb7, 0x0c, 0xdc, 0x92, 0xc0, 0xa7, 0x99,
	0xef, 0x17, 0x85, 0x50, 0xbc, 0x71, 0x7e, 0xc7, 0xb8, 0xf5, 0xa2, 0x10, 0xe2, 0x86, 0x5c, 0x12,
	0x42, 0x91, 0xfb, 0x3d, 0xe3, 0x6a, 0x25, 0x21, 0x96, 0xc0, 0x15, 0x5c, 0xfa, 0x07, 0x06, 0x6e,
	0x94, 0x84, 0x18, 0x0d, 0x8d, 0x77, 0xe1, 0x33, 0x72, 0xc6, 0x55, 0xdc, 0xf4, 0x47, 0x86, 0xde,
	0x35, 0x3b, 0x79, 0xce, 0xa9, 0x51, 0xdf, 0x87, 0x9a, 0x04, 0xaf, 0xc2, 0xfe, 0x89, 0xb1, 0xb7,
	0xcc, 0xac, 0x58, 0xe7, 0xa9, 0x55, 0x17, 0x0b, 0xa8, 0xee, 0xd5, 0x3f, 0x33, 0x70, 0x4d, 0x2a,
	0xa0, 0x30, 0xeb, 0x55, 0x72, 0x28, 0x2a, 0xfd, 0x98, 0x91, 0xdb, 0x4b, 0x72, 0x88, 0x91, 0xa8,
	0x2c, 0x87, 0x22, 0xfb, 0x17, 0xc6, 0xbe, 0x59, 0x96, 0x63, 0xa9, 0x3f, 0x2a, 0x38, 0xf6, 0xaf,
	0x8c, 0x0c, 0xa4, 0xfe, 0x48, 0x2c, 0xbb, 0xd4, 0xd1, 0xea, 0x9e, 0xfd, 0x9b, 0x78, 0xc0, 0x93,
	0x3a, 0x5a, 0x98, 0x76, 0xb1, 0x80, 0xea, 0xae, 0xfd, 0x3b, 0xe3, 0x6e, 0x4b, 0x05, 0x14, 0xb6,
	0xbd, 0x00, 0x05, 0x91, 0xab, 0x18, 0xf7, 0x1f, 0x8c, 0xbc, 0xb9, 0x73, 0x67, 0x09, 0xa4, 0x96,
	0xf1, 0x08, 0x66, 0x77, 0x5d, 0x75, 0xcb, 0xfc, 0x93, 0x05, 0x87, 0x83, 0x97, 0xd6, 0x07, 0x2f,
	0x3c, 0x57, 0x9b, 0xc5, 0x70, 0x4c, 0x02, 0x31, 0x4e, 0xf0, 0x87, 0xe7, 0xe4, 0x01, 0x58, 0x25,
	0x87, 0xcf, 0xeb, 0xdc, 0x36, 0xf1, 0x26, 0xaf, 0x09, 0x58, 0x88, 0x64, 0x22, 0xe3, 0x6b, 0xfe,
	0x78, 0x6b, 0xbc, 0x07, 0x9f, 0x93, 0x32, 0x10, 0x5d, 0x93, 0x24, 0xb2, 0x31, 0xb5, 0x2b, 0x5e,
	0x8c, 0x3c, 0x9b, 0x83, 0x44, 0xe7, 0xf0, 0x08, 0x87, 0x6f, 0x3c, 0x78, 0xdd, 0xa1, 0xd4, 0xf1,
	0x88, 0xee, 0x50, 0xcf, 0xf2, 0x1d, 0x9d, 0x06, 0xce, 0x01, 0x27, 0x1e, 0x5c, 0xcc, 0x6c, 0xb1,
	0xba, 0xb8, 0xeb, 0x10, 0xff, 0xae, 0x43, 0x0f, 0x22, 0x12, 0x46, 0xb6, 0x15, 0xf1, 0x57, 0x4a,
	0xd2, 0x3b, 0xa6, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xa0, 0x8e, 0xb5, 0xc8, 0x12, 0x00,
	0x00,
}