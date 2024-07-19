package data

import (
	"sync"

	"github.com/haormj/cyber/state"
)

var buffersMapMutex sync.Mutex
var buffersMap sync.Map

func AddBuffer[T any](channelBuffer *ChannelBuffer[T]) {
	buffersMapMutex.Lock()
	defer buffersMapMutex.Unlock()

	v, ok := buffersMap.Load(channelBuffer.ChannelID())
	if ok {
		bufs := v.([]*ChannelBuffer[T])
		bufs = append(bufs, channelBuffer)
		buffersMap.Store(channelBuffer.ChannelID(), bufs)
	} else {
		buffersMap.Store(channelBuffer.ChannelID(), []*ChannelBuffer[T]{channelBuffer})
	}
}

func Dispatch[T any](channelID uint64, msg T) bool {
	if state.IsShutdown() {
		return false
	}

	v, ok := buffersMap.Load(channelID)
	if !ok {
		return false
	}
	bufs := v.([]*ChannelBuffer[T])
	for _, buf := range bufs {
		buf.Mutex().Lock()
		buf.Fill(msg)
		buf.Mutex().Unlock()
	}

	return DataNotifierInstance.Notify(channelID)
}
