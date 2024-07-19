package node

import (
	"fmt"

	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport"
	"github.com/haormj/cyber/transport/transmitter"
	"google.golang.org/protobuf/proto"
)

type Writer[M proto.Message] struct {
	roleAttr    *pb.RoleAttributes
	transmitter transmitter.Transmitter[M]
}

func NewWriter[M proto.Message](roleAttr *pb.RoleAttributes) (*Writer[M], error) {
	t, err := transport.CreateTransmitter[M](roleAttr, pb.OptionalMode_SHM)
	if err != nil {
		return nil, fmt.Errorf("transport.CreateTransmitter: %w", err)
	}

	roleAttr.Id = proto.Uint64(t.ID().HashValue())
	w := &Writer[M]{
		roleAttr:    roleAttr,
		transmitter: t,
	}

	return w, nil
}

func (w *Writer[M]) Write(msg M) error {
	if err := w.transmitter.Transmit(msg); err != nil {
		return fmt.Errorf("transmitter.Transmit: %w", err)
	}

	return nil
}
