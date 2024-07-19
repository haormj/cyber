package transmitter

import (
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/identity"
	"github.com/haormj/cyber/transport/message"
	"google.golang.org/protobuf/proto"
)

type Transmitter[M proto.Message] interface {
	ID() *identity.Identity
	Attributes() *pb.RoleAttributes
	Enable()
	Disable()
	Transmit(msg M) error
	TransmitWithMessageInfo(msg M, msgInfo *message.MessageInfo) error
}
