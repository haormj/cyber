package main

import (
	"github.com/haormj/cyber"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
)

type Component1 struct {
	component.BaseComponent1[*pb.SimpleMessage]
	configFilePath string
	node           node.Node
}

func (c *Component1) Init(configFilePath string, n node.Node) bool {
	c.configFilePath = configFilePath
	c.node = n
	return true
}

func (c *Component1) Proc(m0 *pb.SimpleMessage) bool {
	log.Logger.Debug("msg", "m0", m0.GetText(), "i", m0.GetInteger())
	return true
}

var _ = cyber.RegisterComponent1("Component1", &Component1{})
