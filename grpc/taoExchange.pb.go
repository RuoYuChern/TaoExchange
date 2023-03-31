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
	0x00, 0x42, 0x20, 0x0a, 0x15, 0x74, 0x61, 0x6f, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x05, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_taoExchange_proto_goTypes = []interface{}{
	(*QueryReq)(nil), // 0: grpc.QueryReq
	(*OrderReq)(nil), // 1: grpc.OrderReq
	(*QueryRsp)(nil), // 2: grpc.QueryRsp
	(*OrderRsp)(nil), // 3: grpc.OrderRsp
}
var file_taoExchange_proto_depIdxs = []int32{
	0, // 0: grpc.TaoExchangeSrv.queryOrder:input_type -> grpc.QueryReq
	1, // 1: grpc.TaoExchangeSrv.doOrderCommond:input_type -> grpc.OrderReq
	2, // 2: grpc.TaoExchangeSrv.queryOrder:output_type -> grpc.QueryRsp
	3, // 3: grpc.TaoExchangeSrv.doOrderCommond:output_type -> grpc.OrderRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumServices:   1,
		},
		GoTypes:           file_taoExchange_proto_goTypes,
		DependencyIndexes: file_taoExchange_proto_depIdxs,
	}.Build()
	File_taoExchange_proto = out.File
	file_taoExchange_proto_rawDesc = nil
	file_taoExchange_proto_goTypes = nil
	file_taoExchange_proto_depIdxs = nil
}