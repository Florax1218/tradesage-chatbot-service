// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: internal/proto/message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Basic message structure
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId    string `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Read        bool   `protobuf:"varint,5,opt,name=read,proto3" json:"read,omitempty"`
	SentAt      int64  `protobuf:"varint,6,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_internal_proto_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internal_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *Message) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId string `protobuf:"bytes,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_internal_proto_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *SendMessageRequest) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	mi := &file_internal_proto_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	mi := &file_internal_proto_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessagesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetMessagesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Total    int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	mi := &file_internal_proto_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *GetMessagesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_internal_proto_message_proto protoreflect.FileDescriptor

var file_internal_proto_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0xa0, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x12,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xec, 0x01,
	0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x2d,
	0x72, 0x65, 0x70, 0x6f, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_message_proto_rawDescOnce sync.Once
	file_internal_proto_message_proto_rawDescData = file_internal_proto_message_proto_rawDesc
)

func file_internal_proto_message_proto_rawDescGZIP() []byte {
	file_internal_proto_message_proto_rawDescOnce.Do(func() {
		file_internal_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_message_proto_rawDescData)
	})
	return file_internal_proto_message_proto_rawDescData
}

var file_internal_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_message_proto_goTypes = []any{
	(*Message)(nil),             // 0: messaging.Message
	(*SendMessageRequest)(nil),  // 1: messaging.SendMessageRequest
	(*SendMessageResponse)(nil), // 2: messaging.SendMessageResponse
	(*GetMessagesRequest)(nil),  // 3: messaging.GetMessagesRequest
	(*GetMessagesResponse)(nil), // 4: messaging.GetMessagesResponse
}
var file_internal_proto_message_proto_depIdxs = []int32{
	0, // 0: messaging.GetMessagesResponse.messages:type_name -> messaging.Message
	1, // 1: messaging.MessagingService.SendMessage:input_type -> messaging.SendMessageRequest
	3, // 2: messaging.MessagingService.GetMessages:input_type -> messaging.GetMessagesRequest
	0, // 3: messaging.MessagingService.StreamMessages:input_type -> messaging.Message
	2, // 4: messaging.MessagingService.SendMessage:output_type -> messaging.SendMessageResponse
	4, // 5: messaging.MessagingService.GetMessages:output_type -> messaging.GetMessagesResponse
	0, // 6: messaging.MessagingService.StreamMessages:output_type -> messaging.Message
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_message_proto_init() }
func file_internal_proto_message_proto_init() {
	if File_internal_proto_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_message_proto_goTypes,
		DependencyIndexes: file_internal_proto_message_proto_depIdxs,
		MessageInfos:      file_internal_proto_message_proto_msgTypes,
	}.Build()
	File_internal_proto_message_proto = out.File
	file_internal_proto_message_proto_rawDesc = nil
	file_internal_proto_message_proto_goTypes = nil
	file_internal_proto_message_proto_depIdxs = nil
}
