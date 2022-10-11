// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: com/teachingstrategies/temporal/v1beta1/api.proto

package temporalv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TemporalServiceClient is the client API for TemporalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemporalServiceClient interface {
	// CreateLicense starts a Temporal Workflow that will:
	// 1. Create an Org
	// 2. Create a Profile
	// 3. Create a License
	CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*CreateLicenseResponse, error)
}

type temporalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemporalServiceClient(cc grpc.ClientConnInterface) TemporalServiceClient {
	return &temporalServiceClient{cc}
}

func (c *temporalServiceClient) CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*CreateLicenseResponse, error) {
	out := new(CreateLicenseResponse)
	err := c.cc.Invoke(ctx, "/com.teachingstrategies.temporal.v1beta1.TemporalService/CreateLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemporalServiceServer is the server API for TemporalService service.
// All implementations should embed UnimplementedTemporalServiceServer
// for forward compatibility
type TemporalServiceServer interface {
	// CreateLicense starts a Temporal Workflow that will:
	// 1. Create an Org
	// 2. Create a Profile
	// 3. Create a License
	CreateLicense(context.Context, *CreateLicenseRequest) (*CreateLicenseResponse, error)
}

// UnimplementedTemporalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTemporalServiceServer struct {
}

func (UnimplementedTemporalServiceServer) CreateLicense(context.Context, *CreateLicenseRequest) (*CreateLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLicense not implemented")
}

// UnsafeTemporalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemporalServiceServer will
// result in compilation errors.
type UnsafeTemporalServiceServer interface {
	mustEmbedUnimplementedTemporalServiceServer()
}

func RegisterTemporalServiceServer(s grpc.ServiceRegistrar, srv TemporalServiceServer) {
	s.RegisterService(&TemporalService_ServiceDesc, srv)
}

func _TemporalService_CreateLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalServiceServer).CreateLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.teachingstrategies.temporal.v1beta1.TemporalService/CreateLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalServiceServer).CreateLicense(ctx, req.(*CreateLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemporalService_ServiceDesc is the grpc.ServiceDesc for TemporalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemporalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.teachingstrategies.temporal.v1beta1.TemporalService",
	HandlerType: (*TemporalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLicense",
			Handler:    _TemporalService_CreateLicense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/teachingstrategies/temporal/v1beta1/api.proto",
}
