// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: org/v1beta1/api.proto

package orgv1beta1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/org/v1beta1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// OrgServiceName is the fully-qualified name of the OrgService service.
	OrgServiceName = "org.v1beta1.OrgService"
)

// OrgServiceClient is a client for the org.v1beta1.OrgService service.
type OrgServiceClient interface {
	CreateOrg(context.Context, *connect_go.Request[v1beta1.CreateOrgRequest]) (*connect_go.Response[v1beta1.CreateOrgResponse], error)
}

// NewOrgServiceClient constructs a client for the org.v1beta1.OrgService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrgServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) OrgServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &orgServiceClient{
		createOrg: connect_go.NewClient[v1beta1.CreateOrgRequest, v1beta1.CreateOrgResponse](
			httpClient,
			baseURL+"/org.v1beta1.OrgService/CreateOrg",
			opts...,
		),
	}
}

// orgServiceClient implements OrgServiceClient.
type orgServiceClient struct {
	createOrg *connect_go.Client[v1beta1.CreateOrgRequest, v1beta1.CreateOrgResponse]
}

// CreateOrg calls org.v1beta1.OrgService.CreateOrg.
func (c *orgServiceClient) CreateOrg(ctx context.Context, req *connect_go.Request[v1beta1.CreateOrgRequest]) (*connect_go.Response[v1beta1.CreateOrgResponse], error) {
	return c.createOrg.CallUnary(ctx, req)
}

// OrgServiceHandler is an implementation of the org.v1beta1.OrgService service.
type OrgServiceHandler interface {
	CreateOrg(context.Context, *connect_go.Request[v1beta1.CreateOrgRequest]) (*connect_go.Response[v1beta1.CreateOrgResponse], error)
}

// NewOrgServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrgServiceHandler(svc OrgServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/org.v1beta1.OrgService/CreateOrg", connect_go.NewUnaryHandler(
		"/org.v1beta1.OrgService/CreateOrg",
		svc.CreateOrg,
		opts...,
	))
	return "/org.v1beta1.OrgService/", mux
}

// UnimplementedOrgServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrgServiceHandler struct{}

func (UnimplementedOrgServiceHandler) CreateOrg(context.Context, *connect_go.Request[v1beta1.CreateOrgRequest]) (*connect_go.Response[v1beta1.CreateOrgResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("org.v1beta1.OrgService.CreateOrg is not implemented"))
}
