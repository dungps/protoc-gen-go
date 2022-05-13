// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: options/annotations.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_options_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5300,
		Name:          "transformer.map_value_repeat",
		Tag:           "varint,5300,opt,name=map_value_repeat",
		Filename:      "options/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5301,
		Name:          "transformer.is_extend_struct",
		Tag:           "varint,5301,opt,name=is_extend_struct",
		Filename:      "options/annotations.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool map_value_repeat = 5300;
	E_MapValueRepeat = &file_options_annotations_proto_extTypes[0]
	// optional bool is_extend_struct = 5301;
	E_IsExtendStruct = &file_options_annotations_proto_extTypes[1]
)

var File_options_annotations_proto protoreflect.FileDescriptor

var file_options_annotations_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x48, 0x0a, 0x10, 0x6d, 0x61,
	0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0x29,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x3a, 0x48, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x29, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6e,
	0x67, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_options_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_options_annotations_proto_depIdxs = []int32{
	0, // 0: transformer.map_value_repeat:extendee -> google.protobuf.FieldOptions
	0, // 1: transformer.is_extend_struct:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_annotations_proto_init() }
func file_options_annotations_proto_init() {
	if File_options_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_annotations_proto_goTypes,
		DependencyIndexes: file_options_annotations_proto_depIdxs,
		ExtensionInfos:    file_options_annotations_proto_extTypes,
	}.Build()
	File_options_annotations_proto = out.File
	file_options_annotations_proto_rawDesc = nil
	file_options_annotations_proto_goTypes = nil
	file_options_annotations_proto_depIdxs = nil
}