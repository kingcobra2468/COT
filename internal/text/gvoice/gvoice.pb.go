// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: gvoice.proto

package gvoice

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

type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GvoicePhoneNumber    *string `protobuf:"bytes,1,req,name=gvoice_phone_number,json=gvoicePhoneNumber" json:"gvoice_phone_number,omitempty"`
	RecipientPhoneNumber *string `protobuf:"bytes,2,req,name=recipient_phone_number,json=recipientPhoneNumber" json:"recipient_phone_number,omitempty"`
	Message              *string `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{0}
}

func (x *SendSMSRequest) GetGvoicePhoneNumber() string {
	if x != nil && x.GvoicePhoneNumber != nil {
		return *x.GvoicePhoneNumber
	}
	return ""
}

func (x *SendSMSRequest) GetRecipientPhoneNumber() string {
	if x != nil && x.RecipientPhoneNumber != nil {
		return *x.RecipientPhoneNumber
	}
	return ""
}

func (x *SendSMSRequest) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type SendSMSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success *bool   `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Error   *string `protobuf:"bytes,2,opt,name=error,def=" json:"error,omitempty"`
}

// Default values for SendSMSResponse fields.
const (
	Default_SendSMSResponse_Error = string("")
)

func (x *SendSMSResponse) Reset() {
	*x = SendSMSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSResponse) ProtoMessage() {}

func (x *SendSMSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSResponse.ProtoReflect.Descriptor instead.
func (*SendSMSResponse) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{1}
}

func (x *SendSMSResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *SendSMSResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Default_SendSMSResponse_Error
}

type FetchContactListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GvoicePhoneNumber *string `protobuf:"bytes,1,req,name=gvoice_phone_number,json=gvoicePhoneNumber" json:"gvoice_phone_number,omitempty"`
}

func (x *FetchContactListRequest) Reset() {
	*x = FetchContactListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchContactListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchContactListRequest) ProtoMessage() {}

func (x *FetchContactListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchContactListRequest.ProtoReflect.Descriptor instead.
func (*FetchContactListRequest) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{2}
}

func (x *FetchContactListRequest) GetGvoicePhoneNumber() string {
	if x != nil && x.GvoicePhoneNumber != nil {
		return *x.GvoicePhoneNumber
	}
	return ""
}

type FetchContactListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success               *bool    `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Error                 *string  `protobuf:"bytes,2,opt,name=error,def=" json:"error,omitempty"`
	RecipientPhoneNumbers []string `protobuf:"bytes,3,rep,name=recipient_phone_numbers,json=recipientPhoneNumbers" json:"recipient_phone_numbers,omitempty"`
}

// Default values for FetchContactListResponse fields.
const (
	Default_FetchContactListResponse_Error = string("")
)

func (x *FetchContactListResponse) Reset() {
	*x = FetchContactListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchContactListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchContactListResponse) ProtoMessage() {}

func (x *FetchContactListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchContactListResponse.ProtoReflect.Descriptor instead.
func (*FetchContactListResponse) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{3}
}

func (x *FetchContactListResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *FetchContactListResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Default_FetchContactListResponse_Error
}

func (x *FetchContactListResponse) GetRecipientPhoneNumbers() []string {
	if x != nil {
		return x.RecipientPhoneNumbers
	}
	return nil
}

type FetchGVoiceNumbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchGVoiceNumbersRequest) Reset() {
	*x = FetchGVoiceNumbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchGVoiceNumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchGVoiceNumbersRequest) ProtoMessage() {}

func (x *FetchGVoiceNumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchGVoiceNumbersRequest.ProtoReflect.Descriptor instead.
func (*FetchGVoiceNumbersRequest) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{4}
}

type FetchGVoiceNumbersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success            *bool    `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Error              *string  `protobuf:"bytes,2,opt,name=error,def=" json:"error,omitempty"`
	GvoicePhoneNumbers []string `protobuf:"bytes,3,rep,name=gvoice_phone_numbers,json=gvoicePhoneNumbers" json:"gvoice_phone_numbers,omitempty"`
}

