// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: transport_conf.proto

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

type OptionalMode int32

const (
	OptionalMode_HYBRID OptionalMode = 0
	OptionalMode_INTRA  OptionalMode = 1
	OptionalMode_SHM    OptionalMode = 2
	OptionalMode_RTPS   OptionalMode = 3
)

// Enum value maps for OptionalMode.
var (
	OptionalMode_name = map[int32]string{
		0: "HYBRID",
		1: "INTRA",
		2: "SHM",
		3: "RTPS",
	}
	OptionalMode_value = map[string]int32{
		"HYBRID": 0,
		"INTRA":  1,
		"SHM":    2,
		"RTPS":   3,
	}
)

func (x OptionalMode) Enum() *OptionalMode {
	p := new(OptionalMode)
	*p = x
	return p
}

func (x OptionalMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionalMode) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_conf_proto_enumTypes[0].Descriptor()
}

func (OptionalMode) Type() protoreflect.EnumType {
	return &file_transport_conf_proto_enumTypes[0]
}

func (x OptionalMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *OptionalMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = OptionalMode(num)
	return nil
}

// Deprecated: Use OptionalMode.Descriptor instead.
func (OptionalMode) EnumDescriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{0}
}

type ShmMulticastLocator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port *uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (x *ShmMulticastLocator) Reset() {
	*x = ShmMulticastLocator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShmMulticastLocator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShmMulticastLocator) ProtoMessage() {}

func (x *ShmMulticastLocator) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShmMulticastLocator.ProtoReflect.Descriptor instead.
func (*ShmMulticastLocator) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ShmMulticastLocator) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *ShmMulticastLocator) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type ShmConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifierType *string              `protobuf:"bytes,1,opt,name=notifier_type,json=notifierType" json:"notifier_type,omitempty"`
	ShmType      *string              `protobuf:"bytes,2,opt,name=shm_type,json=shmType" json:"shm_type,omitempty"`
	ShmLocator   *ShmMulticastLocator `protobuf:"bytes,3,opt,name=shm_locator,json=shmLocator" json:"shm_locator,omitempty"`
}

func (x *ShmConf) Reset() {
	*x = ShmConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShmConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShmConf) ProtoMessage() {}

func (x *ShmConf) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShmConf.ProtoReflect.Descriptor instead.
func (*ShmConf) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{1}
}

func (x *ShmConf) GetNotifierType() string {
	if x != nil && x.NotifierType != nil {
		return *x.NotifierType
	}
	return ""
}

func (x *ShmConf) GetShmType() string {
	if x != nil && x.ShmType != nil {
		return *x.ShmType
	}
	return ""
}

func (x *ShmConf) GetShmLocator() *ShmMulticastLocator {
	if x != nil {
		return x.ShmLocator
	}
	return nil
}

type RtpsParticipantAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseDuration      *int32  `protobuf:"varint,1,opt,name=lease_duration,json=leaseDuration,def=12" json:"lease_duration,omitempty"`
	AnnouncementPeriod *int32  `protobuf:"varint,2,opt,name=announcement_period,json=announcementPeriod,def=3" json:"announcement_period,omitempty"`
	DomainIdGain       *uint32 `protobuf:"varint,3,opt,name=domain_id_gain,json=domainIdGain,def=200" json:"domain_id_gain,omitempty"`
	PortBase           *uint32 `protobuf:"varint,4,opt,name=port_base,json=portBase,def=10000" json:"port_base,omitempty"`
}

// Default values for RtpsParticipantAttr fields.
const (
	Default_RtpsParticipantAttr_LeaseDuration      = int32(12)
	Default_RtpsParticipantAttr_AnnouncementPeriod = int32(3)
	Default_RtpsParticipantAttr_DomainIdGain       = uint32(200)
	Default_RtpsParticipantAttr_PortBase           = uint32(10000)
)

func (x *RtpsParticipantAttr) Reset() {
	*x = RtpsParticipantAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtpsParticipantAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtpsParticipantAttr) ProtoMessage() {}

func (x *RtpsParticipantAttr) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtpsParticipantAttr.ProtoReflect.Descriptor instead.
func (*RtpsParticipantAttr) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{2}
}

func (x *RtpsParticipantAttr) GetLeaseDuration() int32 {
	if x != nil && x.LeaseDuration != nil {
		return *x.LeaseDuration
	}
	return Default_RtpsParticipantAttr_LeaseDuration
}

func (x *RtpsParticipantAttr) GetAnnouncementPeriod() int32 {
	if x != nil && x.AnnouncementPeriod != nil {
		return *x.AnnouncementPeriod
	}
	return Default_RtpsParticipantAttr_AnnouncementPeriod
}

func (x *RtpsParticipantAttr) GetDomainIdGain() uint32 {
	if x != nil && x.DomainIdGain != nil {
		return *x.DomainIdGain
	}
	return Default_RtpsParticipantAttr_DomainIdGain
}

