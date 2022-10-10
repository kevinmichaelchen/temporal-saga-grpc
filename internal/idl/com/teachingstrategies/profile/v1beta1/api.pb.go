// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: com/teachingstrategies/profile/v1beta1/api.proto

package profilev1beta1

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

type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateProfileResponse) Reset() {
	*x = CreateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileResponse) ProtoMessage() {}

func (x *CreateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileResponse.ProtoReflect.Descriptor instead.
func (*CreateProfileResponse) Descriptor() ([]byte, []int) {
	return file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescGZIP(), []int{1}
}

var File_com_teachingstrategies_profile_v1beta1_api_proto protoreflect.FileDescriptor

var file_com_teachingstrategies_profile_v1beta1_api_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xa1, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xe4, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e,
	0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x73, 0x61, 0x67, 0x61, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x50, 0xaa, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x2e, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x5c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x43, 0x6f, 0x6d, 0x5c,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x29, 0x43, 0x6f, 0x6d, 0x3a, 0x3a, 0x54, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescOnce sync.Once
	file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescData = file_com_teachingstrategies_profile_v1beta1_api_proto_rawDesc
)

func file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescGZIP() []byte {
	file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescOnce.Do(func() {
		file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescData)
	})
	return file_com_teachingstrategies_profile_v1beta1_api_proto_rawDescData
}

var file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_teachingstrategies_profile_v1beta1_api_proto_goTypes = []interface{}{
	(*CreateProfileRequest)(nil),  // 0: com.teachingstrategies.profile.v1beta1.CreateProfileRequest
	(*CreateProfileResponse)(nil), // 1: com.teachingstrategies.profile.v1beta1.CreateProfileResponse
}
var file_com_teachingstrategies_profile_v1beta1_api_proto_depIdxs = []int32{
	0, // 0: com.teachingstrategies.profile.v1beta1.ProfileService.CreateProfile:input_type -> com.teachingstrategies.profile.v1beta1.CreateProfileRequest
	1, // 1: com.teachingstrategies.profile.v1beta1.ProfileService.CreateProfile:output_type -> com.teachingstrategies.profile.v1beta1.CreateProfileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_teachingstrategies_profile_v1beta1_api_proto_init() }
func file_com_teachingstrategies_profile_v1beta1_api_proto_init() {
	if File_com_teachingstrategies_profile_v1beta1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileResponse); i {
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
			RawDescriptor: file_com_teachingstrategies_profile_v1beta1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_teachingstrategies_profile_v1beta1_api_proto_goTypes,
		DependencyIndexes: file_com_teachingstrategies_profile_v1beta1_api_proto_depIdxs,
		MessageInfos:      file_com_teachingstrategies_profile_v1beta1_api_proto_msgTypes,
	}.Build()
	File_com_teachingstrategies_profile_v1beta1_api_proto = out.File
	file_com_teachingstrategies_profile_v1beta1_api_proto_rawDesc = nil
	file_com_teachingstrategies_profile_v1beta1_api_proto_goTypes = nil
	file_com_teachingstrategies_profile_v1beta1_api_proto_depIdxs = nil
}
