package main

import (
	"github.com/haormj/cyber"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
)

type Component3 struct {
	component.BaseComponent3[*pb.SimpleMessage, *pb.SimpleMessage, *pb.SimpleMessage]
	configFilePath string
	node           node.Node
}

func (c *Component3) Init(configFilePath string, n node.Node) bool {
	c.configFilePath = configFilePath
	c.node = n
	return true
}

func (c *Component3) Proc(m0, m1, m2 *pb.SimpleMessage) bool {
	log.Logger.Debug("msg", "m0", m0.GetText(), "i", m0.GetInteger(), "m1", m1.GetText(), "i", m1.GetInteger(),
		"m2", m2.GetText(), "i", m2.GetInteger())
	return true
}

var _ = cyber.RegisterComponent3("Component3", &Component3{})
