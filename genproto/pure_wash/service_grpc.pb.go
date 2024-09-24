// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.0
// source: service.proto

package pure_wash

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceServiceClient interface {
	CreateService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error)
	UpdateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	DeleteService(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetService(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Service, error)
	GetAllService(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ServicesResponse, error)
}

type serviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) CreateService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/carpet_wash_service.ServiceService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) UpdateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/carpet_wash_service.ServiceService/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) DeleteService(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/carpet_wash_service.ServiceService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) GetService(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/carpet_wash_service.ServiceService/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) GetAllService(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ServicesResponse, error) {
	out := new(ServicesResponse)
	err := c.cc.Invoke(ctx, "/carpet_wash_service.ServiceService/GetAllService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
// All implementations must embed UnimplementedServiceServiceServer
// for forward compatibility
type ServiceServiceServer interface {
	CreateService(context.Context, *ServiceRequest) (*Service, error)
	UpdateService(context.Context, *Service) (*Service, error)
	DeleteService(context.Context, *PrimaryKey) (*emptypb.Empty, error)
	GetService(context.Context, *PrimaryKey) (*Service, error)
	GetAllService(context.Context, *GetListRequest) (*ServicesResponse, error)
	mustEmbedUnimplementedServiceServiceServer()
}

// UnimplementedServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (UnimplementedServiceServiceServer) CreateService(context.Context, *ServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceServiceServer) UpdateService(context.Context, *Service) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedServiceServiceServer) DeleteService(context.Context, *PrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServiceServiceServer) GetService(context.Context, *PrimaryKey) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServiceServiceServer) GetAllService(context.Context, *GetListRequest) (*ServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllService not implemented")
}
func (UnimplementedServiceServiceServer) mustEmbedUnimplementedServiceServiceServer() {}

// UnsafeServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServiceServer will
// result in compilation errors.
type UnsafeServiceServiceServer interface {
	mustEmbedUnimplementedServiceServiceServer()
}

func RegisterServiceServiceServer(s grpc.ServiceRegistrar, srv ServiceServiceServer) {
	s.RegisterService(&ServiceService_ServiceDesc, srv)
}

func _ServiceService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carpet_wash_service.ServiceService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).CreateService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carpet_wash_service.ServiceService/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).UpdateService(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carpet_wash_service.ServiceService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).DeleteService(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carpet_wash_service.ServiceService/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).GetService(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_GetAllService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).GetAllService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carpet_wash_service.ServiceService/GetAllService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).GetAllService(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceService_ServiceDesc is the grpc.ServiceDesc for ServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "carpet_wash_service.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _ServiceService_CreateService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _ServiceService_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServiceService_DeleteService_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServiceService_GetService_Handler,
		},
		{
			MethodName: "GetAllService",
			Handler:    _ServiceService_GetAllService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
