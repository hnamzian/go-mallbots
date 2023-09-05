// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: depotpb/api.proto

package depotpb

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
	DepotService_CreateShoppingList_FullMethodName      = "/depotpb.DepotService/CreateShoppingList"
	DepotService_CancelShoppingList_FullMethodName      = "/depotpb.DepotService/CancelShoppingList"
	DepotService_CompleteShoppingList_FullMethodName    = "/depotpb.DepotService/CompleteShoppingList"
	DepotService_AssignBotToShoppingList_FullMethodName = "/depotpb.DepotService/AssignBotToShoppingList"
	DepotService_GetShoppingList_FullMethodName         = "/depotpb.DepotService/GetShoppingList"
)

// DepotServiceClient is the client API for DepotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepotServiceClient interface {
	CreateShoppingList(ctx context.Context, in *CreateShoppingListRequest, opts ...grpc.CallOption) (*CreateShoppingListResponse, error)
	CancelShoppingList(ctx context.Context, in *CancelShoppingListRequest, opts ...grpc.CallOption) (*CancelShoppingListResponse, error)
	CompleteShoppingList(ctx context.Context, in *CompleteShoppingListRequest, opts ...grpc.CallOption) (*CompleteShoppingListResponse, error)
	AssignBotToShoppingList(ctx context.Context, in *AssignBotToShoppingListRequest, opts ...grpc.CallOption) (*AssignBotToShoppingListResponse, error)
	GetShoppingList(ctx context.Context, in *GetShoppingListRequest, opts ...grpc.CallOption) (*GetShoppingListResponse, error)
}

type depotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepotServiceClient(cc grpc.ClientConnInterface) DepotServiceClient {
	return &depotServiceClient{cc}
}

func (c *depotServiceClient) CreateShoppingList(ctx context.Context, in *CreateShoppingListRequest, opts ...grpc.CallOption) (*CreateShoppingListResponse, error) {
	out := new(CreateShoppingListResponse)
	err := c.cc.Invoke(ctx, DepotService_CreateShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depotServiceClient) CancelShoppingList(ctx context.Context, in *CancelShoppingListRequest, opts ...grpc.CallOption) (*CancelShoppingListResponse, error) {
	out := new(CancelShoppingListResponse)
	err := c.cc.Invoke(ctx, DepotService_CancelShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depotServiceClient) CompleteShoppingList(ctx context.Context, in *CompleteShoppingListRequest, opts ...grpc.CallOption) (*CompleteShoppingListResponse, error) {
	out := new(CompleteShoppingListResponse)
	err := c.cc.Invoke(ctx, DepotService_CompleteShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depotServiceClient) AssignBotToShoppingList(ctx context.Context, in *AssignBotToShoppingListRequest, opts ...grpc.CallOption) (*AssignBotToShoppingListResponse, error) {
	out := new(AssignBotToShoppingListResponse)
	err := c.cc.Invoke(ctx, DepotService_AssignBotToShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depotServiceClient) GetShoppingList(ctx context.Context, in *GetShoppingListRequest, opts ...grpc.CallOption) (*GetShoppingListResponse, error) {
	out := new(GetShoppingListResponse)
	err := c.cc.Invoke(ctx, DepotService_GetShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepotServiceServer is the server API for DepotService service.
// All implementations must embed UnimplementedDepotServiceServer
// for forward compatibility
type DepotServiceServer interface {
	CreateShoppingList(context.Context, *CreateShoppingListRequest) (*CreateShoppingListResponse, error)
	CancelShoppingList(context.Context, *CancelShoppingListRequest) (*CancelShoppingListResponse, error)
	CompleteShoppingList(context.Context, *CompleteShoppingListRequest) (*CompleteShoppingListResponse, error)
	AssignBotToShoppingList(context.Context, *AssignBotToShoppingListRequest) (*AssignBotToShoppingListResponse, error)
	GetShoppingList(context.Context, *GetShoppingListRequest) (*GetShoppingListResponse, error)
	mustEmbedUnimplementedDepotServiceServer()
}

// UnimplementedDepotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepotServiceServer struct {
}

func (UnimplementedDepotServiceServer) CreateShoppingList(context.Context, *CreateShoppingListRequest) (*CreateShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShoppingList not implemented")
}
func (UnimplementedDepotServiceServer) CancelShoppingList(context.Context, *CancelShoppingListRequest) (*CancelShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelShoppingList not implemented")
}
func (UnimplementedDepotServiceServer) CompleteShoppingList(context.Context, *CompleteShoppingListRequest) (*CompleteShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteShoppingList not implemented")
}
func (UnimplementedDepotServiceServer) AssignBotToShoppingList(context.Context, *AssignBotToShoppingListRequest) (*AssignBotToShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignBotToShoppingList not implemented")
}
func (UnimplementedDepotServiceServer) GetShoppingList(context.Context, *GetShoppingListRequest) (*GetShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingList not implemented")
}
func (UnimplementedDepotServiceServer) mustEmbedUnimplementedDepotServiceServer() {}

// UnsafeDepotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepotServiceServer will
// result in compilation errors.
type UnsafeDepotServiceServer interface {
	mustEmbedUnimplementedDepotServiceServer()
}

func RegisterDepotServiceServer(s grpc.ServiceRegistrar, srv DepotServiceServer) {
	s.RegisterService(&DepotService_ServiceDesc, srv)
}

func _DepotService_CreateShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepotServiceServer).CreateShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepotService_CreateShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepotServiceServer).CreateShoppingList(ctx, req.(*CreateShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepotService_CancelShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepotServiceServer).CancelShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepotService_CancelShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepotServiceServer).CancelShoppingList(ctx, req.(*CancelShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepotService_CompleteShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepotServiceServer).CompleteShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepotService_CompleteShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepotServiceServer).CompleteShoppingList(ctx, req.(*CompleteShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepotService_AssignBotToShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignBotToShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepotServiceServer).AssignBotToShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepotService_AssignBotToShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepotServiceServer).AssignBotToShoppingList(ctx, req.(*AssignBotToShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepotService_GetShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepotServiceServer).GetShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepotService_GetShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepotServiceServer).GetShoppingList(ctx, req.(*GetShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepotService_ServiceDesc is the grpc.ServiceDesc for DepotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "depotpb.DepotService",
	HandlerType: (*DepotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShoppingList",
			Handler:    _DepotService_CreateShoppingList_Handler,
		},
		{
			MethodName: "CancelShoppingList",
			Handler:    _DepotService_CancelShoppingList_Handler,
		},
		{
			MethodName: "CompleteShoppingList",
			Handler:    _DepotService_CompleteShoppingList_Handler,
		},
		{
			MethodName: "AssignBotToShoppingList",
			Handler:    _DepotService_AssignBotToShoppingList_Handler,
		},
		{
			MethodName: "GetShoppingList",
			Handler:    _DepotService_GetShoppingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "depotpb/api.proto",
}