func (x *RtpsParticipantAttr) GetPortBase() uint32 {
	if x != nil && x.PortBase != nil {
		return *x.PortBase
	}
	return Default_RtpsParticipantAttr_PortBase
}

type CommunicationMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SameProc *OptionalMode `protobuf:"varint,1,opt,name=same_proc,json=sameProc,enum=apollo.cyber.proto.OptionalMode,def=1" json:"same_proc,omitempty"` // INTRA SHM RTPS
	DiffProc *OptionalMode `protobuf:"varint,2,opt,name=diff_proc,json=diffProc,enum=apollo.cyber.proto.OptionalMode,def=2" json:"diff_proc,omitempty"` // SHM RTPS
	DiffHost *OptionalMode `protobuf:"varint,3,opt,name=diff_host,json=diffHost,enum=apollo.cyber.proto.OptionalMode,def=3" json:"diff_host,omitempty"` // RTPS
}

// Default values for CommunicationMode fields.
const (
	Default_CommunicationMode_SameProc = OptionalMode_INTRA
	Default_CommunicationMode_DiffProc = OptionalMode_SHM
	Default_CommunicationMode_DiffHost = OptionalMode_RTPS
)

func (x *CommunicationMode) Reset() {
	*x = CommunicationMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicationMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicationMode) ProtoMessage() {}

func (x *CommunicationMode) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicationMode.ProtoReflect.Descriptor instead.
func (*CommunicationMode) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{3}
}

func (x *CommunicationMode) GetSameProc() OptionalMode {
	if x != nil && x.SameProc != nil {
		return *x.SameProc
	}
	return Default_CommunicationMode_SameProc
}

func (x *CommunicationMode) GetDiffProc() OptionalMode {
	if x != nil && x.DiffProc != nil {
		return *x.DiffProc
	}
	return Default_CommunicationMode_DiffProc
}

func (x *CommunicationMode) GetDiffHost() OptionalMode {
	if x != nil && x.DiffHost != nil {
		return *x.DiffHost
	}
	return Default_CommunicationMode_DiffHost
}

type ResourceLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxHistoryDepth *uint32 `protobuf:"varint,1,opt,name=max_history_depth,json=maxHistoryDepth,def=1000" json:"max_history_depth,omitempty"`
}

// Default values for ResourceLimit fields.
const (
	Default_ResourceLimit_MaxHistoryDepth = uint32(1000)
)

func (x *ResourceLimit) Reset() {
	*x = ResourceLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimit) ProtoMessage() {}

func (x *ResourceLimit) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimit.ProtoReflect.Descriptor instead.
func (*ResourceLimit) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceLimit) GetMaxHistoryDepth() uint32 {
	if x != nil && x.MaxHistoryDepth != nil {
		return *x.MaxHistoryDepth
	}
	return Default_ResourceLimit_MaxHistoryDepth
}

type TransportConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShmConf           *ShmConf             `protobuf:"bytes,1,opt,name=shm_conf,json=shmConf" json:"shm_conf,omitempty"`
	ParticipantAttr   *RtpsParticipantAttr `protobuf:"bytes,2,opt,name=participant_attr,json=participantAttr" json:"participant_attr,omitempty"`
	CommunicationMode *CommunicationMode   `protobuf:"bytes,3,opt,name=communication_mode,json=communicationMode" json:"communication_mode,omitempty"`
	ResourceLimit     *ResourceLimit       `protobuf:"bytes,4,opt,name=resource_limit,json=resourceLimit" json:"resource_limit,omitempty"`
}

func (x *TransportConf) Reset() {
	*x = TransportConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportConf) ProtoMessage() {}

