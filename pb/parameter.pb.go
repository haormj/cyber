// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: parameter.proto

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

type ParamType int32

const (
	ParamType_NOT_SET  ParamType = 0
	ParamType_BOOL     ParamType = 1
	ParamType_INT      ParamType = 2
	ParamType_DOUBLE   ParamType = 3
	ParamType_STRING   ParamType = 4
	ParamType_PROTOBUF ParamType = 5
)

// Enum value maps for ParamType.
var (
	ParamType_name = map[int32]string{
		0: "NOT_SET",
		1: "BOOL",
		2: "INT",
		3: "DOUBLE",
		4: "STRING",
		5: "PROTOBUF",
	}
	ParamType_value = map[string]int32{
		"NOT_SET":  0,
		"BOOL":     1,
		"INT":      2,
		"DOUBLE":   3,
		"STRING":   4,
		"PROTOBUF": 5,
	}
)

func (x ParamType) Enum() *ParamType {
	p := new(ParamType)
	*p = x
	return p
}

func (x ParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_parameter_proto_enumTypes[0].Descriptor()
}

func (ParamType) Type() protoreflect.EnumType {
	return &file_parameter_proto_enumTypes[0]
}

func (x ParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ParamType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ParamType(num)
	return nil
}

// Deprecated: Use ParamType.Descriptor instead.
func (ParamType) EnumDescriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{0}
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type     *ParamType `protobuf:"varint,2,opt,name=type,enum=apollo.cyber.proto.ParamType" json:"type,omitempty"`
	TypeName *string    `protobuf:"bytes,3,opt,name=type_name,json=typeName" json:"type_name,omitempty"`
	// Types that are assignable to OneofValue:
	//
	//	*Param_BoolValue
	//	*Param_IntValue
	//	*Param_DoubleValue
	//	*Param_StringValue
	OneofValue isParam_OneofValue `protobuf_oneof:"oneof_value"`
	ProtoDesc  []byte             `protobuf:"bytes,8,opt,name=proto_desc,json=protoDesc" json:"proto_desc,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parameter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_parameter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Param) GetType() ParamType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ParamType_NOT_SET
}

func (x *Param) GetTypeName() string {
	if x != nil && x.TypeName != nil {
		return *x.TypeName
	}
	return ""
}

func (m *Param) GetOneofValue() isParam_OneofValue {
	if m != nil {
		return m.OneofValue
	}
	return nil
}

func (x *Param) GetBoolValue() bool {
	if x, ok := x.GetOneofValue().(*Param_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Param) GetIntValue() int64 {
	if x, ok := x.GetOneofValue().(*Param_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Param) GetDoubleValue() float64 {
	if x, ok := x.GetOneofValue().(*Param_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Param) GetStringValue() string {
	if x, ok := x.GetOneofValue().(*Param_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Param) GetProtoDesc() []byte {
	if x != nil {
		return x.ProtoDesc
	}
	return nil
}

type isParam_OneofValue interface {
	isParam_OneofValue()
}

type Param_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,oneof"`
}

type Param_IntValue struct {
	IntValue int64 `protobuf:"varint,5,opt,name=int_value,json=intValue,oneof"`
}

type Param_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,oneof"`
}

type Param_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=string_value,json=stringValue,oneof"`
}

func (*Param_BoolValue) isParam_OneofValue() {}

func (*Param_IntValue) isParam_OneofValue() {}

func (*Param_DoubleValue) isParam_OneofValue() {}

func (*Param_StringValue) isParam_OneofValue() {}

type NodeName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (x *NodeName) Reset() {
	*x = NodeName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parameter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeName) ProtoMessage() {}

func (x *NodeName) ProtoReflect() protoreflect.Message {
	mi := &file_parameter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeName.ProtoReflect.Descriptor instead.
func (*NodeName) Descriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{1}
}

func (x *NodeName) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type ParamName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (x *ParamName) Reset() {
	*x = ParamName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parameter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamName) ProtoMessage() {}

func (x *ParamName) ProtoReflect() protoreflect.Message {
	mi := &file_parameter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamName.ProtoReflect.Descriptor instead.
func (*ParamName) Descriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{2}
}

func (x *ParamName) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type BoolResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (x *BoolResult) Reset() {
	*x = BoolResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parameter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResult) ProtoMessage() {}

func (x *BoolResult) ProtoReflect() protoreflect.Message {
	mi := &file_parameter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResult.ProtoReflect.Descriptor instead.
func (*BoolResult) Descriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{3}
}

func (x *BoolResult) GetValue() bool {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return false
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param []*Param `protobuf:"bytes,1,rep,name=param" json:"param,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parameter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_parameter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_parameter_proto_rawDescGZIP(), []int{4}
}

func (x *Params) GetParam() []*Param {
	if x != nil {
		return x.Param
	}
	return nil
}

var File_parameter_proto protoreflect.FileDescriptor

var file_parameter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x42, 0x0d, 0x0a, 0x0b,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a,
	0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x22, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2f,
	0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2a,
	0x51, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f,
	0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x42, 0x55, 0x46,
	0x10, 0x05, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_parameter_proto_rawDescOnce sync.Once
	file_parameter_proto_rawDescData = file_parameter_proto_rawDesc
)

func file_parameter_proto_rawDescGZIP() []byte {
	file_parameter_proto_rawDescOnce.Do(func() {
		file_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(file_parameter_proto_rawDescData)
	})
	return file_parameter_proto_rawDescData
}

var file_parameter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_parameter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_parameter_proto_goTypes = []interface{}{
	(ParamType)(0),     // 0: apollo.cyber.proto.ParamType
	(*Param)(nil),      // 1: apollo.cyber.proto.Param
	(*NodeName)(nil),   // 2: apollo.cyber.proto.NodeName
	(*ParamName)(nil),  // 3: apollo.cyber.proto.ParamName
	(*BoolResult)(nil), // 4: apollo.cyber.proto.BoolResult
	(*Params)(nil),     // 5: apollo.cyber.proto.Params
}
var file_parameter_proto_depIdxs = []int32{
	0, // 0: apollo.cyber.proto.Param.type:type_name -> apollo.cyber.proto.ParamType
	1, // 1: apollo.cyber.proto.Params.param:type_name -> apollo.cyber.proto.Param
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_parameter_proto_init() }
func file_parameter_proto_init() {
	if File_parameter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parameter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_parameter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeName); i {
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
		file_parameter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamName); i {
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
		file_parameter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResult); i {
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
		file_parameter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
	file_parameter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Param_BoolValue)(nil),
		(*Param_IntValue)(nil),
		(*Param_DoubleValue)(nil),
		(*Param_StringValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parameter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_parameter_proto_goTypes,
		DependencyIndexes: file_parameter_proto_depIdxs,
		EnumInfos:         file_parameter_proto_enumTypes,
		MessageInfos:      file_parameter_proto_msgTypes,
	}.Build()
	File_parameter_proto = out.File
	file_parameter_proto_rawDesc = nil
	file_parameter_proto_goTypes = nil
	file_parameter_proto_depIdxs = nil
}
