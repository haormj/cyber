package dispatcher

import (
	"sync"
	"sync/atomic"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/message"
)

type MessageListener[T any] func(t T, info *message.MessageInfo)

type Dispatcher[T any] interface {
	Shutdown()
	AddListener(selfAttr *pb.RoleAttributes, listener MessageListener[T])
	RemoveListener(selfAttr *pb.RoleAttributes)
	AddOppositeListener(selfAttr *pb.RoleAttributes, oppositeAttr *pb.RoleAttributes, listener MessageListener[T])
	RemoveOppositeListener(selfAttr *pb.RoleAttributes, oppositeAttr *pb.RoleAttributes)
	HasChannel(channelID uint64) bool
}

type BaseDispatcher[M any] struct {
	isShutdown   atomic.Bool
	msgListeners sync.Map
}

func NewBaseDispatcher[M any]() *BaseDispatcher[M] {
	return &BaseDispatcher[M]{}
}

func (d *BaseDispatcher[M]) AddListener(selfAttr *pb.RoleAttributes, listener MessageListener[M]) {
	if d.isShutdown.Load() {
		return
	}

	channelID := selfAttr.GetChannelId()
	var listenerHandler *message.ListenerHandler[M]
	v, ok := d.msgListeners.Load(channelID)
	if ok {
		listenerHandler, ok = v.(*message.ListenerHandler[M])
		if !ok {
			log.Logger.Error("please ensure that readers with the same channel in the same process have the same message type",
				"channelName", selfAttr.GetChannelName())
			return
		}
	} else {
		log.Logger.Debug("new reader for channel", "channelName", common.GlobalDataInstance.GetChannelByID(channelID))
		listenerHandler = message.NewListenerHandler[M]()
		d.msgListeners.Store(channelID, listenerHandler)
	}

	listenerHandler.Connect(selfAttr.GetId(), message.Listener[M](listener))
}

func (d *BaseDispatcher[M]) RemoveListener(selfAttr *pb.RoleAttributes) {
	if d.isShutdown.Load() {
		return
	}

	channelID := selfAttr.GetChannelId()
	v, ok := d.msgListeners.Load(channelID)
	if !ok {
		return
	}
	listenerHandler := v.(*message.ListenerHandler[M])
	listenerHandler.Disconnect(selfAttr.GetId())
}

func (d *BaseDispatcher[M]) AddOppositeListener(selfAttr *pb.RoleAttributes, oppositeAttr *pb.RoleAttributes, listener MessageListener[M]) {
	if d.isShutdown.Load() {
		return
	}

	channelID := selfAttr.GetChannelId()
	var listenerHandler *message.ListenerHandler[M]
	v, ok := d.msgListeners.Load(channelID)
	if ok {
		listenerHandler, ok = v.(*message.ListenerHandler[M])
		if !ok {
			log.Logger.Error("please ensure that readers with the same channel in the same process have the same message type",
				"channelName", selfAttr.GetChannelName())
			return
		}
	} else {
		log.Logger.Debug("new reader for channel", "channelName", common.GlobalDataInstance.GetChannelByID(channelID))
		listenerHandler = message.NewListenerHandler[M]()
		d.msgListeners.Store(channelID, listenerHandler)
	}

	listenerHandler.OppositeConnect(selfAttr.GetId(), oppositeAttr.GetId(), message.Listener[M](listener))
}

func (d *BaseDispatcher[M]) RemoveOppositeListener(selfAttr *pb.RoleAttributes, oppositeAttr *pb.RoleAttributes) {
	if d.isShutdown.Load() {
		return
	}

	channelID := selfAttr.GetChannelId()
	v, ok := d.msgListeners.Load(channelID)
	if !ok {
		return
	}

	listenerHandler, ok := v.(*message.ListenerHandler[M])
	if !ok {
		return
	}

	listenerHandler.OppositeDisconnect(selfAttr.GetId(), oppositeAttr.GetId())
}

func (d *BaseDispatcher[M]) HasChannel(channelID uint64) bool {
	_, ok := d.msgListeners.Load(channelID)
	return ok
}

func (d *BaseDispatcher[M]) Shutdown() {
	d.isShutdown.Store(true)
	log.Logger.Debug("BaseDispatcher shutdown")
}
