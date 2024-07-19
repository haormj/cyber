package node

import (
	"sync"
	"time"

	"github.com/haormj/cyber/blocker"
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/data"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/receiver"
	"google.golang.org/protobuf/proto"
)

const DEFAULT_PENDING_QUEUE_SIZE uint32 = 1

var receiverMapMutex sync.Mutex
var receiverMap = make(map[string]any)

func GetReceiver[M proto.Message](roleAttr *pb.RoleAttributes) (receiver.Receiver[M], error) {
	channelName := roleAttr.GetChannelName()
	receiverMapMutex.Lock()
	defer receiverMapMutex.Unlock()

	v, ok := receiverMap[channelName]
	if ok {
		return v.(receiver.Receiver[M]), nil
	}

	r, err := transport.CreateReceiver[M](roleAttr, pb.OptionalMode_SHM,
		func(msg M, msgInfo *message.MessageInfo, attr *pb.RoleAttributes) {
			data.Dispatch(attr.GetChannelId(), msg)
		})
	if err != nil {
		return nil, err
	}
	receiverMap[channelName] = r
	return r, nil
}

type ReaderFunc[M proto.Message] func(msg M)

type Reader[M proto.Message] struct {
	roleAttr               *pb.RoleAttributes
	readerFunc             ReaderFunc[M]
	pendingQueueSize       uint32
	blocker                *blocker.Blocker[M]
	latestRecvTime         time.Time
	secondToLatestRecvTime time.Time
	receiver               receiver.Receiver[M]
	msgNotifyCh            chan struct{}
}

func NewReader[M proto.Message](roleAttr *pb.RoleAttributes, readerFunc ReaderFunc[M], pendingQueueSize uint32) (*Reader[M], error) {
	r := &Reader[M]{
		roleAttr:         roleAttr,
		readerFunc:       readerFunc,
		pendingQueueSize: pendingQueueSize,
		blocker: blocker.NewBlocker[M](&blocker.BlockerAttr{
			Capacity:    roleAttr.GetQosProfile().GetDepth(),
			ChannelName: roleAttr.GetChannelName(),
		}),
		msgNotifyCh: make(chan struct{}),
	}

	if err := r.init(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Reader[M]) init() error {
	var fn ReaderFunc[M]
	if r.readerFunc != nil {
		fn = func(msg M) {
			r.Enqueue(msg)
			r.readerFunc(msg)
		}
	} else {
		fn = func(msg M) {
			r.Enqueue(msg)
		}
	}

	dv := data.NewDataVisitor[M](data.VistorConfig{
		ChannelID: r.roleAttr.GetChannelId(),
		QueueSize: r.pendingQueueSize,
	})
	dv.RegisterNotifyCallback(func() {
		select {
		case r.msgNotifyCh <- struct{}{}:
		default:
		}
	})

	receiver, err := GetReceiver[M](r.roleAttr)
	if err != nil {
		return err
	}
	r.receiver = receiver
	r.roleAttr.Id = proto.Uint64(r.receiver.ID().HashValue())

	go func() {
		for range r.msgNotifyCh {
			m := common.Zero[M]()
			if dv.TryFetch(&m) {
				fn(m)
			}
		}
	}()

	return nil
}

func (r *Reader[M]) Enqueue(msg M) {
	r.secondToLatestRecvTime = r.latestRecvTime
	r.latestRecvTime = time.Now()
	r.blocker.Publish(msg)
}

func (r *Reader[M]) Observe() {
	r.blocker.Observe()
}

func (r *Reader[M]) Shutdown() {
}

func (r *Reader[M]) HasReceived() bool {
	return !r.blocker.IsPublishedEmpty()
}

func (r *Reader[M]) Empty() bool {
	return r.blocker.IsObservedEmpty()
}

// func (r *Reader[M]) GetDelay() time.Duration {
// 	if r.latestRecvTime == nil {
// 		return
// 	}
// }

func (r *Reader[M]) PendingQueueSize() uint32 {
	return r.pendingQueueSize
}

func (r *Reader[M]) GetLatestObserved() M {
	return r.blocker.GetLatestObserved()
}

func (r *Reader[M]) GetOldestObserved() M {
	return r.blocker.GetOldestObserved()
}

func (r *Reader[M]) ClearData() {
	r.blocker.ClearPublished()
	r.blocker.ClearObserved()
}

func (r *Reader[M]) SetHistoryDepth(depth uint32) {
	r.blocker.SetCapacity(depth)
}

func (r *Reader[M]) GetHistoryDepth() uint32 {
	return r.blocker.Capacity()
}

func (r *Reader[M]) ChannelID() uint64 {
	return r.roleAttr.GetChannelId()
}
