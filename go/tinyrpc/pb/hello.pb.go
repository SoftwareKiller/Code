// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/hello.proto

package pb

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

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *String) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_hello_proto protoreflect.FileDescriptor

var file_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x29, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x07, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData = file_proto_hello_proto_rawDesc
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_proto_rawDescData)
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_hello_proto_goTypes = []interface{}{
	(*String)(nil), // 0: String
}
var file_proto_hello_proto_depIdxs = []int32{
	0, // 0: HelloService.Hello:input_type -> String
	0, // 1: HelloService.Hello:output_type -> String
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
			RawDescriptor: file_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_rawDesc = nil
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}
