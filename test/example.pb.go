// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: test/example.proto

package example

import (
	_ "github.com/jmploop/protoc-gen-go-asynq/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateUserPayload) Reset() {
	*x = CreateUserPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPayload) ProtoMessage() {}

func (x *CreateUserPayload) ProtoReflect() protoreflect.Message {
	mi := &file_test_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPayload.ProtoReflect.Descriptor instead.
func (*CreateUserPayload) Descriptor() ([]byte, []int) {
	return file_test_example_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateUserPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateUserPayload) Reset() {
	*x = UpdateUserPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPayload) ProtoMessage() {}

func (x *UpdateUserPayload) ProtoReflect() protoreflect.Message {
	mi := &file_test_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPayload.ProtoReflect.Descriptor instead.
func (*UpdateUserPayload) Descriptor() ([]byte, []int) {
	return file_test_example_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_test_example_proto protoreflect.FileDescriptor

var file_test_example_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0b, 0x61,
	0x73, 0x79, 0x6e, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x27, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x54, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x92, 0xa5, 0x95, 0x02, 0x0d, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x92, 0xa5, 0x95, 0x02,
	0x0d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x16,
	0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x3b, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_example_proto_rawDescOnce sync.Once
	file_test_example_proto_rawDescData = file_test_example_proto_rawDesc
)

func file_test_example_proto_rawDescGZIP() []byte {
	file_test_example_proto_rawDescOnce.Do(func() {
		file_test_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_example_proto_rawDescData)
	})
	return file_test_example_proto_rawDescData
}

var file_test_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_example_proto_goTypes = []interface{}{
	(*CreateUserPayload)(nil), // 0: example.CreateUserPayload
	(*UpdateUserPayload)(nil), // 1: example.UpdateUserPayload
	(*emptypb.Empty)(nil),     // 2: google.protobuf.Empty
}
var file_test_example_proto_depIdxs = []int32{
	0, // 0: example.User.CreateUser:input_type -> example.CreateUserPayload
	1, // 1: example.User.UpdateUser:input_type -> example.UpdateUserPayload
	2, // 2: example.User.CreateUser:output_type -> google.protobuf.Empty
	2, // 3: example.User.UpdateUser:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_example_proto_init() }
func file_test_example_proto_init() {
	if File_test_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPayload); i {
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
		file_test_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserPayload); i {
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
			RawDescriptor: file_test_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_example_proto_goTypes,
		DependencyIndexes: file_test_example_proto_depIdxs,
		MessageInfos:      file_test_example_proto_msgTypes,
	}.Build()
	File_test_example_proto = out.File
	file_test_example_proto_rawDesc = nil
	file_test_example_proto_goTypes = nil
	file_test_example_proto_depIdxs = nil
}
