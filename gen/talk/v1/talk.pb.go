// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: talk/v1/talk.proto

package talkv1

import (
	v1 "github.com/submaline/auth-service/gen/types/v1"
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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *v1.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessage() *v1.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *v1.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetMessage() *v1.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendReadReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendReadReceiptRequest) Reset() {
	*x = SendReadReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReadReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReadReceiptRequest) ProtoMessage() {}

func (x *SendReadReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReadReceiptRequest.ProtoReflect.Descriptor instead.
func (*SendReadReceiptRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{2}
}

func (x *SendReadReceiptRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type SendReadReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendReadReceiptResponse) Reset() {
	*x = SendReadReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReadReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReadReceiptResponse) ProtoMessage() {}

func (x *SendReadReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReadReceiptResponse.ProtoReflect.Descriptor instead.
func (*SendReadReceiptResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{3}
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *v1.Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGroupRequest) GetGroup() *v1.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *v1.Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroupResponse) GetGroup() *v1.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *v1.Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGroupRequest) GetGroup() *v1.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *v1.Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupResponse) Reset() {
	*x = UpdateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupResponse) ProtoMessage() {}

func (x *UpdateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGroupResponse) GetGroup() *v1.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type InviteIntoGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *InviteIntoGroupRequest) Reset() {
	*x = InviteIntoGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteIntoGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteIntoGroupRequest) ProtoMessage() {}

