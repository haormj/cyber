package component

import (
	"time"

	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
)

type TimerComponent struct {
	component Component
	interval  time.Duration
	ticker    *time.Timer
	node      node.Node
}

func NewTimerComponent(component Component) *TimerComponent {
	return &TimerComponent{
		component: component,
	}
}

func (c *TimerComponent) Initialize(config *pb.TimerComponentConfig) bool {
	if config.Name == nil || config.Interval == nil {
		log.Logger.Error("missing required field in config file")
		return false
	}
	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		return false
	}
	c.interval = time.Duration(config.GetInterval()) * time.Millisecond
	c.ticker = (*time.Timer)(time.NewTicker(c.interval))
	// TODO
	go func() {
		for range c.ticker.C {
			c.Process()
		}
	}()
	return true
}

func (c *TimerComponent) Process() bool {
	return c.component.Proc()
}

func (c *TimerComponent) Clear() {
	c.ticker.Stop()
}
