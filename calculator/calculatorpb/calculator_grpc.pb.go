// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculatorpb

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	CalculateEveryone(ctx context.Context, opts ...grpc.CallOption) (CalculateService_CalculateEveryoneClient, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) CalculateEveryone(ctx context.Context, opts ...grpc.CallOption) (CalculateService_CalculateEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[0], "/calculator.CalculateService/CalculateEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceCalculateEveryoneClient{stream}
	return x, nil
}

type CalculateService_CalculateEveryoneClient interface {
	Send(*CalculatorRequest) error
	Recv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculateServiceCalculateEveryoneClient struct {
	grpc.ClientStream
}

func (x *calculateServiceCalculateEveryoneClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceCalculateEveryoneClient) Recv() (*CalculatorResponse, error) {
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations must embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	CalculateEveryone(CalculateService_CalculateEveryoneServer) error
	mustEmbedUnimplementedCalculateServiceServer()
}

// UnimplementedCalculateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) CalculateEveryone(CalculateService_CalculateEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateEveryone not implemented")
}
func (UnimplementedCalculateServiceServer) mustEmbedUnimplementedCalculateServiceServer() {}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_CalculateEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).CalculateEveryone(&calculateServiceCalculateEveryoneServer{stream})
}

type CalculateService_CalculateEveryoneServer interface {
	Send(*CalculatorResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculateServiceCalculateEveryoneServer struct {
	grpc.ServerStream
}

func (x *calculateServiceCalculateEveryoneServer) Send(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceCalculateEveryoneServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculateEveryone",
			Handler:       _CalculateService_CalculateEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
