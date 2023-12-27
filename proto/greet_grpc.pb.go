// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/greet.proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Hello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloServerStream(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_HelloServerStreamClient, error)
	HellowClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HellowClientStreamClient, error)
	HelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloBiDirectionalStreamClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Hello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreetService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) HelloServerStream(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_HelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet_service.GreetService/HelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_HelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) HellowClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HellowClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet_service.GreetService/HellowClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHellowClientStreamClient{stream}
	return x, nil
}

type GreetService_HellowClientStreamClient interface {
	Send(*NamesList) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceHellowClientStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceHellowClientStreamClient) Send(m *NamesList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceHellowClientStreamClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) HelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloBiDirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet_service.GreetService/HelloBiDirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloBiDirectionalStreamClient{stream}
	return x, nil
}

type GreetService_HelloBiDirectionalStreamClient interface {
	Send(*NamesList) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceHelloBiDirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloBiDirectionalStreamClient) Send(m *NamesList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceHelloBiDirectionalStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Hello(context.Context, *NoParam) (*HelloResponse, error)
	HelloServerStream(*NamesList, GreetService_HelloServerStreamServer) error
	HellowClientStream(GreetService_HellowClientStreamServer) error
	HelloBiDirectionalStream(GreetService_HelloBiDirectionalStreamServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Hello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetServiceServer) HelloServerStream(*NamesList, GreetService_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedGreetServiceServer) HellowClientStream(GreetService_HellowClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HellowClientStream not implemented")
}
func (UnimplementedGreetServiceServer) HelloBiDirectionalStream(GreetService_HelloBiDirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloBiDirectionalStream not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreetService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Hello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).HelloServerStream(m, &greetServiceHelloServerStreamServer{stream})
}

type GreetService_HelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_HellowClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).HellowClientStream(&greetServiceHellowClientStreamServer{stream})
}

type GreetService_HellowClientStreamServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*NamesList, error)
	grpc.ServerStream
}

type greetServiceHellowClientStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceHellowClientStreamServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceHellowClientStreamServer) Recv() (*NamesList, error) {
	m := new(NamesList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_HelloBiDirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).HelloBiDirectionalStream(&greetServiceHelloBiDirectionalStreamServer{stream})
}

type GreetService_HelloBiDirectionalStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*NamesList, error)
	grpc.ServerStream
}

type greetServiceHelloBiDirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloBiDirectionalStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceHelloBiDirectionalStreamServer) Recv() (*NamesList, error) {
	m := new(NamesList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _GreetService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HellowClientStream",
			Handler:       _GreetService_HellowClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloBiDirectionalStream",
			Handler:       _GreetService_HelloBiDirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}