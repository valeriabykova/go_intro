// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: project.proto

package dto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AwesomeProject_CreateAccount_FullMethodName = "/dto.AwesomeProject/CreateAccount"
	AwesomeProject_GetAccount_FullMethodName    = "/dto.AwesomeProject/GetAccount"
	AwesomeProject_SetBalance_FullMethodName    = "/dto.AwesomeProject/SetBalance"
	AwesomeProject_RenameAccount_FullMethodName = "/dto.AwesomeProject/RenameAccount"
	AwesomeProject_DeleteAccount_FullMethodName = "/dto.AwesomeProject/DeleteAccount"
)

// AwesomeProjectClient is the client API for AwesomeProject service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwesomeProjectClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	SetBalance(ctx context.Context, in *SetBalanceRequest, opts ...grpc.CallOption) (*Empty, error)
	RenameAccount(ctx context.Context, in *RenameAccountRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*Empty, error)
}

type awesomeProjectClient struct {
	cc grpc.ClientConnInterface
}

func NewAwesomeProjectClient(cc grpc.ClientConnInterface) AwesomeProjectClient {
	return &awesomeProjectClient{cc}
}

func (c *awesomeProjectClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AwesomeProject_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeProjectClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, AwesomeProject_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeProjectClient) SetBalance(ctx context.Context, in *SetBalanceRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AwesomeProject_SetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeProjectClient) RenameAccount(ctx context.Context, in *RenameAccountRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AwesomeProject_RenameAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeProjectClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AwesomeProject_DeleteAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwesomeProjectServer is the server API for AwesomeProject service.
// All implementations must embed UnimplementedAwesomeProjectServer
// for forward compatibility
type AwesomeProjectServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*Empty, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	SetBalance(context.Context, *SetBalanceRequest) (*Empty, error)
	RenameAccount(context.Context, *RenameAccountRequest) (*Empty, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*Empty, error)
	mustEmbedUnimplementedAwesomeProjectServer()
}

// UnimplementedAwesomeProjectServer must be embedded to have forward compatible implementations.
type UnimplementedAwesomeProjectServer struct {
}

func (UnimplementedAwesomeProjectServer) CreateAccount(context.Context, *CreateAccountRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAwesomeProjectServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAwesomeProjectServer) SetBalance(context.Context, *SetBalanceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBalance not implemented")
}
func (UnimplementedAwesomeProjectServer) RenameAccount(context.Context, *RenameAccountRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameAccount not implemented")
}
func (UnimplementedAwesomeProjectServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAwesomeProjectServer) mustEmbedUnimplementedAwesomeProjectServer() {}

// UnsafeAwesomeProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwesomeProjectServer will
// result in compilation errors.
type UnsafeAwesomeProjectServer interface {
	mustEmbedUnimplementedAwesomeProjectServer()
}

func RegisterAwesomeProjectServer(s grpc.ServiceRegistrar, srv AwesomeProjectServer) {
	s.RegisterService(&AwesomeProject_ServiceDesc, srv)
}

func _AwesomeProject_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeProjectServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwesomeProject_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeProjectServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeProject_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeProjectServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwesomeProject_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeProjectServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeProject_SetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeProjectServer).SetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwesomeProject_SetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeProjectServer).SetBalance(ctx, req.(*SetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeProject_RenameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeProjectServer).RenameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwesomeProject_RenameAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeProjectServer).RenameAccount(ctx, req.(*RenameAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeProject_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeProjectServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwesomeProject_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeProjectServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AwesomeProject_ServiceDesc is the grpc.ServiceDesc for AwesomeProject service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwesomeProject_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dto.AwesomeProject",
	HandlerType: (*AwesomeProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AwesomeProject_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AwesomeProject_GetAccount_Handler,
		},
		{
			MethodName: "SetBalance",
			Handler:    _AwesomeProject_SetBalance_Handler,
		},
		{
			MethodName: "RenameAccount",
			Handler:    _AwesomeProject_RenameAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AwesomeProject_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}
