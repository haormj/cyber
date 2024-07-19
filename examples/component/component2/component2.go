package main

import (
	"github.com/haormj/cyber"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
)

type Component2 struct {
	component.BaseComponent2[*pb.SimpleMessage, *pb.SimpleMessage]
	configFilePath string
	node           node.Node
}

func (c *Component2) Init(configFilePath string, n node.Node) bool {
	c.configFilePath = configFilePath
	c.node = n
	return true
}

func (c *Component2) Proc(m0, m1 *pb.SimpleMessage) bool {
	log.Logger.Debug("msg", "m0", m0.GetText(), "i", m0.GetInteger(), "m1", m1.GetText(), "i", m1.GetInteger())
	return true
}

var _ = cyber.RegisterComponent2("Component2", &Component2{})
