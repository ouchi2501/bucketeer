// Copyright 2022 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/environment/command.proto

package environment

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

type CreateEnvironmentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Deprecated: Do not use.
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // optional
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId   string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateEnvironmentCommand) Reset() {
	*x = CreateEnvironmentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnvironmentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironmentCommand) ProtoMessage() {}

func (x *CreateEnvironmentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironmentCommand.ProtoReflect.Descriptor instead.
func (*CreateEnvironmentCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *CreateEnvironmentCommand) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Deprecated: Do not use.
func (x *CreateEnvironmentCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEnvironmentCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEnvironmentCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEnvironmentCommand) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// deprecated
type RenameEnvironmentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RenameEnvironmentCommand) Reset() {
	*x = RenameEnvironmentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameEnvironmentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameEnvironmentCommand) ProtoMessage() {}

func (x *RenameEnvironmentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameEnvironmentCommand.ProtoReflect.Descriptor instead.
func (*RenameEnvironmentCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{1}
}

func (x *RenameEnvironmentCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangeDescriptionEnvironmentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChangeDescriptionEnvironmentCommand) Reset() {
	*x = ChangeDescriptionEnvironmentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDescriptionEnvironmentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDescriptionEnvironmentCommand) ProtoMessage() {}

func (x *ChangeDescriptionEnvironmentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDescriptionEnvironmentCommand.ProtoReflect.Descriptor instead.
func (*ChangeDescriptionEnvironmentCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeDescriptionEnvironmentCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteEnvironmentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEnvironmentCommand) Reset() {
	*x = DeleteEnvironmentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEnvironmentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEnvironmentCommand) ProtoMessage() {}

func (x *DeleteEnvironmentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEnvironmentCommand.ProtoReflect.Descriptor instead.
func (*DeleteEnvironmentCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{3}
}

type CreateProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // optional
}

func (x *CreateProjectCommand) Reset() {
	*x = CreateProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectCommand) ProtoMessage() {}

func (x *CreateProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectCommand.ProtoReflect.Descriptor instead.
func (*CreateProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProjectCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProjectCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateTrialProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateTrialProjectCommand) Reset() {
	*x = CreateTrialProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrialProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrialProjectCommand) ProtoMessage() {}

func (x *CreateTrialProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrialProjectCommand.ProtoReflect.Descriptor instead.
func (*CreateTrialProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTrialProjectCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTrialProjectCommand) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ChangeDescriptionProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChangeDescriptionProjectCommand) Reset() {
	*x = ChangeDescriptionProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDescriptionProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDescriptionProjectCommand) ProtoMessage() {}

func (x *ChangeDescriptionProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDescriptionProjectCommand.ProtoReflect.Descriptor instead.
func (*ChangeDescriptionProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeDescriptionProjectCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EnableProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnableProjectCommand) Reset() {
	*x = EnableProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableProjectCommand) ProtoMessage() {}

func (x *EnableProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableProjectCommand.ProtoReflect.Descriptor instead.
func (*EnableProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{7}
}

type DisableProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisableProjectCommand) Reset() {
	*x = DisableProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableProjectCommand) ProtoMessage() {}

func (x *DisableProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableProjectCommand.ProtoReflect.Descriptor instead.
func (*DisableProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{8}
}

type ConvertTrialProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConvertTrialProjectCommand) Reset() {
	*x = ConvertTrialProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_environment_command_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertTrialProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertTrialProjectCommand) ProtoMessage() {}

func (x *ConvertTrialProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_environment_command_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertTrialProjectCommand.ProtoReflect.Descriptor instead.
func (*ConvertTrialProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_environment_command_proto_rawDescGZIP(), []int{9}
}

var File_proto_environment_command_proto protoreflect.FileDescriptor

var file_proto_environment_command_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x18, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x47, 0x0a, 0x23, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x41, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x43, 0x0a, 0x1f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69,
	0x6f, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_environment_command_proto_rawDescOnce sync.Once
	file_proto_environment_command_proto_rawDescData = file_proto_environment_command_proto_rawDesc
)

func file_proto_environment_command_proto_rawDescGZIP() []byte {
	file_proto_environment_command_proto_rawDescOnce.Do(func() {
		file_proto_environment_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_environment_command_proto_rawDescData)
	})
	return file_proto_environment_command_proto_rawDescData
}

var file_proto_environment_command_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_environment_command_proto_goTypes = []interface{}{
	(*CreateEnvironmentCommand)(nil),            // 0: bucketeer.environment.CreateEnvironmentCommand
	(*RenameEnvironmentCommand)(nil),            // 1: bucketeer.environment.RenameEnvironmentCommand
	(*ChangeDescriptionEnvironmentCommand)(nil), // 2: bucketeer.environment.ChangeDescriptionEnvironmentCommand
	(*DeleteEnvironmentCommand)(nil),            // 3: bucketeer.environment.DeleteEnvironmentCommand
	(*CreateProjectCommand)(nil),                // 4: bucketeer.environment.CreateProjectCommand
	(*CreateTrialProjectCommand)(nil),           // 5: bucketeer.environment.CreateTrialProjectCommand
	(*ChangeDescriptionProjectCommand)(nil),     // 6: bucketeer.environment.ChangeDescriptionProjectCommand
	(*EnableProjectCommand)(nil),                // 7: bucketeer.environment.EnableProjectCommand
	(*DisableProjectCommand)(nil),               // 8: bucketeer.environment.DisableProjectCommand
	(*ConvertTrialProjectCommand)(nil),          // 9: bucketeer.environment.ConvertTrialProjectCommand
}
var file_proto_environment_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_environment_command_proto_init() }
func file_proto_environment_command_proto_init() {
	if File_proto_environment_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_environment_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnvironmentCommand); i {
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
		file_proto_environment_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameEnvironmentCommand); i {
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
		file_proto_environment_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDescriptionEnvironmentCommand); i {
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
		file_proto_environment_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEnvironmentCommand); i {
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
		file_proto_environment_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectCommand); i {
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
		file_proto_environment_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTrialProjectCommand); i {
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
		file_proto_environment_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDescriptionProjectCommand); i {
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
		file_proto_environment_command_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableProjectCommand); i {
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
		file_proto_environment_command_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableProjectCommand); i {
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
		file_proto_environment_command_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertTrialProjectCommand); i {
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
			RawDescriptor: file_proto_environment_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_environment_command_proto_goTypes,
		DependencyIndexes: file_proto_environment_command_proto_depIdxs,
		MessageInfos:      file_proto_environment_command_proto_msgTypes,
	}.Build()
	File_proto_environment_command_proto = out.File
	file_proto_environment_command_proto_rawDesc = nil
	file_proto_environment_command_proto_goTypes = nil
	file_proto_environment_command_proto_depIdxs = nil
}