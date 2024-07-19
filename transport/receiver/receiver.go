package receiver

import (
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/identity"
	"github.com/haormj/cyber/transport/message"
	"google.golang.org/protobuf/proto"
)

type MessageListener[M proto.Message] func(msg M, msgInfo *message.MessageInfo, attr *pb.RoleAttributes)

type Receiver[M proto.Message] interface {
	ID() *identity.Identity
	Attributes() *pb.RoleAttributes
	Enable()
	Disable()
}
