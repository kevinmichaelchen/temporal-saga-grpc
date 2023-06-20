// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: temporal/v1beta1/api.proto

package temporalv1beta1

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

type WorkflowOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID - The business identifier of the workflow execution.
	// Optional: defaulted to a uuid.
	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
}

func (x *WorkflowOptions) Reset() {
	*x = WorkflowOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowOptions) ProtoMessage() {}

func (x *WorkflowOptions) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowOptions.ProtoReflect.Descriptor instead.
func (*WorkflowOptions) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowOptions) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type CreateOnboardingWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowOptions *WorkflowOptions `protobuf:"bytes,1,opt,name=workflow_options,json=workflowOptions,proto3" json:"workflow_options,omitempty"`
	License         *License         `protobuf:"bytes,2,opt,name=license,proto3" json:"license,omitempty"`
	Org             *Org             `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	Profile         *Profile         `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateOnboardingWorkflowRequest) Reset() {
	*x = CreateOnboardingWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOnboardingWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOnboardingWorkflowRequest) ProtoMessage() {}

func (x *CreateOnboardingWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOnboardingWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CreateOnboardingWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOnboardingWorkflowRequest) GetWorkflowOptions() *WorkflowOptions {
	if x != nil {
		return x.WorkflowOptions
	}
	return nil
}

func (x *CreateOnboardingWorkflowRequest) GetLicense() *License {
	if x != nil {
		return x.License
	}
	return nil
}

func (x *CreateOnboardingWorkflowRequest) GetOrg() *Org {
	if x != nil {
		return x.Org
	}
	return nil
}

func (x *CreateOnboardingWorkflowRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type CreateOnboardingWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOnboardingWorkflowResponse) Reset() {
	*x = CreateOnboardingWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOnboardingWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOnboardingWorkflowResponse) ProtoMessage() {}

func (x *CreateOnboardingWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOnboardingWorkflowResponse.ProtoReflect.Descriptor instead.
func (*CreateOnboardingWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{2}
}

type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{3}
}

func (x *License) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{4}
}

func (x *Org) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1beta1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1beta1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_temporal_v1beta1_api_proto_rawDescGZIP(), []int{5}
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_temporal_v1beta1_api_proto protoreflect.FileDescriptor

var file_temporal_v1beta1_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x32,
	0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x49, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x6f, 0x72, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x03, 0x6f,
	0x72, 0x67, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x22, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x03, 0x4f, 0x72,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdf,
	0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x65, 0x76, 0x69, 0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x73, 0x61, 0x67, 0x61, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64, 0x6c,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1c,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_v1beta1_api_proto_rawDescOnce sync.Once
	file_temporal_v1beta1_api_proto_rawDescData = file_temporal_v1beta1_api_proto_rawDesc
)

func file_temporal_v1beta1_api_proto_rawDescGZIP() []byte {
	file_temporal_v1beta1_api_proto_rawDescOnce.Do(func() {
		file_temporal_v1beta1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_v1beta1_api_proto_rawDescData)
	})
	return file_temporal_v1beta1_api_proto_rawDescData
}

var file_temporal_v1beta1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_temporal_v1beta1_api_proto_goTypes = []interface{}{
	(*WorkflowOptions)(nil),                  // 0: temporal.v1beta1.WorkflowOptions
	(*CreateOnboardingWorkflowRequest)(nil),  // 1: temporal.v1beta1.CreateOnboardingWorkflowRequest
	(*CreateOnboardingWorkflowResponse)(nil), // 2: temporal.v1beta1.CreateOnboardingWorkflowResponse
	(*License)(nil),                          // 3: temporal.v1beta1.License
	(*Org)(nil),                              // 4: temporal.v1beta1.Org
	(*Profile)(nil),                          // 5: temporal.v1beta1.Profile
}
var file_temporal_v1beta1_api_proto_depIdxs = []int32{
	0, // 0: temporal.v1beta1.CreateOnboardingWorkflowRequest.workflow_options:type_name -> temporal.v1beta1.WorkflowOptions
	3, // 1: temporal.v1beta1.CreateOnboardingWorkflowRequest.license:type_name -> temporal.v1beta1.License
	4, // 2: temporal.v1beta1.CreateOnboardingWorkflowRequest.org:type_name -> temporal.v1beta1.Org
	5, // 3: temporal.v1beta1.CreateOnboardingWorkflowRequest.profile:type_name -> temporal.v1beta1.Profile
	1, // 4: temporal.v1beta1.TemporalService.CreateOnboardingWorkflow:input_type -> temporal.v1beta1.CreateOnboardingWorkflowRequest
	2, // 5: temporal.v1beta1.TemporalService.CreateOnboardingWorkflow:output_type -> temporal.v1beta1.CreateOnboardingWorkflowResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_temporal_v1beta1_api_proto_init() }
func file_temporal_v1beta1_api_proto_init() {
	if File_temporal_v1beta1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_v1beta1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowOptions); i {
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
		file_temporal_v1beta1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOnboardingWorkflowRequest); i {
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
		file_temporal_v1beta1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOnboardingWorkflowResponse); i {
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
		file_temporal_v1beta1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
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
		file_temporal_v1beta1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Org); i {
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
		file_temporal_v1beta1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_temporal_v1beta1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_temporal_v1beta1_api_proto_goTypes,
		DependencyIndexes: file_temporal_v1beta1_api_proto_depIdxs,
		MessageInfos:      file_temporal_v1beta1_api_proto_msgTypes,
	}.Build()
	File_temporal_v1beta1_api_proto = out.File
	file_temporal_v1beta1_api_proto_rawDesc = nil
	file_temporal_v1beta1_api_proto_goTypes = nil
	file_temporal_v1beta1_api_proto_depIdxs = nil
}
