package main

import (
	"github.com/haormj/cyber"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
)

type Component struct {
	component.BaseComponent
	configFilePath string
	node           node.Node
}

func (c *Component) Init(configFilePath string, n node.Node) bool {
	c.configFilePath = configFilePath
	c.node = n
	log.Logger.Info("component init finished")
	return true
}

var _ = cyber.RegisterComponent("Component", &Component{})