func (x *InviteIntoGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteIntoGroupRequest.ProtoReflect.Descriptor instead.
func (*InviteIntoGroupRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{8}
}

func (x *InviteIntoGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *InviteIntoGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type InviteIntoGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteIntoGroupResponse) Reset() {
	*x = InviteIntoGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteIntoGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteIntoGroupResponse) ProtoMessage() {}

func (x *InviteIntoGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteIntoGroupResponse.ProtoReflect.Descriptor instead.
func (*InviteIntoGroupResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{9}
}

type JoinGroupViaInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *JoinGroupViaInvitationRequest) Reset() {
	*x = JoinGroupViaInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupViaInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupViaInvitationRequest) ProtoMessage() {}

func (x *JoinGroupViaInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupViaInvitationRequest.ProtoReflect.Descriptor instead.
func (*JoinGroupViaInvitationRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{10}
}

func (x *JoinGroupViaInvitationRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type JoinGroupViaInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinGroupViaInvitationResponse) Reset() {
	*x = JoinGroupViaInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupViaInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupViaInvitationResponse) ProtoMessage() {}

func (x *JoinGroupViaInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupViaInvitationResponse.ProtoReflect.Descriptor instead.
func (*JoinGroupViaInvitationResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{11}
}

type LeaveGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *LeaveGroupRequest) Reset() {
	*x = LeaveGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGroupRequest) ProtoMessage() {}

func (x *LeaveGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGroupRequest.ProtoReflect.Descriptor instead.
func (*LeaveGroupRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{12}
}

func (x *LeaveGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type LeaveGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveGroupResponse) Reset() {
	*x = LeaveGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGroupResponse) ProtoMessage() {}

func (x *LeaveGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGroupResponse.ProtoReflect.Descriptor instead.
func (*LeaveGroupResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{13}
}

type KickOutFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *KickOutFromGroupRequest) Reset() {
	*x = KickOutFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutFromGroupRequest) ProtoMessage() {}

func (x *KickOutFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutFromGroupRequest.ProtoReflect.Descriptor instead.
func (*KickOutFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{14}
}

func (x *KickOutFromGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *KickOutFromGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type KickOutFromGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KickOutFromGroupResponse) Reset() {
	*x = KickOutFromGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talk_v1_talk_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutFromGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutFromGroupResponse) ProtoMessage() {}

func (x *KickOutFromGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talk_v1_talk_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutFromGroupResponse.ProtoReflect.Descriptor instead.
func (*KickOutFromGroupResponse) Descriptor() ([]byte, []int) {
	return file_talk_v1_talk_proto_rawDescGZIP(), []int{15}
}

var File_talk_v1_talk_proto protoreflect.FileDescriptor

var file_talk_v1_talk_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x16, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3c, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3b, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x4c, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x6f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a,
	0x1d, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x61, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x61, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x11, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4d, 0x0a, 0x17, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x1a, 0x0a, 0x18, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x05, 0x0a,
	0x0b, 0x54, 0x61, 0x6c, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74,
	0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1b, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74,
	0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x49, 0x6e, 0x74, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x6f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74,
	0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x74,
	0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6b, 0x0a, 0x16, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x61,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x74, 0x61, 0x6c,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69,
	0x61, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x61, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0a, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x6c,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74,
	0x61, 0x6c, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x46, 0x72,
	0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x8b, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x54, 0x61, 0x6c, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x61,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61,
	0x6c, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x54, 0x61, 0x6c,
	0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x54, 0x61, 0x6c, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x13, 0x54, 0x61, 0x6c, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x54, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_talk_v1_talk_proto_rawDescOnce sync.Once
	file_talk_v1_talk_proto_rawDescData = file_talk_v1_talk_proto_rawDesc
)

func file_talk_v1_talk_proto_rawDescGZIP() []byte {
	file_talk_v1_talk_proto_rawDescOnce.Do(func() {
		file_talk_v1_talk_proto_rawDescData = protoimpl.X.CompressGZIP(file_talk_v1_talk_proto_rawDescData)
	})
	return file_talk_v1_talk_proto_rawDescData
}

var file_talk_v1_talk_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_talk_v1_talk_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),             // 0: talk.v1.SendMessageRequest
	(*SendMessageResponse)(nil),            // 1: talk.v1.SendMessageResponse
	(*SendReadReceiptRequest)(nil),         // 2: talk.v1.SendReadReceiptRequest
	(*SendReadReceiptResponse)(nil),        // 3: talk.v1.SendReadReceiptResponse
	(*CreateGroupRequest)(nil),             // 4: talk.v1.CreateGroupRequest
	(*CreateGroupResponse)(nil),            // 5: talk.v1.CreateGroupResponse
	(*UpdateGroupRequest)(nil),             // 6: talk.v1.UpdateGroupRequest
	(*UpdateGroupResponse)(nil),            // 7: talk.v1.UpdateGroupResponse
	(*InviteIntoGroupRequest)(nil),         // 8: talk.v1.InviteIntoGroupRequest
	(*InviteIntoGroupResponse)(nil),        // 9: talk.v1.InviteIntoGroupResponse
	(*JoinGroupViaInvitationRequest)(nil),  // 10: talk.v1.JoinGroupViaInvitationRequest
	(*JoinGroupViaInvitationResponse)(nil), // 11: talk.v1.JoinGroupViaInvitationResponse
	(*LeaveGroupRequest)(nil),              // 12: talk.v1.LeaveGroupRequest
	(*LeaveGroupResponse)(nil),             // 13: talk.v1.LeaveGroupResponse
	(*KickOutFromGroupRequest)(nil),        // 14: talk.v1.KickOutFromGroupRequest
	(*KickOutFromGroupResponse)(nil),       // 15: talk.v1.KickOutFromGroupResponse
	(*v1.Message)(nil),                     // 16: types.v1.Message
	(*v1.Group)(nil),                       // 17: types.v1.Group
}
var file_talk_v1_talk_proto_depIdxs = []int32{
	16, // 0: talk.v1.SendMessageRequest.message:type_name -> types.v1.Message
	16, // 1: talk.v1.SendMessageResponse.message:type_name -> types.v1.Message
	17, // 2: talk.v1.CreateGroupRequest.group:type_name -> types.v1.Group
	17, // 3: talk.v1.CreateGroupResponse.group:type_name -> types.v1.Group
	17, // 4: talk.v1.UpdateGroupRequest.group:type_name -> types.v1.Group
	17, // 5: talk.v1.UpdateGroupResponse.group:type_name -> types.v1.Group
	0,  // 6: talk.v1.TalkService.SendMessage:input_type -> talk.v1.SendMessageRequest
	2,  // 7: talk.v1.TalkService.SendReadReceipt:input_type -> talk.v1.SendReadReceiptRequest
	4,  // 8: talk.v1.TalkService.CreateGroup:input_type -> talk.v1.CreateGroupRequest
	6,  // 9: talk.v1.TalkService.UpdateGroup:input_type -> talk.v1.UpdateGroupRequest
	8,  // 10: talk.v1.TalkService.InviteIntoGroup:input_type -> talk.v1.InviteIntoGroupRequest
	10, // 11: talk.v1.TalkService.JoinGroupViaInvitation:input_type -> talk.v1.JoinGroupViaInvitationRequest
	12, // 12: talk.v1.TalkService.LeaveGroup:input_type -> talk.v1.LeaveGroupRequest
	14, // 13: talk.v1.TalkService.KickOutFromGroup:input_type -> talk.v1.KickOutFromGroupRequest
	1,  // 14: talk.v1.TalkService.SendMessage:output_type -> talk.v1.SendMessageResponse
	3,  // 15: talk.v1.TalkService.SendReadReceipt:output_type -> talk.v1.SendReadReceiptResponse
	5,  // 16: talk.v1.TalkService.CreateGroup:output_type -> talk.v1.CreateGroupResponse
	7,  // 17: talk.v1.TalkService.UpdateGroup:output_type -> talk.v1.UpdateGroupResponse
	9,  // 18: talk.v1.TalkService.InviteIntoGroup:output_type -> talk.v1.InviteIntoGroupResponse
	11, // 19: talk.v1.TalkService.JoinGroupViaInvitation:output_type -> talk.v1.JoinGroupViaInvitationResponse
	13, // 20: talk.v1.TalkService.LeaveGroup:output_type -> talk.v1.LeaveGroupResponse
	15, // 21: talk.v1.TalkService.KickOutFromGroup:output_type -> talk.v1.KickOutFromGroupResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_talk_v1_talk_proto_init() }
func file_talk_v1_talk_proto_init() {
	if File_talk_v1_talk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_talk_v1_talk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReadReceiptRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReadReceiptResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteIntoGroupRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteIntoGroupResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupViaInvitationRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupViaInvitationResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGroupRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGroupResponse); i {
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
		file_talk_v1_talk_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutFromGroupRequest); i {
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
		file_talk_v1_talk_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutFromGroupResponse); i {
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
			RawDescriptor: file_talk_v1_talk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_talk_v1_talk_proto_goTypes,
		DependencyIndexes: file_talk_v1_talk_proto_depIdxs,
		MessageInfos:      file_talk_v1_talk_proto_msgTypes,
	}.Build()
	File_talk_v1_talk_proto = out.File
	file_talk_v1_talk_proto_rawDesc = nil
	file_talk_v1_talk_proto_goTypes = nil
	file_talk_v1_talk_proto_depIdxs = nil
}
