// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: protos/greetings_service.proto

package protos

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

var File_protos_greetings_service_proto protoreflect.FileDescriptor

var file_protos_greetings_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc7, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x61, 0x70, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_greetings_service_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: protos.GreetRequest
	(*GreetResponse)(nil), // 1: protos.GreetResponse
}
var file_protos_greetings_service_proto_depIdxs = []int32{
	0, // 0: protos.GreetService.Greet:input_type -> protos.GreetRequest
	0, // 1: protos.GreetService.GreetWithDeadline:input_type -> protos.GreetRequest
	0, // 2: protos.GreetService.GreetManyTimes:input_type -> protos.GreetRequest
	1, // 3: protos.GreetService.Greet:output_type -> protos.GreetResponse
	1, // 4: protos.GreetService.GreetWithDeadline:output_type -> protos.GreetResponse
	1, // 5: protos.GreetService.GreetManyTimes:output_type -> protos.GreetResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_greetings_service_proto_init() }
func file_protos_greetings_service_proto_init() {
	if File_protos_greetings_service_proto != nil {
		return
	}
	file_protos_greetings_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_greetings_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_greetings_service_proto_goTypes,
		DependencyIndexes: file_protos_greetings_service_proto_depIdxs,
	}.Build()
	File_protos_greetings_service_proto = out.File
	file_protos_greetings_service_proto_rawDesc = nil
	file_protos_greetings_service_proto_goTypes = nil
	file_protos_greetings_service_proto_depIdxs = nil
}
