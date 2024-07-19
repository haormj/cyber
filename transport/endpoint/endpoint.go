package endpoint

import (
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/identity"
	"google.golang.org/protobuf/proto"
)

type Endpoint struct {
	Enabled bool
	id      *identity.Identity
	attr    *pb.RoleAttributes
}

func NewEndpoint(attr *pb.RoleAttributes) *Endpoint {
	e := &Endpoint{
		Enabled: false,
		id:      identity.NewIdentity(true),
		attr:    attr,
	}

	if attr.HostName == nil {
		attr.HostName = proto.String(common.GlobalDataInstance.HostName())
	}

	if attr.ProcessId == nil {
		attr.ProcessId = proto.Int32(int32(common.GlobalDataInstance.ProcessID()))
	}

	if attr.Id == nil {
		attr.Id = proto.Uint64(e.id.HashValue())
	}

	return e
}

func (e *Endpoint) ID() *identity.Identity {
	return e.id
}

func (e *Endpoint) Attributes() *pb.RoleAttributes {
	return e.attr
}
