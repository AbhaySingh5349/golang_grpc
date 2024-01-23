// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: protos/sqrt_message.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SqrtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *SqrtRequest) Reset() {
	*x = SqrtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sqrt_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtRequest) ProtoMessage() {}

func (x *SqrtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sqrt_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtRequest.ProtoReflect.Descriptor instead.
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return file_protos_sqrt_message_proto_rawDescGZIP(), []int{0}
}

func (x *SqrtRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SqrtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqrtResponse float32 `protobuf:"fixed32,1,opt,name=sqrt_response,json=sqrtResponse,proto3" json:"sqrt_response,omitempty"`
}

func (x *SqrtResponse) Reset() {
	*x = SqrtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sqrt_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtResponse) ProtoMessage() {}

func (x *SqrtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sqrt_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtResponse.ProtoReflect.Descriptor instead.
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return file_protos_sqrt_message_proto_rawDescGZIP(), []int{1}
}

func (x *SqrtResponse) GetSqrtResponse() float32 {
	if x != nil {
		return x.SqrtResponse
	}
	return 0
}

var File_protos_sqrt_message_proto protoreflect.FileDescriptor

var file_protos_sqrt_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x71, 0x72, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x22, 0x33, 0x0a, 0x0c, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x71, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x71, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_sqrt_message_proto_rawDescOnce sync.Once
	file_protos_sqrt_message_proto_rawDescData = file_protos_sqrt_message_proto_rawDesc
)

func file_protos_sqrt_message_proto_rawDescGZIP() []byte {
	file_protos_sqrt_message_proto_rawDescOnce.Do(func() {
		file_protos_sqrt_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_sqrt_message_proto_rawDescData)
	})
	return file_protos_sqrt_message_proto_rawDescData
}

var file_protos_sqrt_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_sqrt_message_proto_goTypes = []interface{}{
	(*SqrtRequest)(nil),  // 0: protos.SqrtRequest
	(*SqrtResponse)(nil), // 1: protos.SqrtResponse
}
var file_protos_sqrt_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_sqrt_message_proto_init() }
func file_protos_sqrt_message_proto_init() {
	if File_protos_sqrt_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_sqrt_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_sqrt_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_sqrt_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_sqrt_message_proto_goTypes,
		DependencyIndexes: file_protos_sqrt_message_proto_depIdxs,
		MessageInfos:      file_protos_sqrt_message_proto_msgTypes,
	}.Build()
	File_protos_sqrt_message_proto = out.File
	file_protos_sqrt_message_proto_rawDesc = nil
	file_protos_sqrt_message_proto_goTypes = nil
	file_protos_sqrt_message_proto_depIdxs = nil
}
