// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.1
// source: taoExchange.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_taoExchange_proto protoreflect.FileDescriptor

var file_taoExchange_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x6f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x10, 0x74, 0x61, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x74, 0x0a, 0x0e, 0x54,
	0x61, 0x6f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x72, 0x76, 0x12, 0x2e, 0x0a,
	0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x0e, 0x64, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x64, 0x12,
	0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x32, 0x77, 0x0a, 0x0f, 0x54, 0x61, 0x6f, 0x48, 0x69, 0x67, 0x68, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x72, 0x76, 0x12, 0x30, 0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x32, 0x76, 0x0a, 0x10, 0x54, 0x61,
	0x6f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x72, 0x76, 0x12, 0x32,
	0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x32, 0x9c, 0x02, 0x0a, 0x11, 0x54, 0x61, 0x6f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x72, 0x76, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0b, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12,
	0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x0e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x32, 0x43, 0x0a, 0x0d, 0x54, 0x61, 0x6f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53,
	0x72, 0x76, 0x12, 0x32, 0x0a, 0x0e, 0x64, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x20, 0x0a, 0x15, 0x74, 0x61, 0x6f, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x50,
	0x01, 0x5a, 0x05, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_taoExchange_proto_goTypes = []interface{}{
	(*QueryReq)(nil),         // 0: grpc.QueryReq
	(*OrderReq)(nil),         // 1: grpc.OrderReq
	(*BatchGetReq)(nil),      // 2: grpc.BatchGetReq
	(*CommonReq)(nil),        // 3: grpc.CommonReq
	(*ShardReq)(nil),         // 4: grpc.ShardReq
	(*QueryRsp)(nil),         // 5: grpc.QueryRsp
	(*OrderRsp)(nil),         // 6: grpc.OrderRsp
	(*CommonRsp)(nil),        // 7: grpc.CommonRsp
	(*BatchGetRsp)(nil),      // 8: grpc.BatchGetRsp
	(*LockShardRsp)(nil),     // 9: grpc.LockShardRsp
	(*ListConnectorRsp)(nil), // 10: grpc.ListConnectorRsp
}
var file_taoExchange_proto_depIdxs = []int32{
	0,  // 0: grpc.TaoExchangeSrv.queryOrder:input_type -> grpc.QueryReq
	1,  // 1: grpc.TaoExchangeSrv.doOrderCommond:input_type -> grpc.OrderReq
	1,  // 2: grpc.TaoHighStoreSrv.saveCommond:input_type -> grpc.OrderReq
	2,  // 3: grpc.TaoHighStoreSrv.batchGet:input_type -> grpc.BatchGetReq
	2,  // 4: grpc.TaoMarketDataSrv.batchGet:input_type -> grpc.BatchGetReq
	0,  // 5: grpc.TaoMarketDataSrv.queryOrder:input_type -> grpc.QueryReq
	3,  // 6: grpc.TaoCoordinatorSrv.listShards:input_type -> grpc.CommonReq
	4,  // 7: grpc.TaoCoordinatorSrv.lockShard:input_type -> grpc.ShardReq
	4,  // 8: grpc.TaoCoordinatorSrv.unlockShard:input_type -> grpc.ShardReq
	3,  // 9: grpc.TaoCoordinatorSrv.listConnectorInfo:input_type -> grpc.CommonReq
	4,  // 10: grpc.TaoCoordinatorSrv.keepLive:input_type -> grpc.ShardReq
	1,  // 11: grpc.TaoAdapterSrv.doOrderCommond:input_type -> grpc.OrderReq
	5,  // 12: grpc.TaoExchangeSrv.queryOrder:output_type -> grpc.QueryRsp
	6,  // 13: grpc.TaoExchangeSrv.doOrderCommond:output_type -> grpc.OrderRsp
	7,  // 14: grpc.TaoHighStoreSrv.saveCommond:output_type -> grpc.CommonRsp
	8,  // 15: grpc.TaoHighStoreSrv.batchGet:output_type -> grpc.BatchGetRsp
	8,  // 16: grpc.TaoMarketDataSrv.batchGet:output_type -> grpc.BatchGetRsp
	5,  // 17: grpc.TaoMarketDataSrv.queryOrder:output_type -> grpc.QueryRsp
	9,  // 18: grpc.TaoCoordinatorSrv.listShards:output_type -> grpc.LockShardRsp
	9,  // 19: grpc.TaoCoordinatorSrv.lockShard:output_type -> grpc.LockShardRsp
	7,  // 20: grpc.TaoCoordinatorSrv.unlockShard:output_type -> grpc.CommonRsp
	10, // 21: grpc.TaoCoordinatorSrv.listConnectorInfo:output_type -> grpc.ListConnectorRsp
	7,  // 22: grpc.TaoCoordinatorSrv.keepLive:output_type -> grpc.CommonRsp
	6,  // 23: grpc.TaoAdapterSrv.doOrderCommond:output_type -> grpc.OrderRsp
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_taoExchange_proto_init() }
func file_taoExchange_proto_init() {
	if File_taoExchange_proto != nil {
		return
	}
	file_taoContext_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_taoExchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_taoExchange_proto_goTypes,
		DependencyIndexes: file_taoExchange_proto_depIdxs,
	}.Build()
	File_taoExchange_proto = out.File
	file_taoExchange_proto_rawDesc = nil
	file_taoExchange_proto_goTypes = nil
	file_taoExchange_proto_depIdxs = nil
}
