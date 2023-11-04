// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: noolingo/statistic.proto

package noolingo

import (
	context "context"
	common "github.com/MelnikovNA/noolingoproto/codegen/go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Statistic_StatisticUpdate_FullMethodName   = "/noolingo.Statistic/StatisticUpdate"
	Statistic_StatisticList_FullMethodName     = "/noolingo.Statistic/StatisticList"
	Statistic_StatisticbyDeckID_FullMethodName = "/noolingo.Statistic/StatisticbyDeckID"
)

// StatisticClient is the client API for Statistic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticClient interface {
	StatisticUpdate(ctx context.Context, in *StatisticUpdataRequest, opts ...grpc.CallOption) (*common.Response, error)
	StatisticList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatisticResponse, error)
	StatisticbyDeckID(ctx context.Context, in *StatisticbyDeckIDRequest, opts ...grpc.CallOption) (*StatisticResponse, error)
}

type statisticClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticClient(cc grpc.ClientConnInterface) StatisticClient {
	return &statisticClient{cc}
}

func (c *statisticClient) StatisticUpdate(ctx context.Context, in *StatisticUpdataRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Statistic_StatisticUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticClient) StatisticList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatisticResponse, error) {
	out := new(StatisticResponse)
	err := c.cc.Invoke(ctx, Statistic_StatisticList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticClient) StatisticbyDeckID(ctx context.Context, in *StatisticbyDeckIDRequest, opts ...grpc.CallOption) (*StatisticResponse, error) {
	out := new(StatisticResponse)
	err := c.cc.Invoke(ctx, Statistic_StatisticbyDeckID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticServer is the server API for Statistic service.
// All implementations must embed UnimplementedStatisticServer
// for forward compatibility
type StatisticServer interface {
	StatisticUpdate(context.Context, *StatisticUpdataRequest) (*common.Response, error)
	StatisticList(context.Context, *emptypb.Empty) (*StatisticResponse, error)
	StatisticbyDeckID(context.Context, *StatisticbyDeckIDRequest) (*StatisticResponse, error)
	mustEmbedUnimplementedStatisticServer()
}

// UnimplementedStatisticServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticServer struct {
}

func (UnimplementedStatisticServer) StatisticUpdate(context.Context, *StatisticUpdataRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatisticUpdate not implemented")
}
func (UnimplementedStatisticServer) StatisticList(context.Context, *emptypb.Empty) (*StatisticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatisticList not implemented")
}
func (UnimplementedStatisticServer) StatisticbyDeckID(context.Context, *StatisticbyDeckIDRequest) (*StatisticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatisticbyDeckID not implemented")
}
func (UnimplementedStatisticServer) mustEmbedUnimplementedStatisticServer() {}

// UnsafeStatisticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticServer will
// result in compilation errors.
type UnsafeStatisticServer interface {
	mustEmbedUnimplementedStatisticServer()
}

func RegisterStatisticServer(s grpc.ServiceRegistrar, srv StatisticServer) {
	s.RegisterService(&Statistic_ServiceDesc, srv)
}

func _Statistic_StatisticUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticUpdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticServer).StatisticUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistic_StatisticUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticServer).StatisticUpdate(ctx, req.(*StatisticUpdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Statistic_StatisticList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticServer).StatisticList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistic_StatisticList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticServer).StatisticList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Statistic_StatisticbyDeckID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticbyDeckIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticServer).StatisticbyDeckID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistic_StatisticbyDeckID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticServer).StatisticbyDeckID(ctx, req.(*StatisticbyDeckIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Statistic_ServiceDesc is the grpc.ServiceDesc for Statistic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Statistic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noolingo.Statistic",
	HandlerType: (*StatisticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatisticUpdate",
			Handler:    _Statistic_StatisticUpdate_Handler,
		},
		{
			MethodName: "StatisticList",
			Handler:    _Statistic_StatisticList_Handler,
		},
		{
			MethodName: "StatisticbyDeckID",
			Handler:    _Statistic_StatisticbyDeckID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noolingo/statistic.proto",
}