func (x *TransportConf) ProtoReflect() protoreflect.Message {
	mi := &file_transport_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportConf.ProtoReflect.Descriptor instead.
func (*TransportConf) Descriptor() ([]byte, []int) {
	return file_transport_conf_proto_rawDescGZIP(), []int{5}
}

func (x *TransportConf) GetShmConf() *ShmConf {
	if x != nil {
		return x.ShmConf
	}
	return nil
}

func (x *TransportConf) GetParticipantAttr() *RtpsParticipantAttr {
	if x != nil {
		return x.ParticipantAttr
	}
	return nil
}

func (x *TransportConf) GetCommunicationMode() *CommunicationMode {
	if x != nil {
		return x.CommunicationMode
	}
	return nil
}

func (x *TransportConf) GetResourceLimit() *ResourceLimit {
	if x != nil {
		return x.ResourceLimit
	}
	return nil
}

var File_transport_conf_proto protoreflect.FileDescriptor

var file_transport_conf_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x13, 0x53, 0x68,
	0x6d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x53, 0x68, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x68, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e,
	0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x6d, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x0a, 0x73, 0x68, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xc3, 0x01, 0x0a, 0x13,
	0x52, 0x74, 0x70, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x12, 0x29, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x31, 0x32, 0x52,
	0x0d, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x0a, 0x13, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x33, 0x52, 0x12,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x29, 0x0a, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f,
	0x67, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x03, 0x32, 0x30, 0x30, 0x52,
	0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x47, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a,
	0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x3a, 0x05, 0x31, 0x30, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x73,
	0x65, 0x22, 0xe2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x73, 0x61, 0x6d, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x6f,
	0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x05, 0x49, 0x4e,
	0x54, 0x52, 0x41, 0x52, 0x08, 0x73, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x12, 0x42, 0x0a,
	0x09, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x6f,
	0x64, 0x65, 0x3a, 0x03, 0x53, 0x48, 0x4d, 0x52, 0x08, 0x64, 0x69, 0x66, 0x66, 0x50, 0x72, 0x6f,
	0x63, 0x12, 0x43, 0x0a, 0x09, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x04, 0x52, 0x54, 0x50, 0x53, 0x52, 0x08, 0x64, 0x69,
	0x66, 0x66, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x3a, 0x04, 0x31, 0x30, 0x30, 0x30, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x70, 0x74, 0x68, 0x22, 0xbb, 0x02, 0x0a, 0x0d, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x36, 0x0a, 0x08, 0x73,
	0x68, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x68, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x07, 0x73, 0x68, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x12, 0x52, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x74, 0x70, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x54, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63, 0x79, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a,
	0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2a, 0x38, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x59, 0x42, 0x52, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x52, 0x41, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x48, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x54, 0x50, 0x53, 0x10,
	0x03, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_transport_conf_proto_rawDescOnce sync.Once
	file_transport_conf_proto_rawDescData = file_transport_conf_proto_rawDesc
)

func file_transport_conf_proto_rawDescGZIP() []byte {
	file_transport_conf_proto_rawDescOnce.Do(func() {
		file_transport_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_conf_proto_rawDescData)
	})
	return file_transport_conf_proto_rawDescData
}

var file_transport_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transport_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transport_conf_proto_goTypes = []interface{}{
	(OptionalMode)(0),           // 0: apollo.cyber.proto.OptionalMode
	(*ShmMulticastLocator)(nil), // 1: apollo.cyber.proto.ShmMulticastLocator
	(*ShmConf)(nil),             // 2: apollo.cyber.proto.ShmConf
	(*RtpsParticipantAttr)(nil), // 3: apollo.cyber.proto.RtpsParticipantAttr
	(*CommunicationMode)(nil),   // 4: apollo.cyber.proto.CommunicationMode
	(*ResourceLimit)(nil),       // 5: apollo.cyber.proto.ResourceLimit
	(*TransportConf)(nil),       // 6: apollo.cyber.proto.TransportConf
}
var file_transport_conf_proto_depIdxs = []int32{
	1, // 0: apollo.cyber.proto.ShmConf.shm_locator:type_name -> apollo.cyber.proto.ShmMulticastLocator
	0, // 1: apollo.cyber.proto.CommunicationMode.same_proc:type_name -> apollo.cyber.proto.OptionalMode
	0, // 2: apollo.cyber.proto.CommunicationMode.diff_proc:type_name -> apollo.cyber.proto.OptionalMode
	0, // 3: apollo.cyber.proto.CommunicationMode.diff_host:type_name -> apollo.cyber.proto.OptionalMode
	2, // 4: apollo.cyber.proto.TransportConf.shm_conf:type_name -> apollo.cyber.proto.ShmConf
	3, // 5: apollo.cyber.proto.TransportConf.participant_attr:type_name -> apollo.cyber.proto.RtpsParticipantAttr
	4, // 6: apollo.cyber.proto.TransportConf.communication_mode:type_name -> apollo.cyber.proto.CommunicationMode
	5, // 7: apollo.cyber.proto.TransportConf.resource_limit:type_name -> apollo.cyber.proto.ResourceLimit
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_transport_conf_proto_init() }
func file_transport_conf_proto_init() {
	if File_transport_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShmMulticastLocator); i {
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
		file_transport_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShmConf); i {
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
		file_transport_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RtpsParticipantAttr); i {
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
		file_transport_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicationMode); i {
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
		file_transport_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimit); i {
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
		file_transport_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportConf); i {
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
			RawDescriptor: file_transport_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_conf_proto_goTypes,
		DependencyIndexes: file_transport_conf_proto_depIdxs,
		EnumInfos:         file_transport_conf_proto_enumTypes,
		MessageInfos:      file_transport_conf_proto_msgTypes,
	}.Build()
	File_transport_conf_proto = out.File
	file_transport_conf_proto_rawDesc = nil
	file_transport_conf_proto_goTypes = nil
	file_transport_conf_proto_depIdxs = nil
}