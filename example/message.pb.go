// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: example/message.proto

package example

import (
	_ "github.com/dungps/protoc-gen-go/options"
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

type SampleForExtend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SampleForExtend) Reset() {
	*x = SampleForExtend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleForExtend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleForExtend) ProtoMessage() {}

func (x *SampleForExtend) ProtoReflect() protoreflect.Message {
	mi := &file_example_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleForExtend.ProtoReflect.Descriptor instead.
func (*SampleForExtend) Descriptor() ([]byte, []int) {
	return file_example_message_proto_rawDescGZIP(), []int{0}
}

func (x *SampleForExtend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MapRepeatListSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test             map[string][]*SampleForExtend `protobuf:"bytes,1,rep,name=test,proto3" json:"test,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	*SampleForExtend `protobuf:"bytes,2,opt,name=sample_for_extend,json=sampleForExtend,proto3"`
}

func (x *MapRepeatListSample) Reset() {
	*x = MapRepeatListSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRepeatListSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRepeatListSample) ProtoMessage() {}

func (x *MapRepeatListSample) ProtoReflect() protoreflect.Message {
	mi := &file_example_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRepeatListSample.ProtoReflect.Descriptor instead.
func (*MapRepeatListSample) Descriptor() ([]byte, []int) {
	return file_example_message_proto_rawDescGZIP(), []int{1}
}

func (x *MapRepeatListSample) GetTest() map[string][]*SampleForExtend {
	if x != nil {
		return x.Test
	}
	return nil
}

func (x *MapRepeatListSample) GetSampleForExtend() *SampleForExtend {
	if x != nil {
		return x.SampleForExtend
	}
	return nil
}

var File_example_message_proto protoreflect.FileDescriptor

var file_example_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x25, 0x0a, 0x0f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x44,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0xa0, 0xcb, 0x02, 0x01, 0x52, 0x04,
	0x74, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x42, 0x04, 0xa8,
	0xcb, 0x02, 0x01, 0x52, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x1a, 0x55, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6e, 0x67, 0x70, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_message_proto_rawDescOnce sync.Once
	file_example_message_proto_rawDescData = file_example_message_proto_rawDesc
)

func file_example_message_proto_rawDescGZIP() []byte {
	file_example_message_proto_rawDescOnce.Do(func() {
		file_example_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_message_proto_rawDescData)
	})
	return file_example_message_proto_rawDescData
}

var file_example_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_message_proto_goTypes = []interface{}{
	(*SampleForExtend)(nil),     // 0: svc.example.SampleForExtend
	(*MapRepeatListSample)(nil), // 1: svc.example.MapRepeatListSample
	nil,                         // 2: svc.example.MapRepeatListSample.TestEntry
}
var file_example_message_proto_depIdxs = []int32{
	2, // 0: svc.example.MapRepeatListSample.test:type_name -> svc.example.MapRepeatListSample.TestEntry
	0, // 1: svc.example.MapRepeatListSample.sample_for_extend:type_name -> svc.example.SampleForExtend
	0, // 2: svc.example.MapRepeatListSample.TestEntry.value:type_name -> svc.example.SampleForExtend
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example_message_proto_init() }
func file_example_message_proto_init() {
	if File_example_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleForExtend); i {
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
		file_example_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRepeatListSample); i {
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
			RawDescriptor: file_example_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_message_proto_goTypes,
		DependencyIndexes: file_example_message_proto_depIdxs,
		MessageInfos:      file_example_message_proto_msgTypes,
	}.Build()
	File_example_message_proto = out.File
	file_example_message_proto_rawDesc = nil
	file_example_message_proto_goTypes = nil
	file_example_message_proto_depIdxs = nil
}
