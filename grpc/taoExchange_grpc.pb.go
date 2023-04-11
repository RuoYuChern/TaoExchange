// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.1
// source: taoExchange.proto

package grpc

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
	TaoExchangeSrv_QueryOrder_FullMethodName     = "/grpc.TaoExchangeSrv/queryOrder"
	TaoExchangeSrv_DoOrderCommond_FullMethodName = "/grpc.TaoExchangeSrv/doOrderCommond"
)

// TaoExchangeSrvClient is the client API for TaoExchangeSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaoExchangeSrvClient interface {
	QueryOrder(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error)
	DoOrderCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRsp, error)
}

type taoExchangeSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewTaoExchangeSrvClient(cc grpc.ClientConnInterface) TaoExchangeSrvClient {
	return &taoExchangeSrvClient{cc}
}

func (c *taoExchangeSrvClient) QueryOrder(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error) {
	out := new(QueryRsp)
	err := c.cc.Invoke(ctx, TaoExchangeSrv_QueryOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoExchangeSrvClient) DoOrderCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRsp, error) {
	out := new(OrderRsp)
	err := c.cc.Invoke(ctx, TaoExchangeSrv_DoOrderCommond_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaoExchangeSrvServer is the server API for TaoExchangeSrv service.
// All implementations must embed UnimplementedTaoExchangeSrvServer
// for forward compatibility
type TaoExchangeSrvServer interface {
	QueryOrder(context.Context, *QueryReq) (*QueryRsp, error)
	DoOrderCommond(context.Context, *OrderReq) (*OrderRsp, error)
	mustEmbedUnimplementedTaoExchangeSrvServer()
}

// UnimplementedTaoExchangeSrvServer must be embedded to have forward compatible implementations.
type UnimplementedTaoExchangeSrvServer struct {
}

func (UnimplementedTaoExchangeSrvServer) QueryOrder(context.Context, *QueryReq) (*QueryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (UnimplementedTaoExchangeSrvServer) DoOrderCommond(context.Context, *OrderReq) (*OrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoOrderCommond not implemented")
}
func (UnimplementedTaoExchangeSrvServer) mustEmbedUnimplementedTaoExchangeSrvServer() {}

// UnsafeTaoExchangeSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaoExchangeSrvServer will
// result in compilation errors.
type UnsafeTaoExchangeSrvServer interface {
	mustEmbedUnimplementedTaoExchangeSrvServer()
}

func RegisterTaoExchangeSrvServer(s grpc.ServiceRegistrar, srv TaoExchangeSrvServer) {
	s.RegisterService(&TaoExchangeSrv_ServiceDesc, srv)
}

func _TaoExchangeSrv_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoExchangeSrvServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoExchangeSrv_QueryOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoExchangeSrvServer).QueryOrder(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoExchangeSrv_DoOrderCommond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoExchangeSrvServer).DoOrderCommond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoExchangeSrv_DoOrderCommond_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoExchangeSrvServer).DoOrderCommond(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaoExchangeSrv_ServiceDesc is the grpc.ServiceDesc for TaoExchangeSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaoExchangeSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TaoExchangeSrv",
	HandlerType: (*TaoExchangeSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "queryOrder",
			Handler:    _TaoExchangeSrv_QueryOrder_Handler,
		},
		{
			MethodName: "doOrderCommond",
			Handler:    _TaoExchangeSrv_DoOrderCommond_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taoExchange.proto",
}

const (
	TaoHighStoreSrv_SaveCommond_FullMethodName = "/grpc.TaoHighStoreSrv/saveCommond"
	TaoHighStoreSrv_BatchGet_FullMethodName    = "/grpc.TaoHighStoreSrv/batchGet"
)

// TaoHighStoreSrvClient is the client API for TaoHighStoreSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaoHighStoreSrvClient interface {
	SaveCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*CommonRsp, error)
	BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error)
}

type taoHighStoreSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewTaoHighStoreSrvClient(cc grpc.ClientConnInterface) TaoHighStoreSrvClient {
	return &taoHighStoreSrvClient{cc}
}

