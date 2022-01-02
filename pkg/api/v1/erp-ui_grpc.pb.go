// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ErpUIClient is the client API for ErpUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErpUIClient interface {
	// ListOrganisationSpecs returns a list of Organisation(s) that can be started through the UI.
	ListOrganisationSpecs(ctx context.Context, in *ListOrganisationSpecsRequest, opts ...grpc.CallOption) (ErpUI_ListOrganisationSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type erpUIClient struct {
	cc grpc.ClientConnInterface
}

func NewErpUIClient(cc grpc.ClientConnInterface) ErpUIClient {
	return &erpUIClient{cc}
}

func (c *erpUIClient) ListOrganisationSpecs(ctx context.Context, in *ListOrganisationSpecsRequest, opts ...grpc.CallOption) (ErpUI_ListOrganisationSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ErpUI_ServiceDesc.Streams[0], "/v1.ErpUI/ListOrganisationSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &erpUIListOrganisationSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ErpUI_ListOrganisationSpecsClient interface {
	Recv() (*ListOrganisationSpecsResponse, error)
	grpc.ClientStream
}

type erpUIListOrganisationSpecsClient struct {
	grpc.ClientStream
}

func (x *erpUIListOrganisationSpecsClient) Recv() (*ListOrganisationSpecsResponse, error) {
	m := new(ListOrganisationSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *erpUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.ErpUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErpUIServer is the server API for ErpUI service.
// All implementations must embed UnimplementedErpUIServer
// for forward compatibility
type ErpUIServer interface {
	// ListOrganisationSpecs returns a list of Organisation(s) that can be started through the UI.
	ListOrganisationSpecs(*ListOrganisationSpecsRequest, ErpUI_ListOrganisationSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedErpUIServer()
}

// UnimplementedErpUIServer must be embedded to have forward compatible implementations.
type UnimplementedErpUIServer struct {
}

func (UnimplementedErpUIServer) ListOrganisationSpecs(*ListOrganisationSpecsRequest, ErpUI_ListOrganisationSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOrganisationSpecs not implemented")
}
func (UnimplementedErpUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedErpUIServer) mustEmbedUnimplementedErpUIServer() {}

// UnsafeErpUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErpUIServer will
// result in compilation errors.
type UnsafeErpUIServer interface {
	mustEmbedUnimplementedErpUIServer()
}

func RegisterErpUIServer(s grpc.ServiceRegistrar, srv ErpUIServer) {
	s.RegisterService(&ErpUI_ServiceDesc, srv)
}

func _ErpUI_ListOrganisationSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrganisationSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErpUIServer).ListOrganisationSpecs(m, &erpUIListOrganisationSpecsServer{stream})
}

type ErpUI_ListOrganisationSpecsServer interface {
	Send(*ListOrganisationSpecsResponse) error
	grpc.ServerStream
}

type erpUIListOrganisationSpecsServer struct {
	grpc.ServerStream
}

func (x *erpUIListOrganisationSpecsServer) Send(m *ListOrganisationSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ErpUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErpUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ErpUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErpUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ErpUI_ServiceDesc is the grpc.ServiceDesc for ErpUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ErpUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ErpUI",
	HandlerType: (*ErpUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _ErpUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOrganisationSpecs",
			Handler:       _ErpUI_ListOrganisationSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "erp-ui.proto",
}
