// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/jd.proto

package jdv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Jd_Check_FullMethodName = "/Jd/Check"
)

// JdClient is the client API for Jd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JdClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type jdClient struct {
	cc grpc.ClientConnInterface
}

func NewJdClient(cc grpc.ClientConnInterface) JdClient {
	return &jdClient{cc}
}

func (c *jdClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, Jd_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JdServer is the server API for Jd service.
// All implementations must embed UnimplementedJdServer
// for forward compatibility.
type JdServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	mustEmbedUnimplementedJdServer()
}

// UnimplementedJdServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJdServer struct{}

func (UnimplementedJdServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedJdServer) mustEmbedUnimplementedJdServer() {}
func (UnimplementedJdServer) testEmbeddedByValue()            {}

// UnsafeJdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JdServer will
// result in compilation errors.
type UnsafeJdServer interface {
	mustEmbedUnimplementedJdServer()
}

func RegisterJdServer(s grpc.ServiceRegistrar, srv JdServer) {
	// If the following call pancis, it indicates UnimplementedJdServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Jd_ServiceDesc, srv)
}

func _Jd_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JdServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jd_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JdServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jd_ServiceDesc is the grpc.ServiceDesc for Jd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Jd",
	HandlerType: (*JdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Jd_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/jd.proto",
}