// Default values for FetchGVoiceNumbersResponse fields.
const (
	Default_FetchGVoiceNumbersResponse_Error = string("")
)

func (x *FetchGVoiceNumbersResponse) Reset() {
	*x = FetchGVoiceNumbersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchGVoiceNumbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchGVoiceNumbersResponse) ProtoMessage() {}

func (x *FetchGVoiceNumbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchGVoiceNumbersResponse.ProtoReflect.Descriptor instead.
func (*FetchGVoiceNumbersResponse) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{5}
}

func (x *FetchGVoiceNumbersResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *FetchGVoiceNumbersResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Default_FetchGVoiceNumbersResponse_Error
}

func (x *FetchGVoiceNumbersResponse) GetGvoicePhoneNumbers() []string {
	if x != nil {
		return x.GvoicePhoneNumbers
	}
	return nil
}

type FetchContactHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GvoicePhoneNumber    *string `protobuf:"bytes,1,req,name=gvoice_phone_number,json=gvoicePhoneNumber" json:"gvoice_phone_number,omitempty"`
	RecipientPhoneNumber *string `protobuf:"bytes,2,req,name=recipient_phone_number,json=recipientPhoneNumber" json:"recipient_phone_number,omitempty"`
	NumMessages          *uint64 `protobuf:"varint,3,req,name=num_messages,json=numMessages,def=0" json:"num_messages,omitempty"`
}

// Default values for FetchContactHistoryRequest fields.
const (
	Default_FetchContactHistoryRequest_NumMessages = uint64(0)
)

func (x *FetchContactHistoryRequest) Reset() {
	*x = FetchContactHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchContactHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchContactHistoryRequest) ProtoMessage() {}

func (x *FetchContactHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchContactHistoryRequest.ProtoReflect.Descriptor instead.
func (*FetchContactHistoryRequest) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{6}
}

func (x *FetchContactHistoryRequest) GetGvoicePhoneNumber() string {
	if x != nil && x.GvoicePhoneNumber != nil {
		return *x.GvoicePhoneNumber
	}
	return ""
}

func (x *FetchContactHistoryRequest) GetRecipientPhoneNumber() string {
	if x != nil && x.RecipientPhoneNumber != nil {
		return *x.RecipientPhoneNumber
	}
	return ""
}

func (x *FetchContactHistoryRequest) GetNumMessages() uint64 {
	if x != nil && x.NumMessages != nil {
		return *x.NumMessages
	}
	return Default_FetchContactHistoryRequest_NumMessages
}

type FetchContactHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  *bool          `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Error    *string        `protobuf:"bytes,2,opt,name=error,def=" json:"error,omitempty"`
	Messages []*MessageNode `protobuf:"bytes,3,rep,name=messages" json:"messages,omitempty"`
}

// Default values for FetchContactHistoryResponse fields.
const (
	Default_FetchContactHistoryResponse_Error = string("")
)

func (x *FetchContactHistoryResponse) Reset() {
	*x = FetchContactHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchContactHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchContactHistoryResponse) ProtoMessage() {}

func (x *FetchContactHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchContactHistoryResponse.ProtoReflect.Descriptor instead.
func (*FetchContactHistoryResponse) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{7}
}

func (x *FetchContactHistoryResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *FetchContactHistoryResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Default_FetchContactHistoryResponse_Error
}

func (x *FetchContactHistoryResponse) GetMessages() []*MessageNode {
	if x != nil {
		return x.Messages
	}
	return nil
}

type MessageNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp       *uint64 `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"` // Timestamp of the message.
	MessageContents *string `protobuf:"bytes,2,req,name=message_contents,json=messageContents" json:"message_contents,omitempty"`
	Source          *bool   `protobuf:"varint,3,req,name=source" json:"source,omitempty"` // Message origin with True meaning recipient and False meaning Gvoice.
}

