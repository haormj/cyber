package receiver

import (
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/dispatcher"
	"github.com/haormj/cyber/transport/endpoint"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/shm"
	"google.golang.org/protobuf/proto"
)

type ShmReceiver[M proto.Message] struct {
	*endpoint.Endpoint
	msgListener MessageListener[M]
}

func NewShmReceiver[M proto.Message](attr *pb.RoleAttributes, msgListener MessageListener[M]) *ShmReceiver[M] {
	r := &ShmReceiver[M]{
		Endpoint:    endpoint.NewEndpoint(attr),
		msgListener: msgListener,
	}

	return r
}

func (r *ShmReceiver[M]) onNewMessage(readableBlock *shm.ReadableBlock, info *message.MessageInfo) {
	if r.msgListener != nil {
		m := common.Zero[M]()
		if err := proto.Unmarshal(readableBlock.Buf[:readableBlock.Block.MsgSize()], m); err != nil {
			log.Logger.Error("parse message error", "err", err)
			return
		}
		r.msgListener(m, info, r.Endpoint.Attributes())
	}
}

func (r *ShmReceiver[M]) Enable() {
	if r.Endpoint.Enabled {
		return
	}
	dispatcher.ShmDispatcherInstance.AddListener(r.Endpoint.Attributes(), r.onNewMessage)
	r.Enabled = true
}

func (r *ShmReceiver[M]) Disable() {
	if !r.Enabled {
		return
	}
	dispatcher.ShmDispatcherInstance.RemoveListener(r.Attributes())
	r.Enabled = false
}

func (r *ShmReceiver[M]) OppositeEnable(oppoAttr *pb.RoleAttributes) {
	dispatcher.ShmDispatcherInstance.AddOppositeListener(r.Attributes(), oppoAttr, r.onNewMessage)
}

func (r *ShmReceiver[M]) OppositeDisable(oppoAttr *pb.RoleAttributes) {
	dispatcher.ShmDispatcherInstance.RemoveOppositeListener(r.Attributes(), oppoAttr)
}
