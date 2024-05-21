// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: noolingo/deck.proto

package noolingo

import (
	context "context"
	common "github.com/noolingo/proto/codegen/go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Deck_Create_FullMethodName     = "/noolingo.Deck/Create"
	Deck_Delete_FullMethodName     = "/noolingo.Deck/Delete"
	Deck_List_FullMethodName       = "/noolingo.Deck/List"
	Deck_Get_FullMethodName        = "/noolingo.Deck/Get"
	Deck_CardAdd_FullMethodName    = "/noolingo.Deck/CardAdd"
	Deck_CardDelete_FullMethodName = "/noolingo.Deck/CardDelete"
)

// DeckClient is the client API for Deck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeckClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*common.Response, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	CardAdd(ctx context.Context, in *CardAddRequest, opts ...grpc.CallOption) (*common.Response, error)
	CardDelete(ctx context.Context, in *CardDeleteRequest, opts ...grpc.CallOption) (*common.Response, error)
}

type deckClient struct {
	cc grpc.ClientConnInterface
}

func NewDeckClient(cc grpc.ClientConnInterface) DeckClient {
	return &deckClient{cc}
}

func (c *deckClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Deck_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Deck_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Deck_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Deck_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckClient) CardAdd(ctx context.Context, in *CardAddRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Deck_CardAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckClient) CardDelete(ctx context.Context, in *CardDeleteRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Deck_CardDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeckServer is the server API for Deck service.
// All implementations must embed UnimplementedDeckServer
// for forward compatibility
type DeckServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*common.Response, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	CardAdd(context.Context, *CardAddRequest) (*common.Response, error)
	CardDelete(context.Context, *CardDeleteRequest) (*common.Response, error)
	mustEmbedUnimplementedDeckServer()
}

// UnimplementedDeckServer must be embedded to have forward compatible implementations.
type UnimplementedDeckServer struct {
}

func (UnimplementedDeckServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDeckServer) Delete(context.Context, *DeleteRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDeckServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDeckServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeckServer) CardAdd(context.Context, *CardAddRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardAdd not implemented")
}
func (UnimplementedDeckServer) CardDelete(context.Context, *CardDeleteRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardDelete not implemented")
}
func (UnimplementedDeckServer) mustEmbedUnimplementedDeckServer() {}

// UnsafeDeckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeckServer will
// result in compilation errors.
type UnsafeDeckServer interface {
	mustEmbedUnimplementedDeckServer()
}

func RegisterDeckServer(s grpc.ServiceRegistrar, srv DeckServer) {
	s.RegisterService(&Deck_ServiceDesc, srv)
}

func _Deck_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deck_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deck_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deck_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deck_CardAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).CardAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_CardAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).CardAdd(ctx, req.(*CardAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deck_CardDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServer).CardDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deck_CardDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServer).CardDelete(ctx, req.(*CardDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deck_ServiceDesc is the grpc.ServiceDesc for Deck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noolingo.Deck",
	HandlerType: (*DeckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Deck_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Deck_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Deck_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Deck_Get_Handler,
		},
		{
			MethodName: "CardAdd",
			Handler:    _Deck_CardAdd_Handler,
		},
		{
			MethodName: "CardDelete",
			Handler:    _Deck_CardDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noolingo/deck.proto",
}
