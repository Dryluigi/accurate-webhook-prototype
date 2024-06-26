// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: internal/webhook/api/forward.proto

package pb

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

// WebhookForwardServiceClient is the client API for WebhookForwardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookForwardServiceClient interface {
	Forward(ctx context.Context, opts ...grpc.CallOption) (WebhookForwardService_ForwardClient, error)
}

type webhookForwardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookForwardServiceClient(cc grpc.ClientConnInterface) WebhookForwardServiceClient {
	return &webhookForwardServiceClient{cc}
}

func (c *webhookForwardServiceClient) Forward(ctx context.Context, opts ...grpc.CallOption) (WebhookForwardService_ForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &WebhookForwardService_ServiceDesc.Streams[0], "/chat.WebhookForwardService/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &webhookForwardServiceForwardClient{stream}
	return x, nil
}

type WebhookForwardService_ForwardClient interface {
	Send(*WebhookForwardData) error
	Recv() (*WebhookForwardResponse, error)
	grpc.ClientStream
}

type webhookForwardServiceForwardClient struct {
	grpc.ClientStream
}

func (x *webhookForwardServiceForwardClient) Send(m *WebhookForwardData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *webhookForwardServiceForwardClient) Recv() (*WebhookForwardResponse, error) {
	m := new(WebhookForwardResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebhookForwardServiceServer is the server API for WebhookForwardService service.
// All implementations must embed UnimplementedWebhookForwardServiceServer
// for forward compatibility
type WebhookForwardServiceServer interface {
	Forward(WebhookForwardService_ForwardServer) error
	mustEmbedUnimplementedWebhookForwardServiceServer()
}

// UnimplementedWebhookForwardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookForwardServiceServer struct {
}

func (UnimplementedWebhookForwardServiceServer) Forward(WebhookForwardService_ForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (UnimplementedWebhookForwardServiceServer) mustEmbedUnimplementedWebhookForwardServiceServer() {}

// UnsafeWebhookForwardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookForwardServiceServer will
// result in compilation errors.
type UnsafeWebhookForwardServiceServer interface {
	mustEmbedUnimplementedWebhookForwardServiceServer()
}

func RegisterWebhookForwardServiceServer(s grpc.ServiceRegistrar, srv WebhookForwardServiceServer) {
	s.RegisterService(&WebhookForwardService_ServiceDesc, srv)
}

func _WebhookForwardService_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WebhookForwardServiceServer).Forward(&webhookForwardServiceForwardServer{stream})
}

type WebhookForwardService_ForwardServer interface {
	Send(*WebhookForwardResponse) error
	Recv() (*WebhookForwardData, error)
	grpc.ServerStream
}

type webhookForwardServiceForwardServer struct {
	grpc.ServerStream
}

func (x *webhookForwardServiceForwardServer) Send(m *WebhookForwardResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *webhookForwardServiceForwardServer) Recv() (*WebhookForwardData, error) {
	m := new(WebhookForwardData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebhookForwardService_ServiceDesc is the grpc.ServiceDesc for WebhookForwardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebhookForwardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.WebhookForwardService",
	HandlerType: (*WebhookForwardServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Forward",
			Handler:       _WebhookForwardService_Forward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/webhook/api/forward.proto",
}
