// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: classic_conf.proto

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

type ClassicTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Prio      *uint32 `protobuf:"varint,2,opt,name=prio,def=1" json:"prio,omitempty"`
	GroupName *string `protobuf:"bytes,3,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
}

// Default values for ClassicTask fields.
const (
	Default_ClassicTask_Prio = uint32(1)
)

func (x *ClassicTask) Reset() {
	*x = ClassicTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_classic_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassicTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassicTask) ProtoMessage() {}

func (x *ClassicTask) ProtoReflect() protoreflect.Message {
	mi := &file_classic_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassicTask.ProtoReflect.Descriptor instead.
func (*ClassicTask) Descriptor() ([]byte, []int) {
	return file_classic_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ClassicTask) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ClassicTask) GetPrio() uint32 {
	if x != nil && x.Prio != nil {
		return *x.Prio
	}
	return Default_ClassicTask_Prio
}

func (x *ClassicTask) GetGroupName() string {
	if x != nil && x.GroupName != nil {
		return *x.GroupName
	}
	return ""
}

type SchedGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            *string        `protobuf:"bytes,1,req,name=name,def=default_grp" json:"name,omitempty"`
	ProcessorNum    *uint32        `protobuf:"varint,2,opt,name=processor_num,json=processorNum" json:"processor_num,omitempty"`
	Affinity        *string        `protobuf:"bytes,3,opt,name=affinity" json:"affinity,omitempty"`
	Cpuset          *string        `protobuf:"bytes,4,opt,name=cpuset" json:"cpuset,omitempty"`
	ProcessorPolicy *string        `protobuf:"bytes,5,opt,name=processor_policy,json=processorPolicy" json:"processor_policy,omitempty"`
	ProcessorPrio   *int32         `protobuf:"varint,6,opt,name=processor_prio,json=processorPrio,def=0" json:"processor_prio,omitempty"`
	Tasks           []*ClassicTask `protobuf:"bytes,7,rep,name=tasks" json:"tasks,omitempty"`
}

// Default values for SchedGroup fields.
const (
	Default_SchedGroup_Name          = string("default_grp")
	Default_SchedGroup_ProcessorPrio = int32(0)
)

func (x *SchedGroup) Reset() {
	*x = SchedGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_classic_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedGroup) ProtoMessage() {}

func (x *SchedGroup) ProtoReflect() protoreflect.Message {
	mi := &file_classic_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedGroup.ProtoReflect.Descriptor instead.
func (*SchedGroup) Descriptor() ([]byte, []int) {
	return file_classic_conf_proto_rawDescGZIP(), []int{1}
}

func (x *SchedGroup) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return Default_SchedGroup_Name
}

func (x *SchedGroup) GetProcessorNum() uint32 {
	if x != nil && x.ProcessorNum != nil {
		return *x.ProcessorNum
	}
	return 0
}

func (x *SchedGroup) GetAffinity() string {
	if x != nil && x.Affinity != nil {
		return *x.Affinity
	}
	return ""
}

func (x *SchedGroup) GetCpuset() string {
	if x != nil && x.Cpuset != nil {
		return *x.Cpuset
	}
	return ""
}

func (x *SchedGroup) GetProcessorPolicy() string {
	if x != nil && x.ProcessorPolicy != nil {
		return *x.ProcessorPolicy
	}
	return ""
}

func (x *SchedGroup) GetProcessorPrio() int32 {
	if x != nil && x.ProcessorPrio != nil {
		return *x.ProcessorPrio
	}
	return Default_SchedGroup_ProcessorPrio
}

func (x *SchedGroup) GetTasks() []*ClassicTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ClassicConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*SchedGroup `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (x *ClassicConf) Reset() {
	*x = ClassicConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_classic_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassicConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassicConf) ProtoMessage() {}

func (x *ClassicConf) ProtoReflect() protoreflect.Message {
	mi := &file_classic_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassicConf.ProtoReflect.Descriptor instead.
func (*ClassicConf) Descriptor() ([]byte, []int) {
	return file_classic_conf_proto_rawDescGZIP(), []int{2}
}

func (x *ClassicConf) GetGroups() []*SchedGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_classic_conf_proto protoreflect.FileDescriptor

var file_classic_conf_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x04, 0x70,
	0x72, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x31, 0x52, 0x04, 0x70, 0x72,
	0x69, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x92, 0x02, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x3a, 0x0b,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x70, 0x75, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x70, 0x75, 0x73, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x12,
	0x35, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x63, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x36, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_classic_conf_proto_rawDescOnce sync.Once
	file_classic_conf_proto_rawDescData = file_classic_conf_proto_rawDesc
)

func file_classic_conf_proto_rawDescGZIP() []byte {
	file_classic_conf_proto_rawDescOnce.Do(func() {
		file_classic_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_classic_conf_proto_rawDescData)
	})
	return file_classic_conf_proto_rawDescData
}

var file_classic_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_classic_conf_proto_goTypes = []interface{}{
	(*ClassicTask)(nil), // 0: apollo.cyber.proto.ClassicTask
	(*SchedGroup)(nil),  // 1: apollo.cyber.proto.SchedGroup
	(*ClassicConf)(nil), // 2: apollo.cyber.proto.ClassicConf
}
var file_classic_conf_proto_depIdxs = []int32{
	0, // 0: apollo.cyber.proto.SchedGroup.tasks:type_name -> apollo.cyber.proto.ClassicTask
	1, // 1: apollo.cyber.proto.ClassicConf.groups:type_name -> apollo.cyber.proto.SchedGroup
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_classic_conf_proto_init() }
func file_classic_conf_proto_init() {
	if File_classic_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_classic_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassicTask); i {
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
		file_classic_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedGroup); i {
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
		file_classic_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassicConf); i {
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
			RawDescriptor: file_classic_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_classic_conf_proto_goTypes,
		DependencyIndexes: file_classic_conf_proto_depIdxs,
		MessageInfos:      file_classic_conf_proto_msgTypes,
	}.Build()
	File_classic_conf_proto = out.File
	file_classic_conf_proto_rawDesc = nil
	file_classic_conf_proto_goTypes = nil
	file_classic_conf_proto_depIdxs = nil
}