func (c *taoHighStoreSrvClient) SaveCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, TaoHighStoreSrv_SaveCommond_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoHighStoreSrvClient) BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error) {
	out := new(BatchGetRsp)
	err := c.cc.Invoke(ctx, TaoHighStoreSrv_BatchGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaoHighStoreSrvServer is the server API for TaoHighStoreSrv service.
// All implementations must embed UnimplementedTaoHighStoreSrvServer
// for forward compatibility
type TaoHighStoreSrvServer interface {
	SaveCommond(context.Context, *OrderReq) (*CommonRsp, error)
	BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error)
	mustEmbedUnimplementedTaoHighStoreSrvServer()
}

// UnimplementedTaoHighStoreSrvServer must be embedded to have forward compatible implementations.
type UnimplementedTaoHighStoreSrvServer struct {
}

func (UnimplementedTaoHighStoreSrvServer) SaveCommond(context.Context, *OrderReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCommond not implemented")
}
func (UnimplementedTaoHighStoreSrvServer) BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedTaoHighStoreSrvServer) mustEmbedUnimplementedTaoHighStoreSrvServer() {}

// UnsafeTaoHighStoreSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaoHighStoreSrvServer will
// result in compilation errors.
type UnsafeTaoHighStoreSrvServer interface {
	mustEmbedUnimplementedTaoHighStoreSrvServer()
}

func RegisterTaoHighStoreSrvServer(s grpc.ServiceRegistrar, srv TaoHighStoreSrvServer) {
	s.RegisterService(&TaoHighStoreSrv_ServiceDesc, srv)
}

func _TaoHighStoreSrv_SaveCommond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoHighStoreSrvServer).SaveCommond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoHighStoreSrv_SaveCommond_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoHighStoreSrvServer).SaveCommond(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoHighStoreSrv_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoHighStoreSrvServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoHighStoreSrv_BatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoHighStoreSrvServer).BatchGet(ctx, req.(*BatchGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaoHighStoreSrv_ServiceDesc is the grpc.ServiceDesc for TaoHighStoreSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaoHighStoreSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TaoHighStoreSrv",
	HandlerType: (*TaoHighStoreSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveCommond",
			Handler:    _TaoHighStoreSrv_SaveCommond_Handler,
		},
		{
			MethodName: "batchGet",
			Handler:    _TaoHighStoreSrv_BatchGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taoExchange.proto",
}

const (
	TaoMarketDataSrv_BatchGet_FullMethodName   = "/grpc.TaoMarketDataSrv/batchGet"
	TaoMarketDataSrv_QueryOrder_FullMethodName = "/grpc.TaoMarketDataSrv/queryOrder"
)

// TaoMarketDataSrvClient is the client API for TaoMarketDataSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaoMarketDataSrvClient interface {
	BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error)
	QueryOrder(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error)
}

type taoMarketDataSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewTaoMarketDataSrvClient(cc grpc.ClientConnInterface) TaoMarketDataSrvClient {
	return &taoMarketDataSrvClient{cc}
}

