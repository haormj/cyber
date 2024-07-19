package cyber

import (
	"reflect"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

func RegisterTimerComponent(name string, c component.Component) bool {
	component.ComponentMap.Store(name, component.StartTimerComponent(func(config *pb.TimerComponentConfig) bool {
		com := component.NewTimerComponent(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component))
		return com.Initialize(config)
	}))
	return true
}

func RegisterComponent(name string, c component.Component) bool {
	component.ComponentMap.Store(name, component.StartComponent(func(config *pb.ComponentConfig) bool {
		com := component.NewDataDrivenComponent(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component))
		return com.Initialize(config)
	}))
	return true
}

func RegisterComponent1[M proto.Message](name string, c component.Component1[M]) bool {
	component.ComponentMap.Store(name, component.StartComponent(func(config *pb.ComponentConfig) bool {
		com := component.NewDataDrivenComponent1(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component1[M]))
		return com.Initialize(config)
	}))
	return true
}

func RegisterComponent2[M0, M1 proto.Message](name string, c component.Component2[M0, M1]) bool {
	component.ComponentMap.Store(name, component.StartComponent(func(config *pb.ComponentConfig) bool {
		com := component.NewDataDrivenComponent2(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component2[M0, M1]))
		return com.Initialize(config)
	}))
	return true
}

func RegisterComponent3[M0, M1, M2 proto.Message](name string, c component.Component3[M0, M1, M2]) bool {
	component.ComponentMap.Store(name, component.StartComponent(func(config *pb.ComponentConfig) bool {
		com := component.NewDataDrivenComponent3(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component3[M0, M1, M2]))
		return com.Initialize(config)
	}))
	return true
}

func RegisterComponent4[M0, M1, M2, M3 proto.Message](name string, c component.Component4[M0, M1, M2, M3]) bool {
	component.ComponentMap.Store(name, component.StartComponent(func(config *pb.ComponentConfig) bool {
		com := component.NewDataDrivenComponent4(common.ZeroByType(reflect.TypeOf(c)).Interface().(component.Component4[M0, M1, M2, M3]))
		return com.Initialize(config)
	}))
	return true
}
