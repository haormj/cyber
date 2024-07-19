package transport

import (
	"errors"

	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/receiver"
	"github.com/haormj/cyber/transport/transmitter"
	"google.golang.org/protobuf/proto"
)

func CreateTransmitter[M proto.Message](attr *pb.RoleAttributes, mode pb.OptionalMode) (transmitter.Transmitter[M], error) {
	if attr == nil {
		return nil, errors.New("attr nil")
	}

	var t transmitter.Transmitter[M]
	var err error
	switch mode {
	case pb.OptionalMode_SHM:
		t, err = transmitter.NewShmTransmitter[M](attr)
	default:
		return nil, errors.New("not support mode" + mode.String())
	}

	if err != nil {
		return nil, err
	}

	if mode != pb.OptionalMode_HYBRID {
		t.Enable()
	}

	return t, nil
}

func CreateReceiver[M proto.Message](attr *pb.RoleAttributes, mode pb.OptionalMode,
	msgListener receiver.MessageListener[M]) (receiver.Receiver[M], error) {
	if attr == nil {
		return nil, errors.New("attr nil")
	}

	var r receiver.Receiver[M]
	switch mode {
	case pb.OptionalMode_SHM:
		r = receiver.NewShmReceiver(attr, msgListener)
	default:
		return nil, errors.New("not support mode" + mode.String())
	}

	if mode != pb.OptionalMode_INTRA {
		r.Enable()
	}

	return r, nil
}
