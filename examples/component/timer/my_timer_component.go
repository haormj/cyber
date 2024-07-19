package main

import (
	"github.com/haormj/cyber"
	"github.com/haormj/cyber/common"
	pb2 "github.com/haormj/cyber/examples/component/timer/pb"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type MyTimerComponent struct {
	configFilePath string
	node           node.Node
	i              int
	w              *node.Writer[*pb.SimpleMessage]
	r              *node.Reader[*pb.SimpleMessage]
}

func (c *MyTimerComponent) Init(configFilePath string, n node.Node) bool {
	c.configFilePath = configFilePath
	c.node = n

	config := &pb2.Config{}
	if err := common.GetProtoFromFile(c.configFilePath, config); err != nil {
		log.Logger.Error("get config error", "err", err)
		return false
	}

	creator := node.NewChannelCreator[*pb.SimpleMessage](c.node)
	w, err := creator.CreateWriter(config.GetChannelName())
	if err != nil {
		log.Logger.Error("create writer error", "err", err)
		return false
	}
	c.w = w

	r, err := creator.CreateReader(config.GetChannelName(), func(msg *pb.SimpleMessage) {
		log.Logger.Debug(c.node.Name(), "integer", msg.GetInteger(), "text", msg.GetText())
	})
	if err != nil {
		log.Logger.Error("create reader error", "err", err)
		return false
	}
	c.r = r

	return true
}

func (c *MyTimerComponent) Proc() bool {
	if err := c.w.Write(&pb.SimpleMessage{
		Integer: proto.Int32(int32(c.i)),
		Text:    proto.String(c.node.Name()),
	}); err != nil {
		log.Logger.Error("write error", "err", err)
		return false
	}

	c.i++
	return true
}

var _ = cyber.RegisterTimerComponent("MyTimerComponent", &MyTimerComponent{})
