// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.1
// source: dropbox/proto/build_tools/svclib/service.proto

package svclib

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

// SvcCtlClient is the client API for SvcCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcCtlClient interface {
	// Create multiple unstarted services.
	CreateBatch(ctx context.Context, in *CreateBatchReq, opts ...grpc.CallOption) (*Empty, error)
	// Stop (if necessary) and remove multiple services, not in any particular order.
	RemoveBatch(ctx context.Context, in *RemoveBatchReq, opts ...grpc.CallOption) (*Empty, error)
	// Start batch of created services. Attempting to start a running service
	// will be a no-op.
	Start(ctx context.Context, opts ...grpc.CallOption) (SvcCtl_StartClient, error)
	// Stop a service. Stopping an already stopped service is a no-op.
	Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*Empty, error)
	// Stop all services in reversed dependency order.
	StopAll(ctx context.Context, in *StopAllReq, opts ...grpc.CallOption) (*Empty, error)
	// Get status of selected services (or all services if empty).
	Status(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*StatusResp, error)
	// Get detailed diagnostics, including metrics (like cpu time) of selected services (or all services if empty).
	// This call is expensive and can take close to 1s to execute for large service graphs, so use this sparingly
	Diagnostics(ctx context.Context, in *DiagnosticsReq, opts ...grpc.CallOption) (*DiagnosticsResp, error)
}

type svcCtlClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcCtlClient(cc grpc.ClientConnInterface) SvcCtlClient {
	return &svcCtlClient{cc}
}

func (c *svcCtlClient) CreateBatch(ctx context.Context, in *CreateBatchReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/CreateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcCtlClient) RemoveBatch(ctx context.Context, in *RemoveBatchReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/RemoveBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcCtlClient) Start(ctx context.Context, opts ...grpc.CallOption) (SvcCtl_StartClient, error) {
	stream, err := c.cc.NewStream(ctx, &SvcCtl_ServiceDesc.Streams[0], "/build_tools.SvcCtl/Start", opts...)
	if err != nil {
		return nil, err
	}
	x := &svcCtlStartClient{stream}
	return x, nil
}

type SvcCtl_StartClient interface {
	Send(*StartReq) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type svcCtlStartClient struct {
	grpc.ClientStream
}

func (x *svcCtlStartClient) Send(m *StartReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *svcCtlStartClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *svcCtlClient) Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcCtlClient) StopAll(ctx context.Context, in *StopAllReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/StopAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcCtlClient) Status(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*StatusResp, error) {
	out := new(StatusResp)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcCtlClient) Diagnostics(ctx context.Context, in *DiagnosticsReq, opts ...grpc.CallOption) (*DiagnosticsResp, error) {
	out := new(DiagnosticsResp)
	err := c.cc.Invoke(ctx, "/build_tools.SvcCtl/Diagnostics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcCtlServer is the server API for SvcCtl service.
// All implementations must embed UnimplementedSvcCtlServer
// for forward compatibility
type SvcCtlServer interface {
	// Create multiple unstarted services.
	CreateBatch(context.Context, *CreateBatchReq) (*Empty, error)
	// Stop (if necessary) and remove multiple services, not in any particular order.
	RemoveBatch(context.Context, *RemoveBatchReq) (*Empty, error)
	// Start batch of created services. Attempting to start a running service
	// will be a no-op.
	Start(SvcCtl_StartServer) error
	// Stop a service. Stopping an already stopped service is a no-op.
	Stop(context.Context, *StopReq) (*Empty, error)
	// Stop all services in reversed dependency order.
	StopAll(context.Context, *StopAllReq) (*Empty, error)
	// Get status of selected services (or all services if empty).
	Status(context.Context, *StatusReq) (*StatusResp, error)
	// Get detailed diagnostics, including metrics (like cpu time) of selected services (or all services if empty).
	// This call is expensive and can take close to 1s to execute for large service graphs, so use this sparingly
	Diagnostics(context.Context, *DiagnosticsReq) (*DiagnosticsResp, error)
	mustEmbedUnimplementedSvcCtlServer()
}

// UnimplementedSvcCtlServer must be embedded to have forward compatible implementations.
type UnimplementedSvcCtlServer struct {
}

func (UnimplementedSvcCtlServer) CreateBatch(context.Context, *CreateBatchReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatch not implemented")
}
func (UnimplementedSvcCtlServer) RemoveBatch(context.Context, *RemoveBatchReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBatch not implemented")
}
func (UnimplementedSvcCtlServer) Start(SvcCtl_StartServer) error {
	return status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSvcCtlServer) Stop(context.Context, *StopReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSvcCtlServer) StopAll(context.Context, *StopAllReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAll not implemented")
}
func (UnimplementedSvcCtlServer) Status(context.Context, *StatusReq) (*StatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedSvcCtlServer) Diagnostics(context.Context, *DiagnosticsReq) (*DiagnosticsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diagnostics not implemented")
}
func (UnimplementedSvcCtlServer) mustEmbedUnimplementedSvcCtlServer() {}

// UnsafeSvcCtlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SvcCtlServer will
// result in compilation errors.
type UnsafeSvcCtlServer interface {
	mustEmbedUnimplementedSvcCtlServer()
}

func RegisterSvcCtlServer(s grpc.ServiceRegistrar, srv SvcCtlServer) {
	s.RegisterService(&SvcCtl_ServiceDesc, srv)
}

func _SvcCtl_CreateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).CreateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/CreateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).CreateBatch(ctx, req.(*CreateBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcCtl_RemoveBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).RemoveBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/RemoveBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).RemoveBatch(ctx, req.(*RemoveBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcCtl_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SvcCtlServer).Start(&svcCtlStartServer{stream})
}

type SvcCtl_StartServer interface {
	SendAndClose(*Empty) error
	Recv() (*StartReq, error)
	grpc.ServerStream
}

type svcCtlStartServer struct {
	grpc.ServerStream
}

func (x *svcCtlStartServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *svcCtlStartServer) Recv() (*StartReq, error) {
	m := new(StartReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SvcCtl_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).Stop(ctx, req.(*StopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcCtl_StopAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).StopAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/StopAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).StopAll(ctx, req.(*StopAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcCtl_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).Status(ctx, req.(*StatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcCtl_Diagnostics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiagnosticsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcCtlServer).Diagnostics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build_tools.SvcCtl/Diagnostics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcCtlServer).Diagnostics(ctx, req.(*DiagnosticsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SvcCtl_ServiceDesc is the grpc.ServiceDesc for SvcCtl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SvcCtl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "build_tools.SvcCtl",
	HandlerType: (*SvcCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBatch",
			Handler:    _SvcCtl_CreateBatch_Handler,
		},
		{
			MethodName: "RemoveBatch",
			Handler:    _SvcCtl_RemoveBatch_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _SvcCtl_Stop_Handler,
		},
		{
			MethodName: "StopAll",
			Handler:    _SvcCtl_StopAll_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _SvcCtl_Status_Handler,
		},
		{
			MethodName: "Diagnostics",
			Handler:    _SvcCtl_Diagnostics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _SvcCtl_Start_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "dropbox/proto/build_tools/svclib/service.proto",
}