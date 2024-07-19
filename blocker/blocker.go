package blocker

import (
	"container/list"
	"sync"
)

type Callback[T any] func(T)

type BlockerAttr struct {
	Capacity    uint32
	ChannelName string
}

type Blocker[T any] struct {
	blockerAttr        *BlockerAttr
	observedMsgQueue   *list.List
	publishedMsgQueue  *list.List
	msgMutex           sync.Mutex
	publishedCallbacks map[string]Callback[T]
	callbackMutex      sync.Mutex
	dummyMsg           T
}

func NewBlocker[T any](blockerAttr *BlockerAttr) *Blocker[T] {
	return &Blocker[T]{
		blockerAttr:        blockerAttr,
		observedMsgQueue:   list.New(),
		publishedMsgQueue:  list.New(),
		publishedCallbacks: make(map[string]Callback[T]),
	}
}

func (b *Blocker[T]) enqueue(msg T) {
	if b.blockerAttr.Capacity == 0 {
		return
	}

	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	b.publishedMsgQueue.PushFront(msg)
	for b.publishedMsgQueue.Len() > int(b.blockerAttr.Capacity) {
		e := b.publishedMsgQueue.Back()
		b.publishedMsgQueue.Remove(e)
	}
}

func (b *Blocker[T]) notify(msg T) {
	b.callbackMutex.Lock()
	defer b.callbackMutex.Unlock()

	for _, callback := range b.publishedCallbacks {
		callback(msg)
	}
}

func (b *Blocker[T]) reset(msg T) {
	b.msgMutex.Lock()
	b.observedMsgQueue = list.New()
	b.publishedMsgQueue = list.New()
	b.msgMutex.Unlock()

	b.callbackMutex.Lock()
	b.publishedCallbacks = make(map[string]Callback[T])
	b.callbackMutex.Unlock()

}

func (b *Blocker[T]) Publish(msg T) {
	b.enqueue(msg)
	b.notify(msg)
}

func (b *Blocker[T]) ClearObserved() {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	b.observedMsgQueue = list.New()
}

func (b *Blocker[T]) ClearPublished() {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	b.publishedMsgQueue = list.New()
}

func (b *Blocker[T]) Observe() {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	b.observedMsgQueue.PushBackList(b.publishedMsgQueue)
}

func (b *Blocker[T]) IsObservedEmpty() bool {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	return b.observedMsgQueue.Len() == 0
}

func (b *Blocker[T]) IsPublishedEmpty() bool {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	return b.publishedMsgQueue.Len() == 0
}

func (b *Blocker[T]) Subscribe(callbackID string, callback Callback[T]) bool {
	b.callbackMutex.Lock()
	defer b.callbackMutex.Unlock()

	if _, ok := b.publishedCallbacks[callbackID]; ok {
		return false
	}

	b.publishedCallbacks[callbackID] = callback
	return true
}

func (b *Blocker[T]) Unsubscribe(callbackID string) bool {
	b.callbackMutex.Lock()
	defer b.callbackMutex.Unlock()

	_, ok := b.publishedCallbacks[callbackID]
	delete(b.publishedCallbacks, callbackID)
	return ok
}

func (b *Blocker[T]) GetLatestObserved() T {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	if b.observedMsgQueue.Len() == 0 {
		return b.dummyMsg
	}

	return b.observedMsgQueue.Front().Value.(T)
}

func (b *Blocker[T]) GetOldestObserved() T {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	if b.observedMsgQueue.Len() == 0 {
		return b.dummyMsg
	}

	return b.observedMsgQueue.Back().Value.(T)
}

func (b *Blocker[T]) GetLatestPublished() T {
	b.msgMutex.Lock()
	defer b.msgMutex.Unlock()

	if b.publishedMsgQueue.Len() == 0 {
		return b.dummyMsg
	}

	return b.publishedMsgQueue.Front().Value.(T)
}

func (b *Blocker[T]) Capacity() uint32 {
	return b.blockerAttr.Capacity
}

func (b *Blocker[T]) SetCapacity(cap uint32) {
	b.blockerAttr.Capacity = cap
}

func (b *Blocker[T]) ChannelName() string {
	return b.blockerAttr.ChannelName
}
