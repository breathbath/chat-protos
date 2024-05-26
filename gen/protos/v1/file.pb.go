// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: protos/v1/file.proto

package gen

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Enum to represent known image or sound formats
type Format int32

const (
	Format_FORMAT_UNKNOWN_UNSPECIFIED Format = 0
	Format_FORMAT_JPEG                Format = 1
	Format_FORMAT_PNG                 Format = 2
	Format_FORMAT_MP3                 Format = 3
	Format_FORMAT_WAV                 Format = 4
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "FORMAT_UNKNOWN_UNSPECIFIED",
		1: "FORMAT_JPEG",
		2: "FORMAT_PNG",
		3: "FORMAT_MP3",
		4: "FORMAT_WAV",
	}
	Format_value = map[string]int32{
		"FORMAT_UNKNOWN_UNSPECIFIED": 0,
		"FORMAT_JPEG":                1,
		"FORMAT_PNG":                 2,
		"FORMAT_MP3":                 3,
		"FORMAT_WAV":                 4,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_v1_file_proto_enumTypes[0].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_protos_v1_file_proto_enumTypes[0]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_protos_v1_file_proto_rawDescGZIP(), []int{0}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents []byte `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Format   Format `protobuf:"varint,5,opt,name=format,proto3,enum=protos.v1.Format" json:"format,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_protos_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetFormat() Format {
	if x != nil {
		return x.Format
	}
	return Format_FORMAT_UNKNOWN_UNSPECIFIED
}

var File_protos_v1_file_proto protoreflect.FileDescriptor

var file_protos_v1_file_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xba, 0x48, 0x04, 0x7a, 0x02, 0x10,
	0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2a, 0x69, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1e,
	0x0a, 0x1a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4a, 0x50, 0x45, 0x47, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x50, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4d, 0x50, 0x33, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x57, 0x41, 0x56, 0x10, 0x04, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72,
	0x65, 0x61, 0x74, 0x68, 0x62, 0x61, 0x74, 0x68, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_file_proto_rawDescOnce sync.Once
	file_protos_v1_file_proto_rawDescData = file_protos_v1_file_proto_rawDesc
)

func file_protos_v1_file_proto_rawDescGZIP() []byte {
	file_protos_v1_file_proto_rawDescOnce.Do(func() {
		file_protos_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_file_proto_rawDescData)
	})
	return file_protos_v1_file_proto_rawDescData
}

var file_protos_v1_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_v1_file_proto_goTypes = []interface{}{
	(Format)(0),  // 0: protos.v1.Format
	(*File)(nil), // 1: protos.v1.File
}
var file_protos_v1_file_proto_depIdxs = []int32{
	0, // 0: protos.v1.File.format:type_name -> protos.v1.Format
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_v1_file_proto_init() }
func file_protos_v1_file_proto_init() {
	if File_protos_v1_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_protos_v1_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_v1_file_proto_goTypes,
		DependencyIndexes: file_protos_v1_file_proto_depIdxs,
		EnumInfos:         file_protos_v1_file_proto_enumTypes,
		MessageInfos:      file_protos_v1_file_proto_msgTypes,
	}.Build()
	File_protos_v1_file_proto = out.File
	file_protos_v1_file_proto_rawDesc = nil
	file_protos_v1_file_proto_goTypes = nil
	file_protos_v1_file_proto_depIdxs = nil
}