func (x *MessageNode) Reset() {
	*x = MessageNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gvoice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageNode) ProtoMessage() {}

func (x *MessageNode) ProtoReflect() protoreflect.Message {
	mi := &file_gvoice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageNode.ProtoReflect.Descriptor instead.
func (*MessageNode) Descriptor() ([]byte, []int) {
	return file_gvoice_proto_rawDescGZIP(), []int{8}
}

func (x *MessageNode) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *MessageNode) GetMessageContents() string {
	if x != nil && x.MessageContents != nil {
		return *x.MessageContents
	}
	return ""
}

func (x *MessageNode) GetSource() bool {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return false
}

var File_gvoice_proto protoreflect.FileDescriptor

var file_gvoice_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x11,
	0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x43, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x11,
	0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x15, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x47, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x47,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x1a, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x11, 0x67, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x79, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6e,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x08, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0xa2,
	0x02, 0x0a, 0x06, 0x47, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x4d, 0x53, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x47, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x47, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x69, 0x6e, 0x67, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x32, 0x34, 0x36, 0x38, 0x2f,
	0x63, 0x6f, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x76, 0x6f,
	0x69, 0x63, 0x65,
}

var (
	file_gvoice_proto_rawDescOnce sync.Once
	file_gvoice_proto_rawDescData = file_gvoice_proto_rawDesc
)

func file_gvoice_proto_rawDescGZIP() []byte {
	file_gvoice_proto_rawDescOnce.Do(func() {
		file_gvoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_gvoice_proto_rawDescData)
	})
	return file_gvoice_proto_rawDescData
}

var file_gvoice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gvoice_proto_goTypes = []interface{}{
	(*SendSMSRequest)(nil),              // 0: SendSMSRequest
	(*SendSMSResponse)(nil),             // 1: SendSMSResponse
	(*FetchContactListRequest)(nil),     // 2: FetchContactListRequest
	(*FetchContactListResponse)(nil),    // 3: FetchContactListResponse
	(*FetchGVoiceNumbersRequest)(nil),   // 4: FetchGVoiceNumbersRequest
	(*FetchGVoiceNumbersResponse)(nil),  // 5: FetchGVoiceNumbersResponse
	(*FetchContactHistoryRequest)(nil),  // 6: FetchContactHistoryRequest
	(*FetchContactHistoryResponse)(nil), // 7: FetchContactHistoryResponse
	(*MessageNode)(nil),                 // 8: MessageNode
}
var file_gvoice_proto_depIdxs = []int32{
	8, // 0: FetchContactHistoryResponse.messages:type_name -> MessageNode
	0, // 1: GVoice.SendSMS:input_type -> SendSMSRequest
	2, // 2: GVoice.GetContactList:input_type -> FetchContactListRequest
	4, // 3: GVoice.GetGVoiceNumbers:input_type -> FetchGVoiceNumbersRequest
	6, // 4: GVoice.GetContactHistory:input_type -> FetchContactHistoryRequest
	1, // 5: GVoice.SendSMS:output_type -> SendSMSResponse
	3, // 6: GVoice.GetContactList:output_type -> FetchContactListResponse
	5, // 7: GVoice.GetGVoiceNumbers:output_type -> FetchGVoiceNumbersResponse
	7, // 8: GVoice.GetContactHistory:output_type -> FetchContactHistoryResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gvoice_proto_init() }
func file_gvoice_proto_init() {
	if File_gvoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gvoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchContactListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchContactListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchGVoiceNumbersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchGVoiceNumbersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchContactHistoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchContactHistoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gvoice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gvoice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gvoice_proto_goTypes,
		DependencyIndexes: file_gvoice_proto_depIdxs,
		MessageInfos:      file_gvoice_proto_msgTypes,
	}.Build()
	File_gvoice_proto = out.File
	file_gvoice_proto_rawDesc = nil
	file_gvoice_proto_goTypes = nil
	file_gvoice_proto_depIdxs = nil
}
