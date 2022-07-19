// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: example.proto

package example

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

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseOk, error)
	PingWithDelay(ctx context.Context, in *Delay, opts ...grpc.CallOption) (*ResponseOk, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseOk, error) {
	out := new(ResponseOk)
	err := c.cc.Invoke(ctx, "/example.Example/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) PingWithDelay(ctx context.Context, in *Delay, opts ...grpc.CallOption) (*ResponseOk, error) {
	out := new(ResponseOk)
	err := c.cc.Invoke(ctx, "/example.Example/PingWithDelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
// All implementations should embed UnimplementedExampleServer
// for forward compatibility
type ExampleServer interface {
	Ping(context.Context, *Empty) (*ResponseOk, error)
	PingWithDelay(context.Context, *Delay) (*ResponseOk, error)
}

// UnimplementedExampleServer should be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (UnimplementedExampleServer) Ping(context.Context, *Empty) (*ResponseOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedExampleServer) PingWithDelay(context.Context, *Delay) (*ResponseOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingWithDelay not implemented")
}

// UnsafeExampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServer will
// result in compilation errors.
type UnsafeExampleServer interface {
	mustEmbedUnimplementedExampleServer()
}

func RegisterExampleServer(s grpc.ServiceRegistrar, srv ExampleServer) {
	s.RegisterService(&Example_ServiceDesc, srv)
}

func _Example_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Example/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_PingWithDelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).PingWithDelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Example/PingWithDelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).PingWithDelay(ctx, req.(*Delay))
	}
	return interceptor(ctx, in, info, handler)
}

// Example_ServiceDesc is the grpc.ServiceDesc for Example service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Example_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Example_Ping_Handler,
		},
		{
			MethodName: "PingWithDelay",
			Handler:    _Example_PingWithDelay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
