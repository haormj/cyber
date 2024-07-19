package node

import (
	"errors"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type Node interface {
	Attributes() *pb.RoleAttributes
	Name() string
}

type ChannelCreator[M proto.Message] struct {
	node Node
}

func NewChannelCreator[M proto.Message](n Node) *ChannelCreator[M] {
	return &ChannelCreator[M]{
		node: n,
	}
}

func (c *ChannelCreator[M]) CreateWriter(channelName string) (*Writer[M], error) {
	roleAttr := &pb.RoleAttributes{}
	roleAttr.ChannelName = proto.String(channelName)
	return c.CreateWriterByAttr(roleAttr)
}

func (c *ChannelCreator[M]) CreateWriterByAttr(roleAttr *pb.RoleAttributes) (*Writer[M], error) {
	if roleAttr.ChannelName == nil || len(roleAttr.GetChannelName()) == 0 {
		return nil, errors.New("can't create a writer with empty channel name")
	}
	c.fillInAttr(roleAttr)
	return NewWriter[M](roleAttr)
}

func (c *ChannelCreator[M]) CreateReader(channelName string, readerFunc ReaderFunc[M]) (*Reader[M], error) {
	roleAttr := &pb.RoleAttributes{}
	roleAttr.ChannelName = proto.String(channelName)
	return c.CreateReaderByAttr(roleAttr, readerFunc, DEFAULT_PENDING_QUEUE_SIZE)
}

func (c *ChannelCreator[M]) CreateReaderByConfig(config *ReaderConfig, readerFunc ReaderFunc[M]) (*Reader[M], error) {
	roleAttr := &pb.RoleAttributes{}
	roleAttr.ChannelName = proto.String(config.ChannelName)
	roleAttr.QosProfile = config.QosProfile
	return c.CreateReaderByAttr(roleAttr, readerFunc, config.PendingQueueSize)
}

func (c *ChannelCreator[M]) CreateReaderByAttr(roleAttr *pb.RoleAttributes, readerFunc ReaderFunc[M],
	pendingQueueSize uint32) (*Reader[M], error) {
	if roleAttr.ChannelName == nil || len(roleAttr.GetChannelName()) == 0 {
		return nil, errors.New("can't create a reader with empty channel name")
	}

	c.fillInAttr(roleAttr)
	return NewReader(roleAttr, readerFunc, pendingQueueSize)
}

// TODO
func (c *ChannelCreator[M]) fillInAttr(attr *pb.RoleAttributes) {
	attr.HostName = proto.String(c.node.Attributes().GetHostName())
	attr.HostIp = proto.String(c.node.Attributes().GetHostIp())
	attr.ProcessId = proto.Int32(c.node.Attributes().GetProcessId())
	attr.NodeName = proto.String(c.node.Attributes().GetNodeName())
	attr.NodeId = proto.Uint64(c.node.Attributes().GetNodeId())
	channelID := common.GlobalDataInstance.RegisterChannel(attr.GetChannelName())
	attr.ChannelId = proto.Uint64(channelID)
	if attr.MessageType == nil {
		m := common.Zero[M]()
		attr.MessageType = proto.String(string(m.ProtoReflect().Descriptor().FullName()))
	}
	// TODO descriptor
	// if attr.QosProfile == nil {
	// 	attr.QosProfile =
	// }
}

type ServiceCreator[Request, Response proto.Message] struct {
	node Node
}
