package data

import "sync"

type FusionCallback[T any] func(T)

type CacheBuffer[T any] struct {
	head           uint64
	tail           uint64
	capacity       uint64
	buffer         []T
	mutex          *sync.Mutex
	fusionCallback FusionCallback[T]
}

func NewCacheBuffer[T any](size int) *CacheBuffer[T] {
	return &CacheBuffer[T]{
		capacity: uint64(size + 1),
		buffer:   make([]T, size+1),
		mutex:    &sync.Mutex{},
	}
}

func (b *CacheBuffer[T]) getIndex(pos uint64) uint64 {
	return pos % b.capacity
}

func (b *CacheBuffer[T]) Head() uint64 {
	return b.head + 1
}

func (b *CacheBuffer[T]) Tail() uint64 {
	return b.tail
}

func (b *CacheBuffer[T]) Size() uint64 {
	return b.tail - b.head
}

func (b *CacheBuffer[T]) Front() T {
	return b.buffer[b.getIndex(b.head+1)]
}

func (b *CacheBuffer[T]) Back() T {
	return b.buffer[b.getIndex(b.tail)]
}

func (b *CacheBuffer[T]) Empty() bool {
	return b.tail == 0
}

func (b *CacheBuffer[T]) Full() bool {
	return b.capacity-1 == b.tail-b.head
}

func (b *CacheBuffer[T]) Capacity() uint64 {
	return b.capacity
}

func (b *CacheBuffer[T]) Fill(value T) {
	if b.fusionCallback != nil {
		b.fusionCallback(value)
	} else {
		if b.Full() {
			b.buffer[b.getIndex(b.head)] = value
			b.head++
			b.tail++
		} else {
			b.buffer[b.getIndex(b.tail+1)] = value
			b.tail++
		}
	}
}

func (b *CacheBuffer[T]) At(index uint64) T {
	return b.buffer[b.getIndex(index)]
}

func (b *CacheBuffer[T]) Mutex() *sync.Mutex {
	return b.mutex
}

func (b *CacheBuffer[T]) SetFusionCallback(callback FusionCallback[T]) {
	b.fusionCallback = callback
}
