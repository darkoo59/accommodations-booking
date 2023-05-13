// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.4
// source: accommodation/accommodation-service.proto

package accommodation

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

const (
	AccommodationService_GetAllAccommodations_FullMethodName         = "/AccommodationService/GetAllAccommodations"
	AccommodationService_CreateAccommodation_FullMethodName          = "/AccommodationService/CreateAccommodation"
	AccommodationService_GetAllPeriodsByAccommodation_FullMethodName = "/AccommodationService/GetAllPeriodsByAccommodation"
	AccommodationService_CreatePeriod_FullMethodName                 = "/AccommodationService/CreatePeriod"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	GetAllAccommodations(ctx context.Context, in *AM_GetAllAccommodations_Request, opts ...grpc.CallOption) (*AM_GetAllAccommodations_Response, error)
	CreateAccommodation(ctx context.Context, in *AM_CreateAccommodation_Request, opts ...grpc.CallOption) (*AM_CreateAccommodation_Response, error)
	GetAllPeriodsByAccommodation(ctx context.Context, in *AM_GetAllPeriodsByAccommodation_Request, opts ...grpc.CallOption) (*AM_GetAllPeriodsByAccommodation_Response, error)
	CreatePeriod(ctx context.Context, in *AM_CreatePeriod_Request, opts ...grpc.CallOption) (*AM_CreatePeriod_Response, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) GetAllAccommodations(ctx context.Context, in *AM_GetAllAccommodations_Request, opts ...grpc.CallOption) (*AM_GetAllAccommodations_Response, error) {
	out := new(AM_GetAllAccommodations_Response)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *AM_CreateAccommodation_Request, opts ...grpc.CallOption) (*AM_CreateAccommodation_Response, error) {
	out := new(AM_CreateAccommodation_Response)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllPeriodsByAccommodation(ctx context.Context, in *AM_GetAllPeriodsByAccommodation_Request, opts ...grpc.CallOption) (*AM_GetAllPeriodsByAccommodation_Response, error) {
	out := new(AM_GetAllPeriodsByAccommodation_Response)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllPeriodsByAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreatePeriod(ctx context.Context, in *AM_CreatePeriod_Request, opts ...grpc.CallOption) (*AM_CreatePeriod_Response, error) {
	out := new(AM_CreatePeriod_Response)
	err := c.cc.Invoke(ctx, AccommodationService_CreatePeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	GetAllAccommodations(context.Context, *AM_GetAllAccommodations_Request) (*AM_GetAllAccommodations_Response, error)
	CreateAccommodation(context.Context, *AM_CreateAccommodation_Request) (*AM_CreateAccommodation_Response, error)
	GetAllPeriodsByAccommodation(context.Context, *AM_GetAllPeriodsByAccommodation_Request) (*AM_GetAllPeriodsByAccommodation_Response, error)
	CreatePeriod(context.Context, *AM_CreatePeriod_Request) (*AM_CreatePeriod_Response, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) GetAllAccommodations(context.Context, *AM_GetAllAccommodations_Request) (*AM_GetAllAccommodations_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *AM_CreateAccommodation_Request) (*AM_CreateAccommodation_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllPeriodsByAccommodation(context.Context, *AM_GetAllPeriodsByAccommodation_Request) (*AM_GetAllPeriodsByAccommodation_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPeriodsByAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreatePeriod(context.Context, *AM_CreatePeriod_Request) (*AM_CreatePeriod_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeriod not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_GetAllAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_GetAllAccommodations_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAccommodations(ctx, req.(*AM_GetAllAccommodations_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_CreateAccommodation_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*AM_CreateAccommodation_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllPeriodsByAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_GetAllPeriodsByAccommodation_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllPeriodsByAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllPeriodsByAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllPeriodsByAccommodation(ctx, req.(*AM_GetAllPeriodsByAccommodation_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreatePeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_CreatePeriod_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreatePeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreatePeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreatePeriod(ctx, req.(*AM_CreatePeriod_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAccommodations",
			Handler:    _AccommodationService_GetAllAccommodations_Handler,
		},
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
			MethodName: "GetAllPeriodsByAccommodation",
			Handler:    _AccommodationService_GetAllPeriodsByAccommodation_Handler,
		},
		{
			MethodName: "CreatePeriod",
			Handler:    _AccommodationService_CreatePeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation/accommodation-service.proto",
}
