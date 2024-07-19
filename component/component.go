package component

import (
	"sync"

	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type StartComponent func(config *pb.ComponentConfig) bool
type StartTimerComponent func(config *pb.TimerComponentConfig) bool

var ComponentMap sync.Map

type Component interface {
	Init(configFilePath string, n node.Node) bool
	Proc() bool
}

type Component1[M proto.Message] interface {
	Init(configFilePath string, n node.Node) bool
	Proc(msg M) bool
}

type Component2[M0, M1 proto.Message] interface {
	Init(configFilePath string, n node.Node) bool
	Proc(m0 M0, m1 M1) bool
}

type Component3[M0, M1, M2 proto.Message] interface {
	Init(configFilePath string, n node.Node) bool
	Proc(m0 M0, m1 M1, m2 M2) bool
}

type Component4[M0, M1, M2, M3 proto.Message] interface {
	Init(configFilePath string, n node.Node) bool
	Proc(m0 M0, m1 M1, m2 M2, m3 M3) bool
}
