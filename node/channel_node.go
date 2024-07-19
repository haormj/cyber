package node

import (
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type ReaderConfig struct {
	ChannelName      string
	QosProfile       *pb.QosProfile
	PendingQueueSize uint32
}

func NewReaderConfig() *ReaderConfig {
	return &ReaderConfig{
		QosProfile: &pb.QosProfile{
			History:     pb.QosHistoryPolicy_HISTORY_KEEP_LAST.Enum(),
			Depth:       proto.Uint32(1),
			Mps:         proto.Uint32(0),
			Reliability: pb.QosReliabilityPolicy_RELIABILITY_RELIABLE.Enum(),
			Durability:  pb.QosDurabilityPolicy_DURABILITY_VOLATILE.Enum(),
		},
		PendingQueueSize: DEFAULT_PENDING_QUEUE_SIZE,
	}
}

type ChannelNode struct {
	nodeAttr *pb.RoleAttributes
}

func NewChannelNode(nodeName string) *ChannelNode {
	nodeAttr := &pb.RoleAttributes{}
	nodeAttr.HostName = proto.String(common.GlobalDataInstance.HostName())
	nodeAttr.HostIp = proto.String(common.GlobalDataInstance.HostIP())
	nodeAttr.ProcessId = proto.Int32(int32(common.GlobalDataInstance.ProcessID()))
	nodeAttr.NodeName = proto.String(nodeName)
	nodeID := common.GlobalDataInstance.RegisterNode(nodeName)
	nodeAttr.NodeId = proto.Uint64(nodeID)

	return &ChannelNode{
		nodeAttr: nodeAttr,
	}
}

func (n *ChannelNode) Name() string {
	return n.nodeAttr.GetNodeName()
}

func (n *ChannelNode) Attributes() *pb.RoleAttributes {
	return n.nodeAttr
}
