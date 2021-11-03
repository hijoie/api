// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AppearenceServiceClient is the client API for AppearenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppearenceServiceClient interface {
	GetUAppearence(ctx context.Context, in *GetUAppearenceReq, opts ...grpc.CallOption) (*GetUAppearenceResp, error)
	//更新用户皮肤，未解锁的会购买解锁。
	UpdateUAppearence(ctx context.Context, in *UpdateUAppearenceReq, opts ...grpc.CallOption) (*UpdateUAppearenceResp, error)
}

type appearenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppearenceServiceClient(cc grpc.ClientConnInterface) AppearenceServiceClient {
	return &appearenceServiceClient{cc}
}

func (c *appearenceServiceClient) GetUAppearence(ctx context.Context, in *GetUAppearenceReq, opts ...grpc.CallOption) (*GetUAppearenceResp, error) {
	out := new(GetUAppearenceResp)
	err := c.cc.Invoke(ctx, "/daoying.AppearenceService/GetUAppearence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appearenceServiceClient) UpdateUAppearence(ctx context.Context, in *UpdateUAppearenceReq, opts ...grpc.CallOption) (*UpdateUAppearenceResp, error) {
	out := new(UpdateUAppearenceResp)
	err := c.cc.Invoke(ctx, "/daoying.AppearenceService/UpdateUAppearence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppearenceServiceServer is the server API for AppearenceService service.
// All implementations must embed UnimplementedAppearenceServiceServer
// for forward compatibility
type AppearenceServiceServer interface {
	GetUAppearence(context.Context, *GetUAppearenceReq) (*GetUAppearenceResp, error)
	//更新用户皮肤，未解锁的会购买解锁。
	UpdateUAppearence(context.Context, *UpdateUAppearenceReq) (*UpdateUAppearenceResp, error)
	mustEmbedUnimplementedAppearenceServiceServer()
}

// UnimplementedAppearenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppearenceServiceServer struct {
}

func (UnimplementedAppearenceServiceServer) GetUAppearence(context.Context, *GetUAppearenceReq) (*GetUAppearenceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUAppearence not implemented")
}
func (UnimplementedAppearenceServiceServer) UpdateUAppearence(context.Context, *UpdateUAppearenceReq) (*UpdateUAppearenceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUAppearence not implemented")
}
func (UnimplementedAppearenceServiceServer) mustEmbedUnimplementedAppearenceServiceServer() {}

// UnsafeAppearenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppearenceServiceServer will
// result in compilation errors.
type UnsafeAppearenceServiceServer interface {
	mustEmbedUnimplementedAppearenceServiceServer()
}

func RegisterAppearenceServiceServer(s grpc.ServiceRegistrar, srv AppearenceServiceServer) {
	s.RegisterService(&AppearenceService_ServiceDesc, srv)
}

func _AppearenceService_GetUAppearence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUAppearenceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppearenceServiceServer).GetUAppearence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daoying.AppearenceService/GetUAppearence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppearenceServiceServer).GetUAppearence(ctx, req.(*GetUAppearenceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppearenceService_UpdateUAppearence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUAppearenceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppearenceServiceServer).UpdateUAppearence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daoying.AppearenceService/UpdateUAppearence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppearenceServiceServer).UpdateUAppearence(ctx, req.(*UpdateUAppearenceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppearenceService_ServiceDesc is the grpc.ServiceDesc for AppearenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppearenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daoying.AppearenceService",
	HandlerType: (*AppearenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUAppearence",
			Handler:    _AppearenceService_GetUAppearence_Handler,
		},
		{
			MethodName: "UpdateUAppearence",
			Handler:    _AppearenceService_UpdateUAppearence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/daoying/pb/appearence.proto",
}