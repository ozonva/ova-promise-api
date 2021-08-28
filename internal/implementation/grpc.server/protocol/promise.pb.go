// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/implementation/grpc.server/protocol/promise.proto

package promise

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

type SuccessMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessMessage) Reset() {
	*x = SuccessMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessMessage) ProtoMessage() {}

func (x *SuccessMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessMessage.ProtoReflect.Descriptor instead.
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{0}
}

func (x *SuccessMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Promise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID       int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Status       string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	DateDeadline string `protobuf:"bytes,5,opt,name=date_deadline,json=dateDeadline,proto3" json:"date_deadline,omitempty"`
	CreatedAt    string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Promise) Reset() {
	*x = Promise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promise) ProtoMessage() {}

func (x *Promise) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promise.ProtoReflect.Descriptor instead.
func (*Promise) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{3}
}

func (x *Promise) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Promise) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Promise) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Promise) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Promise) GetDateDeadline() string {
	if x != nil {
		return x.DateDeadline
	}
	return ""
}

func (x *Promise) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Promise) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListPromisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListPromisesRequest) Reset() {
	*x = ListPromisesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromisesRequest) ProtoMessage() {}

func (x *ListPromisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromisesRequest.ProtoReflect.Descriptor instead.
func (*ListPromisesRequest) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{4}
}

func (x *ListPromisesRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPromisesRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListPromisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Promises []*Promise `protobuf:"bytes,1,rep,name=Promises,proto3" json:"Promises,omitempty"`
}

func (x *ListPromisesResponse) Reset() {
	*x = ListPromisesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromisesResponse) ProtoMessage() {}

func (x *ListPromisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromisesResponse.ProtoReflect.Descriptor instead.
func (*ListPromisesResponse) Descriptor() ([]byte, []int) {
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP(), []int{5}
}

func (x *ListPromisesResponse) GetPromises() []*Promise {
	if x != nil {
		return x.Promises
	}
	return nil
}

var File_internal_implementation_grpc_server_protocol_promise_proto protoreflect.FileDescriptor

var file_internal_implementation_grpc_server_protocol_promise_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xce, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x49, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6d, 0x69,
	0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65,
	0x52, 0x08, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73, 0x32, 0xad, 0x02, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x43, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69,
	0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f,
	0x6f, 0x76, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_implementation_grpc_server_protocol_promise_proto_rawDescOnce sync.Once
	file_internal_implementation_grpc_server_protocol_promise_proto_rawDescData = file_internal_implementation_grpc_server_protocol_promise_proto_rawDesc
)

func file_internal_implementation_grpc_server_protocol_promise_proto_rawDescGZIP() []byte {
	file_internal_implementation_grpc_server_protocol_promise_proto_rawDescOnce.Do(func() {
		file_internal_implementation_grpc_server_protocol_promise_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_implementation_grpc_server_protocol_promise_proto_rawDescData)
	})
	return file_internal_implementation_grpc_server_protocol_promise_proto_rawDescData
}

var file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_implementation_grpc_server_protocol_promise_proto_goTypes = []interface{}{
	(*SuccessMessage)(nil),       // 0: promise_grpc.SuccessMessage
	(*UUID)(nil),                 // 1: promise_grpc.UUID
	(*CreateRequest)(nil),        // 2: promise_grpc.CreateRequest
	(*Promise)(nil),              // 3: promise_grpc.Promise
	(*ListPromisesRequest)(nil),  // 4: promise_grpc.ListPromisesRequest
	(*ListPromisesResponse)(nil), // 5: promise_grpc.ListPromisesResponse
}
var file_internal_implementation_grpc_server_protocol_promise_proto_depIdxs = []int32{
	3, // 0: promise_grpc.ListPromisesResponse.Promises:type_name -> promise_grpc.Promise
	2, // 1: promise_grpc.PromiseHandler.CreatePromise:input_type -> promise_grpc.CreateRequest
	1, // 2: promise_grpc.PromiseHandler.DescribePromise:input_type -> promise_grpc.UUID
	4, // 3: promise_grpc.PromiseHandler.ListPromises:input_type -> promise_grpc.ListPromisesRequest
	1, // 4: promise_grpc.PromiseHandler.RemovePromise:input_type -> promise_grpc.UUID
	3, // 5: promise_grpc.PromiseHandler.CreatePromise:output_type -> promise_grpc.Promise
	3, // 6: promise_grpc.PromiseHandler.DescribePromise:output_type -> promise_grpc.Promise
	5, // 7: promise_grpc.PromiseHandler.ListPromises:output_type -> promise_grpc.ListPromisesResponse
	0, // 8: promise_grpc.PromiseHandler.RemovePromise:output_type -> promise_grpc.SuccessMessage
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_implementation_grpc_server_protocol_promise_proto_init() }
func file_internal_implementation_grpc_server_protocol_promise_proto_init() {
	if File_internal_implementation_grpc_server_protocol_promise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessMessage); i {
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
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promise); i {
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
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromisesRequest); i {
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
		file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromisesResponse); i {
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
			RawDescriptor: file_internal_implementation_grpc_server_protocol_promise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_implementation_grpc_server_protocol_promise_proto_goTypes,
		DependencyIndexes: file_internal_implementation_grpc_server_protocol_promise_proto_depIdxs,
		MessageInfos:      file_internal_implementation_grpc_server_protocol_promise_proto_msgTypes,
	}.Build()
	File_internal_implementation_grpc_server_protocol_promise_proto = out.File
	file_internal_implementation_grpc_server_protocol_promise_proto_rawDesc = nil
	file_internal_implementation_grpc_server_protocol_promise_proto_goTypes = nil
	file_internal_implementation_grpc_server_protocol_promise_proto_depIdxs = nil
}