func (c *taoMarketDataSrvClient) BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error) {
	out := new(BatchGetRsp)
	err := c.cc.Invoke(ctx, TaoMarketDataSrv_BatchGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoMarketDataSrvClient) QueryOrder(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error) {
	out := new(QueryRsp)
	err := c.cc.Invoke(ctx, TaoMarketDataSrv_QueryOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaoMarketDataSrvServer is the server API for TaoMarketDataSrv service.
// All implementations must embed UnimplementedTaoMarketDataSrvServer
// for forward compatibility
type TaoMarketDataSrvServer interface {
	BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error)
	QueryOrder(context.Context, *QueryReq) (*QueryRsp, error)
	mustEmbedUnimplementedTaoMarketDataSrvServer()
}

// UnimplementedTaoMarketDataSrvServer must be embedded to have forward compatible implementations.
type UnimplementedTaoMarketDataSrvServer struct {
}

func (UnimplementedTaoMarketDataSrvServer) BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedTaoMarketDataSrvServer) QueryOrder(context.Context, *QueryReq) (*QueryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (UnimplementedTaoMarketDataSrvServer) mustEmbedUnimplementedTaoMarketDataSrvServer() {}

// UnsafeTaoMarketDataSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaoMarketDataSrvServer will
// result in compilation errors.
type UnsafeTaoMarketDataSrvServer interface {
	mustEmbedUnimplementedTaoMarketDataSrvServer()
}

func RegisterTaoMarketDataSrvServer(s grpc.ServiceRegistrar, srv TaoMarketDataSrvServer) {
	s.RegisterService(&TaoMarketDataSrv_ServiceDesc, srv)
}

func _TaoMarketDataSrv_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoMarketDataSrvServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoMarketDataSrv_BatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoMarketDataSrvServer).BatchGet(ctx, req.(*BatchGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoMarketDataSrv_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoMarketDataSrvServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoMarketDataSrv_QueryOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoMarketDataSrvServer).QueryOrder(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaoMarketDataSrv_ServiceDesc is the grpc.ServiceDesc for TaoMarketDataSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaoMarketDataSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TaoMarketDataSrv",
	HandlerType: (*TaoMarketDataSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "batchGet",
			Handler:    _TaoMarketDataSrv_BatchGet_Handler,
		},
		{
			MethodName: "queryOrder",
			Handler:    _TaoMarketDataSrv_QueryOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taoExchange.proto",
}

const (
	TaoCoordinatorSrv_ListShards_FullMethodName        = "/grpc.TaoCoordinatorSrv/listShards"
	TaoCoordinatorSrv_LockShard_FullMethodName         = "/grpc.TaoCoordinatorSrv/lockShard"
	TaoCoordinatorSrv_UnlockShard_FullMethodName       = "/grpc.TaoCoordinatorSrv/unlockShard"
	TaoCoordinatorSrv_ListConnectorInfo_FullMethodName = "/grpc.TaoCoordinatorSrv/listConnectorInfo"
	TaoCoordinatorSrv_KeepLive_FullMethodName          = "/grpc.TaoCoordinatorSrv/keepLive"
)

// TaoCoordinatorSrvClient is the client API for TaoCoordinatorSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaoCoordinatorSrvClient interface {
	ListShards(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*LockShardRsp, error)
	LockShard(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*LockShardRsp, error)
	UnlockShard(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*CommonRsp, error)
	ListConnectorInfo(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*ListConnectorRsp, error)
	KeepLive(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*CommonRsp, error)
}

type taoCoordinatorSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewTaoCoordinatorSrvClient(cc grpc.ClientConnInterface) TaoCoordinatorSrvClient {
	return &taoCoordinatorSrvClient{cc}
}

func (c *taoCoordinatorSrvClient) ListShards(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*LockShardRsp, error) {
	out := new(LockShardRsp)
	err := c.cc.Invoke(ctx, TaoCoordinatorSrv_ListShards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoCoordinatorSrvClient) LockShard(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*LockShardRsp, error) {
	out := new(LockShardRsp)
	err := c.cc.Invoke(ctx, TaoCoordinatorSrv_LockShard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoCoordinatorSrvClient) UnlockShard(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, TaoCoordinatorSrv_UnlockShard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoCoordinatorSrvClient) ListConnectorInfo(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*ListConnectorRsp, error) {
	out := new(ListConnectorRsp)
	err := c.cc.Invoke(ctx, TaoCoordinatorSrv_ListConnectorInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taoCoordinatorSrvClient) KeepLive(ctx context.Context, in *ShardReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, TaoCoordinatorSrv_KeepLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaoCoordinatorSrvServer is the server API for TaoCoordinatorSrv service.
// All implementations must embed UnimplementedTaoCoordinatorSrvServer
// for forward compatibility
type TaoCoordinatorSrvServer interface {
	ListShards(context.Context, *CommonReq) (*LockShardRsp, error)
	LockShard(context.Context, *ShardReq) (*LockShardRsp, error)
	UnlockShard(context.Context, *ShardReq) (*CommonRsp, error)
	ListConnectorInfo(context.Context, *CommonReq) (*ListConnectorRsp, error)
	KeepLive(context.Context, *ShardReq) (*CommonRsp, error)
	mustEmbedUnimplementedTaoCoordinatorSrvServer()
}

// UnimplementedTaoCoordinatorSrvServer must be embedded to have forward compatible implementations.
type UnimplementedTaoCoordinatorSrvServer struct {
}

func (UnimplementedTaoCoordinatorSrvServer) ListShards(context.Context, *CommonReq) (*LockShardRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShards not implemented")
}
func (UnimplementedTaoCoordinatorSrvServer) LockShard(context.Context, *ShardReq) (*LockShardRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockShard not implemented")
}
func (UnimplementedTaoCoordinatorSrvServer) UnlockShard(context.Context, *ShardReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockShard not implemented")
}
func (UnimplementedTaoCoordinatorSrvServer) ListConnectorInfo(context.Context, *CommonReq) (*ListConnectorRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorInfo not implemented")
}
func (UnimplementedTaoCoordinatorSrvServer) KeepLive(context.Context, *ShardReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepLive not implemented")
}
func (UnimplementedTaoCoordinatorSrvServer) mustEmbedUnimplementedTaoCoordinatorSrvServer() {}

// UnsafeTaoCoordinatorSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaoCoordinatorSrvServer will
// result in compilation errors.
type UnsafeTaoCoordinatorSrvServer interface {
	mustEmbedUnimplementedTaoCoordinatorSrvServer()
}

func RegisterTaoCoordinatorSrvServer(s grpc.ServiceRegistrar, srv TaoCoordinatorSrvServer) {
	s.RegisterService(&TaoCoordinatorSrv_ServiceDesc, srv)
}

func _TaoCoordinatorSrv_ListShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoCoordinatorSrvServer).ListShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoCoordinatorSrv_ListShards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoCoordinatorSrvServer).ListShards(ctx, req.(*CommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoCoordinatorSrv_LockShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoCoordinatorSrvServer).LockShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoCoordinatorSrv_LockShard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoCoordinatorSrvServer).LockShard(ctx, req.(*ShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoCoordinatorSrv_UnlockShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoCoordinatorSrvServer).UnlockShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoCoordinatorSrv_UnlockShard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoCoordinatorSrvServer).UnlockShard(ctx, req.(*ShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoCoordinatorSrv_ListConnectorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoCoordinatorSrvServer).ListConnectorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoCoordinatorSrv_ListConnectorInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoCoordinatorSrvServer).ListConnectorInfo(ctx, req.(*CommonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaoCoordinatorSrv_KeepLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoCoordinatorSrvServer).KeepLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoCoordinatorSrv_KeepLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoCoordinatorSrvServer).KeepLive(ctx, req.(*ShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaoCoordinatorSrv_ServiceDesc is the grpc.ServiceDesc for TaoCoordinatorSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaoCoordinatorSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TaoCoordinatorSrv",
	HandlerType: (*TaoCoordinatorSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listShards",
			Handler:    _TaoCoordinatorSrv_ListShards_Handler,
		},
		{
			MethodName: "lockShard",
			Handler:    _TaoCoordinatorSrv_LockShard_Handler,
		},
		{
			MethodName: "unlockShard",
			Handler:    _TaoCoordinatorSrv_UnlockShard_Handler,
		},
		{
			MethodName: "listConnectorInfo",
			Handler:    _TaoCoordinatorSrv_ListConnectorInfo_Handler,
		},
		{
			MethodName: "keepLive",
			Handler:    _TaoCoordinatorSrv_KeepLive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taoExchange.proto",
}

const (
	TaoAdapterSrv_DoOrderCommond_FullMethodName = "/grpc.TaoAdapterSrv/doOrderCommond"
)

// TaoAdapterSrvClient is the client API for TaoAdapterSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaoAdapterSrvClient interface {
	DoOrderCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRsp, error)
}

type taoAdapterSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewTaoAdapterSrvClient(cc grpc.ClientConnInterface) TaoAdapterSrvClient {
	return &taoAdapterSrvClient{cc}
}

func (c *taoAdapterSrvClient) DoOrderCommond(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderRsp, error) {
	out := new(OrderRsp)
	err := c.cc.Invoke(ctx, TaoAdapterSrv_DoOrderCommond_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaoAdapterSrvServer is the server API for TaoAdapterSrv service.
// All implementations must embed UnimplementedTaoAdapterSrvServer
// for forward compatibility
type TaoAdapterSrvServer interface {
	DoOrderCommond(context.Context, *OrderReq) (*OrderRsp, error)
	mustEmbedUnimplementedTaoAdapterSrvServer()
}

// UnimplementedTaoAdapterSrvServer must be embedded to have forward compatible implementations.
type UnimplementedTaoAdapterSrvServer struct {
}

func (UnimplementedTaoAdapterSrvServer) DoOrderCommond(context.Context, *OrderReq) (*OrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoOrderCommond not implemented")
}
func (UnimplementedTaoAdapterSrvServer) mustEmbedUnimplementedTaoAdapterSrvServer() {}

// UnsafeTaoAdapterSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaoAdapterSrvServer will
// result in compilation errors.
type UnsafeTaoAdapterSrvServer interface {
	mustEmbedUnimplementedTaoAdapterSrvServer()
}

func RegisterTaoAdapterSrvServer(s grpc.ServiceRegistrar, srv TaoAdapterSrvServer) {
	s.RegisterService(&TaoAdapterSrv_ServiceDesc, srv)
}

func _TaoAdapterSrv_DoOrderCommond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaoAdapterSrvServer).DoOrderCommond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaoAdapterSrv_DoOrderCommond_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaoAdapterSrvServer).DoOrderCommond(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaoAdapterSrv_ServiceDesc is the grpc.ServiceDesc for TaoAdapterSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaoAdapterSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TaoAdapterSrv",
	HandlerType: (*TaoAdapterSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doOrderCommond",
			Handler:    _TaoAdapterSrv_DoOrderCommond_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taoExchange.proto",
}
