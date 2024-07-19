// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: perf_conf.proto

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

type PerfType int32

const (
	PerfType_SCHED      PerfType = 1
	PerfType_TRANSPORT  PerfType = 2
	PerfType_DATA_CACHE PerfType = 3
	PerfType_ALL        PerfType = 4
)

// Enum value maps for PerfType.
var (
	PerfType_name = map[int32]string{
		1: "SCHED",
		2: "TRANSPORT",
		3: "DATA_CACHE",
		4: "ALL",
	}
	PerfType_value = map[string]int32{
		"SCHED":      1,
		"TRANSPORT":  2,
		"DATA_CACHE": 3,
		"ALL":        4,
	}
)

func (x PerfType) Enum() *PerfType {
	p := new(PerfType)
	*p = x
	return p
}

func (x PerfType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PerfType) Descriptor() protoreflect.EnumDescriptor {
	return file_perf_conf_proto_enumTypes[0].Descriptor()
}

func (PerfType) Type() protoreflect.EnumType {
	return &file_perf_conf_proto_enumTypes[0]
}

func (x PerfType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PerfType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PerfType(num)
	return nil
}

// Deprecated: Use PerfType.Descriptor instead.
func (PerfType) EnumDescriptor() ([]byte, []int) {
	return file_perf_conf_proto_rawDescGZIP(), []int{0}
}

type PerfConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable *bool     `protobuf:"varint,1,opt,name=enable,def=0" json:"enable,omitempty"`
	Type   *PerfType `protobuf:"varint,2,opt,name=type,enum=apollo.cyber.proto.PerfType,def=4" json:"type,omitempty"`
}

// Default values for PerfConf fields.
const (
	Default_PerfConf_Enable = bool(false)
	Default_PerfConf_Type   = PerfType_ALL
)

func (x *PerfConf) Reset() {
	*x = PerfConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_perf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerfConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerfConf) ProtoMessage() {}

func (x *PerfConf) ProtoReflect() protoreflect.Message {
	mi := &file_perf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerfConf.ProtoReflect.Descriptor instead.
func (*PerfConf) Descriptor() ([]byte, []int) {
	return file_perf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *PerfConf) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return Default_PerfConf_Enable
}

func (x *PerfConf) GetType() PerfType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_PerfConf_Type
}

var File_perf_conf_proto protoreflect.FileDescriptor

var file_perf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x65, 0x72, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x66, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x1d, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x03, 0x41, 0x4c,
	0x4c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x3d, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x66, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x43, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x10, 0x03, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x4c, 0x4c, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_perf_conf_proto_rawDescOnce sync.Once
	file_perf_conf_proto_rawDescData = file_perf_conf_proto_rawDesc
)

func file_perf_conf_proto_rawDescGZIP() []byte {
	file_perf_conf_proto_rawDescOnce.Do(func() {
		file_perf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_perf_conf_proto_rawDescData)
	})
	return file_perf_conf_proto_rawDescData
}

var file_perf_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_perf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_perf_conf_proto_goTypes = []interface{}{
	(PerfType)(0),    // 0: apollo.cyber.proto.PerfType
	(*PerfConf)(nil), // 1: apollo.cyber.proto.PerfConf
}
var file_perf_conf_proto_depIdxs = []int32{
	0, // 0: apollo.cyber.proto.PerfConf.type:type_name -> apollo.cyber.proto.PerfType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_perf_conf_proto_init() }
func file_perf_conf_proto_init() {
	if File_perf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_perf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerfConf); i {
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
			RawDescriptor: file_perf_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_perf_conf_proto_goTypes,
		DependencyIndexes: file_perf_conf_proto_depIdxs,
		EnumInfos:         file_perf_conf_proto_enumTypes,
		MessageInfos:      file_perf_conf_proto_msgTypes,
	}.Build()
	File_perf_conf_proto = out.File
	file_perf_conf_proto_rawDesc = nil
	file_perf_conf_proto_goTypes = nil
	file_perf_conf_proto_depIdxs = nil
}
