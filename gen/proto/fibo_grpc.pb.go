// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FiboApiClient is the client API for FiboApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FiboApiClient interface {
	CalculateList(ctx context.Context, in *FiboRequest, opts ...grpc.CallOption) (*FiboResponse, error)
}

type fiboApiClient struct {
	cc grpc.ClientConnInterface
}

func NewFiboApiClient(cc grpc.ClientConnInterface) FiboApiClient {
	return &fiboApiClient{cc}
}

func (c *fiboApiClient) CalculateList(ctx context.Context, in *FiboRequest, opts ...grpc.CallOption) (*FiboResponse, error) {
	out := new(FiboResponse)
	err := c.cc.Invoke(ctx, "/main.FiboApi/CalculateList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiboApiServer is the server API for FiboApi service.
// All implementations must embed UnimplementedFiboApiServer
// for forward compatibility
type FiboApiServer interface {
	CalculateList(context.Context, *FiboRequest) (*FiboResponse, error)
	mustEmbedUnimplementedFiboApiServer()
}

// UnimplementedFiboApiServer must be embedded to have forward compatible implementations.
type UnimplementedFiboApiServer struct {
}

func (UnimplementedFiboApiServer) CalculateList(context.Context, *FiboRequest) (*FiboResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateList not implemented")
}
func (UnimplementedFiboApiServer) mustEmbedUnimplementedFiboApiServer() {}

// UnsafeFiboApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FiboApiServer will
// result in compilation errors.
type UnsafeFiboApiServer interface {
	mustEmbedUnimplementedFiboApiServer()
}

func RegisterFiboApiServer(s grpc.ServiceRegistrar, srv FiboApiServer) {
	s.RegisterService(&FiboApi_ServiceDesc, srv)
}

func _FiboApi_CalculateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiboRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiboApiServer).CalculateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.FiboApi/CalculateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiboApiServer).CalculateList(ctx, req.(*FiboRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FiboApi_ServiceDesc is the grpc.ServiceDesc for FiboApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FiboApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.FiboApi",
	HandlerType: (*FiboApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateList",
			Handler:    _FiboApi_CalculateList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fibo.proto",
}
