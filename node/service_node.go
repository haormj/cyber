package node

import (
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type ServiceNode struct {
	nodeAttr *pb.RoleAttributes
}

func NewServiceNode(nodeName string) *ServiceNode {
	nodeAttr := &pb.RoleAttributes{}
	nodeAttr.HostName = proto.String(common.GlobalDataInstance.HostName())
	nodeAttr.ProcessId = proto.Int32(int32(common.GlobalDataInstance.ProcessID()))
	nodeAttr.NodeName = proto.String(nodeName)
	nodeID := common.GlobalDataInstance.RegisterNode(nodeName)
	nodeAttr.NodeId = proto.Uint64(nodeID)

	return &ServiceNode{
		nodeAttr: nodeAttr,
	}
}

func (n *ServiceNode) NodeName() string {
	return n.nodeAttr.GetNodeName()
}

func (n *ServiceNode) Attributes() *pb.RoleAttributes {
	return n.nodeAttr
}
