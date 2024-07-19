package data

import (
	"github.com/haormj/cyber/log"
)

type ChannelBuffer[T any] struct {
	*CacheBuffer[T]
	channelID uint64
}

func NewChannelBuffer[T any](channelID uint64, cacheBuffer *CacheBuffer[T]) *ChannelBuffer[T] {
	return &ChannelBuffer[T]{
		channelID:   channelID,
		CacheBuffer: cacheBuffer,
	}
}

func (b *ChannelBuffer[T]) ChannelID() uint64 {
	return b.channelID
}

func (b *ChannelBuffer[T]) Fetch(index *uint64, m *T) bool {
	b.Mutex().Lock()
	defer b.Mutex().Unlock()

	if b.Empty() {
		return false
	}

	switch {
	case *index == 0:
		*index = b.Tail()
	case *index == b.Tail()+1:
		return false
	case *index < b.Head():
		log.Logger.Warn("read buffer overflow, drop_message", "channelID", b.channelID,
			"index", *index, "tail", b.Tail(), "head", b.Head())
		*index = b.Tail()
	}
	*m = b.At(*index)
	return true
}

func (b *ChannelBuffer[T]) Latest(m *T) bool {
	b.Mutex().Lock()
	defer b.Mutex().Unlock()

	if b.Empty() {
		return false
	}

	*m = b.Back()
	return true
}

func (b *ChannelBuffer[T]) FetchMulti(fetchSize uint64, vec *[]T) bool {
	b.Mutex().Lock()
	defer b.Mutex().Unlock()

	if b.Empty() {
		return false
	}

	num := b.Size()
	if fetchSize < num {
		num = fetchSize
	}

	*vec = make([]T, num)
	for index := b.Tail() - num + 1; index <= b.Tail(); index++ {
		*vec = append(*vec, b.At(index))
	}
	return true
}
