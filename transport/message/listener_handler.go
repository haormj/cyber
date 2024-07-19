package message

import (
	"sync"

	"github.com/haormj/cyber/base"
	"github.com/haormj/cyber/log"
)

type Listener[T any] func(t T, info *MessageInfo)

type ListenerHandler[T any] struct {
	isRawMessage bool
	signal       *base.Signal
	signalConns  map[uint64]*base.Connection
	signals      map[uint64]*base.Signal
	signalsConns map[uint64]map[uint64]*base.Connection
	rwLock       sync.RWMutex
}

func NewListenerHandler[T any]() *ListenerHandler[T] {
	return &ListenerHandler[T]{
		signal:       base.NewSignal(),
		signalConns:  make(map[uint64]*base.Connection),
		signals:      make(map[uint64]*base.Signal),
		signalsConns: make(map[uint64]map[uint64]*base.Connection),
	}
}

func (h *ListenerHandler[T]) Connect(selfID uint64, listener Listener[T]) {
	connection := h.signal.Connect(func(a ...any) {
		msg := a[0].(T)
		messageInfo := a[1].(*MessageInfo)
		listener(msg, messageInfo)
	})

	if !connection.IsConnected() {
		return
	}

	h.rwLock.Lock()
	h.signalConns[selfID] = connection
	h.rwLock.Unlock()
}

func (h *ListenerHandler[T]) Disconnect(selfID uint64) {
	h.rwLock.Lock()
	defer h.rwLock.Unlock()

	connection, ok := h.signalConns[selfID]
	if !ok {
		return
	}

	connection.Disconnect()
	delete(h.signalConns, selfID)
}

func (h *ListenerHandler[T]) OppositeConnect(selfID, oppoID uint64, listener Listener[T]) {
	h.rwLock.Lock()
	defer h.rwLock.Unlock()

	if _, ok := h.signals[oppoID]; !ok {
		h.signals[oppoID] = base.NewSignal()
	}

	connection := h.signals[oppoID].Connect(func(a ...any) {
		msg := a[0].(T)
		messageInfo := a[1].(*MessageInfo)
		listener(msg, messageInfo)
	})

	if !connection.IsConnected() {
		log.Logger.Warn("connect failed", "selfID", selfID, "oppoID", oppoID)
		return
	}

	if _, ok := h.signalsConns[oppoID]; !ok {
		h.signalsConns[oppoID] = make(map[uint64]*base.Connection)
	}

	h.signalsConns[oppoID][selfID] = connection
}

func (h *ListenerHandler[T]) OppositeDisconnect(selfID, oppoID uint64) {
	h.rwLock.Lock()
	defer h.rwLock.Unlock()

	if _, ok := h.signals[oppoID]; !ok {
		return
	}

	if _, ok := h.signalsConns[oppoID][selfID]; !ok {
		return
	}

	h.signalsConns[oppoID][selfID].Disconnect()
	delete(h.signalsConns[oppoID], selfID)
}

func (h *ListenerHandler[T]) Run(msg T, messageInfo *MessageInfo) {
	h.signal.Call([]any{msg, messageInfo}...)
	oppoID := messageInfo.SenderID.HashValue()

	h.rwLock.RLock()
	defer h.rwLock.RUnlock()
	oppoSignal, ok := h.signals[oppoID]
	if !ok {
		return
	}

	oppoSignal.Call([]any{msg, messageInfo}...)
}
