// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: packets.proto

package packets

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define your messages
type ChatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_packets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IdMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdMessage) Reset() {
	*x = IdMessage{}
	mi := &file_packets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMessage) ProtoMessage() {}

func (x *IdMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMessage.ProtoReflect.Descriptor instead.
func (*IdMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{1}
}

func (x *IdMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Define the main Packet message
type Packet struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	SenderId uint64                 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Types that are valid to be assigned to Msg:
	//
	//	*Packet_Chat
	//	*Packet_Id
	Msg           isPacket_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Packet) Reset() {
	*x = Packet{}
	mi := &file_packets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{2}
}

func (x *Packet) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *Packet) GetMsg() isPacket_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Packet) GetChat() *ChatMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Chat); ok {
			return x.Chat
		}
	}
	return nil
}

func (x *Packet) GetId() *IdMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Id); ok {
			return x.Id
		}
	}
	return nil
}

type isPacket_Msg interface {
	isPacket_Msg()
}

type Packet_Chat struct {
	Chat *ChatMessage `protobuf:"bytes,2,opt,name=chat,proto3,oneof"`
}

type Packet_Id struct {
	Id *IdMessage `protobuf:"bytes,3,opt,name=id,proto3,oneof"`
}

func (*Packet_Chat) isPacket_Msg() {}

func (*Packet_Id) isPacket_Msg() {}

var File_packets_proto protoreflect.FileDescriptor

const file_packets_proto_rawDesc = "" +
	"\n" +
	"\rpackets.proto\x12\apackets\"\x1f\n" +
	"\vChatMessage\x12\x10\n" +
	"\x03msg\x18\x01 \x01(\tR\x03msg\"\x1b\n" +
	"\tIdMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"~\n" +
	"\x06Packet\x12\x1b\n" +
	"\tsender_id\x18\x01 \x01(\x04R\bsenderId\x12*\n" +
	"\x04chat\x18\x02 \x01(\v2\x14.packets.ChatMessageH\x00R\x04chat\x12$\n" +
	"\x02id\x18\x03 \x01(\v2\x12.packets.IdMessageH\x00R\x02idB\x05\n" +
	"\x03msgB\rZ\vpkg/packetsb\x06proto3"

var (
	file_packets_proto_rawDescOnce sync.Once
	file_packets_proto_rawDescData []byte
)

func file_packets_proto_rawDescGZIP() []byte {
	file_packets_proto_rawDescOnce.Do(func() {
		file_packets_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_packets_proto_rawDesc), len(file_packets_proto_rawDesc)))
	})
	return file_packets_proto_rawDescData
}

var file_packets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packets_proto_goTypes = []any{
	(*ChatMessage)(nil), // 0: packets.ChatMessage
	(*IdMessage)(nil),   // 1: packets.IdMessage
	(*Packet)(nil),      // 2: packets.Packet
}
var file_packets_proto_depIdxs = []int32{
	0, // 0: packets.Packet.chat:type_name -> packets.ChatMessage
	1, // 1: packets.Packet.id:type_name -> packets.IdMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_packets_proto_init() }
func file_packets_proto_init() {
	if File_packets_proto != nil {
		return
	}
	file_packets_proto_msgTypes[2].OneofWrappers = []any{
		(*Packet_Chat)(nil),
		(*Packet_Id)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_packets_proto_rawDesc), len(file_packets_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packets_proto_goTypes,
		DependencyIndexes: file_packets_proto_depIdxs,
		MessageInfos:      file_packets_proto_msgTypes,
	}.Build()
	File_packets_proto = out.File
	file_packets_proto_goTypes = nil
	file_packets_proto_depIdxs = nil
}
