package shm

import (
	"github.com/haormj/cyber/log"

	"github.com/haormj/cyber/common"
)

type WritableBlock struct {
	Index uint32
	Block *Block
	Buf   []byte
}

type ReadableBlock WritableBlock

type Segment interface {
	AcquireBlockToWrite(msgSize uint64, writableBlock *WritableBlock) bool
	ReleaseWrittenBlock(writableBlock *WritableBlock)
	AcquireBlockToRead(readableBlock *ReadableBlock) bool
	ReleaseReadBlock(readableBlock *ReadableBlock)
	Destroy() bool
	Reset()
	Remove() bool
	OpenOnly() bool
	OpenOrCreate() bool
	Type() string
}

func NewSegment(channelID uint64) Segment {
	segmentType := "xsi"
	config := common.GlobalDataInstance.Config()
	if config != nil && config.TransportConf != nil &&
		config.TransportConf.ShmConf != nil &&
		config.TransportConf.ShmConf.ShmType != nil {
		segmentType = config.TransportConf.ShmConf.GetShmType()
	}

	log.Logger.Debug("segment type", "segmentType", segmentType)

	switch segmentType {
	case "xsi":
		return NewXSISegment(channelID)
	case "posix":
		panic("not impl posix")
	default:
		panic("not support" + segmentType)
	}
}
