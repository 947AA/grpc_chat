// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: computer_message.proto

package computer_message

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

type Cpu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand   string  `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Croes   uint32  `protobuf:"varint,3,opt,name=Croes,proto3" json:"Croes,omitempty"`
	Threads uint32  `protobuf:"varint,4,opt,name=threads,proto3" json:"threads,omitempty"`
	MinGhz  float32 `protobuf:"fixed32,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz  float32 `protobuf:"fixed32,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
}

func (x *Cpu) Reset() {
	*x = Cpu{}
	mi := &file_computer_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cpu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cpu) ProtoMessage() {}

func (x *Cpu) ProtoReflect() protoreflect.Message {
	mi := &file_computer_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cpu.ProtoReflect.Descriptor instead.
func (*Cpu) Descriptor() ([]byte, []int) {
	return file_computer_message_proto_rawDescGZIP(), []int{0}
}

func (x *Cpu) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Cpu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cpu) GetCroes() uint32 {
	if x != nil {
		return x.Croes
	}
	return 0
}

func (x *Cpu) GetThreads() uint32 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *Cpu) GetMinGhz() float32 {
	if x != nil {
		return x.MinGhz
	}
	return 0
}

func (x *Cpu) GetMaxGhz() float32 {
	if x != nil {
		return x.MaxGhz
	}
	return 0
}

var File_computer_message_proto protoreflect.FileDescriptor

var file_computer_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x72, 0x6f, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x72, 0x6f,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x47, 0x68, 0x7a, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x68, 0x7a,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x68, 0x7a, 0x42, 0x2a,
	0x5a, 0x28, 0x43, 0x3a, 0x2f, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_computer_message_proto_rawDescOnce sync.Once
	file_computer_message_proto_rawDescData = file_computer_message_proto_rawDesc
)

func file_computer_message_proto_rawDescGZIP() []byte {
	file_computer_message_proto_rawDescOnce.Do(func() {
		file_computer_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_computer_message_proto_rawDescData)
	})
	return file_computer_message_proto_rawDescData
}

var file_computer_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_computer_message_proto_goTypes = []any{
	(*Cpu)(nil), // 0: computer_message.cpu
}
var file_computer_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_computer_message_proto_init() }
func file_computer_message_proto_init() {
	if File_computer_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_computer_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_computer_message_proto_goTypes,
		DependencyIndexes: file_computer_message_proto_depIdxs,
		MessageInfos:      file_computer_message_proto_msgTypes,
	}.Build()
	File_computer_message_proto = out.File
	file_computer_message_proto_rawDesc = nil
	file_computer_message_proto_goTypes = nil
	file_computer_message_proto_depIdxs = nil
}
