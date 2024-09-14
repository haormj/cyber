package record

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
)

func CastMessageToAny(m proto.Message) *anypb.Any {
	return CastProtoReflectMessageToAny(m.ProtoReflect())
}

func CastDynamicMessageToAny(m *dynamicpb.Message) *anypb.Any {
	return CastProtoReflectMessageToAny(m.ProtoReflect())
}

func CastProtoReflectMessageToAny(m protoreflect.Message) *anypb.Any {
	typeURL := m.Get(m.Descriptor().Fields().ByName("type_url")).String()
	value := m.Get(m.Descriptor().Fields().ByName("value")).Bytes()
	return &anypb.Any{
		TypeUrl: typeURL,
		Value:   value,
	}
}

func CastAnyToDynamicMessage(a *anypb.Any) *dynamicpb.Message {
	t := &anypb.Any{}
	m := dynamicpb.NewMessage(t.ProtoReflect().Descriptor())
	m.Set(m.Descriptor().Fields().ByName("type_url"), protoreflect.ValueOf(a.TypeUrl))
	m.Set(m.Descriptor().Fields().ByName("value"), protoreflect.ValueOf(a.Value))
	return m
}
