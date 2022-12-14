// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: grpc.proto

package driver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/AlexandrBurak/TaxiAppOrder/pkg/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResult, error)
	GetOrdersByUser(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*ListOfOrders, error)
	RateLastTrip(ctx context.Context, in *RateTripByUserRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResult, error) {
	out := new(OrderResult)
	err := c.cc.Invoke(ctx, "/api.UserService/OrderTaxi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOrdersByUser(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*ListOfOrders, error) {
	out := new(ListOfOrders)
	err := c.cc.Invoke(ctx, "/api.UserService/GetOrdersByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RateLastTrip(ctx context.Context, in *RateTripByUserRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/RateLastTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	OrderTaxi(context.Context, *OrderRequest) (*OrderResult, error)
	GetOrdersByUser(context.Context, *OrderRequest) (*ListOfOrders, error)
	RateLastTrip(context.Context, *RateTripByUserRequest) (*RateResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) OrderTaxi(context.Context, *OrderRequest) (*OrderResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderTaxi not implemented")
}
func (UnimplementedUserServiceServer) GetOrdersByUser(context.Context, *OrderRequest) (*ListOfOrders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByUser not implemented")
}
func (UnimplementedUserServiceServer) RateLastTrip(context.Context, *RateTripByUserRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLastTrip not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_OrderTaxi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).OrderTaxi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/OrderTaxi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).OrderTaxi(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOrdersByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOrdersByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/GetOrdersByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOrdersByUser(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RateLastTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateTripByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RateLastTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/RateLastTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RateLastTrip(ctx, req.(*RateTripByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderTaxi",
			Handler:    _UserService_OrderTaxi_Handler,
		},
		{
			MethodName: "GetOrdersByUser",
			Handler:    _UserService_GetOrdersByUser_Handler,
		},
		{
			MethodName: "RateLastTrip",
			Handler:    _UserService_RateLastTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverServiceClient interface {
	GetAllFreeDrivers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllFreeDrivers, error)
	GetAllDrivers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllDrivers, error)
	RateTripByDriver(ctx context.Context, in *RateTripByDriverRequest, opts ...grpc.CallOption) (*Error, error)
	GetAllOrdersOfDriver(ctx context.Context, in *DriversPhone, opts ...grpc.CallOption) (*AllOrders, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) GetAllFreeDrivers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllFreeDrivers, error) {
	out := new(AllFreeDrivers)
	err := c.cc.Invoke(ctx, "/api.DriverService/GetAllFreeDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetAllDrivers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllDrivers, error) {
	out := new(AllDrivers)
	err := c.cc.Invoke(ctx, "/api.DriverService/GetAllDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) RateTripByDriver(ctx context.Context, in *RateTripByDriverRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.DriverService/RateTripByDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetAllOrdersOfDriver(ctx context.Context, in *DriversPhone, opts ...grpc.CallOption) (*AllOrders, error) {
	out := new(AllOrders)
	err := c.cc.Invoke(ctx, "/api.DriverService/GetAllOrdersOfDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility
type DriverServiceServer interface {
	GetAllFreeDrivers(context.Context, *Request) (*AllFreeDrivers, error)
	GetAllDrivers(context.Context, *Request) (*AllDrivers, error)
	RateTripByDriver(context.Context, *RateTripByDriverRequest) (*Error, error)
	GetAllOrdersOfDriver(context.Context, *DriversPhone) (*AllOrders, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServiceServer struct {
}

func (UnimplementedDriverServiceServer) GetAllFreeDrivers(context.Context, *Request) (*AllFreeDrivers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFreeDrivers not implemented")
}
func (UnimplementedDriverServiceServer) GetAllDrivers(context.Context, *Request) (*AllDrivers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDrivers not implemented")
}
func (UnimplementedDriverServiceServer) RateTripByDriver(context.Context, *RateTripByDriverRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateTripByDriver not implemented")
}
func (UnimplementedDriverServiceServer) GetAllOrdersOfDriver(context.Context, *DriversPhone) (*AllOrders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrdersOfDriver not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv *api.DriverServer) {
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_GetAllFreeDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetAllFreeDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DriverService/GetAllFreeDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetAllFreeDrivers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetAllDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetAllDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DriverService/GetAllDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetAllDrivers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_RateTripByDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateTripByDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).RateTripByDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DriverService/RateTripByDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).RateTripByDriver(ctx, req.(*RateTripByDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetAllOrdersOfDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriversPhone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetAllOrdersOfDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DriverService/GetAllOrdersOfDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetAllOrdersOfDriver(ctx, req.(*DriversPhone))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllFreeDrivers",
			Handler:    _DriverService_GetAllFreeDrivers_Handler,
		},
		{
			MethodName: "GetAllDrivers",
			Handler:    _DriverService_GetAllDrivers_Handler,
		},
		{
			MethodName: "RateTripByDriver",
			Handler:    _DriverService_RateTripByDriver_Handler,
		},
		{
			MethodName: "GetAllOrdersOfDriver",
			Handler:    _DriverService_GetAllOrdersOfDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
