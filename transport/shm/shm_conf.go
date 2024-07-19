package shm

import (
	"fmt"

	"github.com/haormj/cyber/log"
)

type ShmConf struct {
	// Extra size, Byte
	extraSize uint64
	// State size, Byte
	stateSize uint64
	// Block size, Byte
	blockSize uint64
	// Message info size, Byte
	messageInfoSize uint64
	// For message 0-10K
	blockNum16K    uint32
	messageSize16K uint64
	// For message 10K-100K
	blockNum128K    uint32
	messageSize128K uint64
	// For message 100K-1M
	blockNum1M    uint32
	messageSize1M uint64
	// For message 1M-6M
	blockNum8M    uint32
	messageSize8M uint64
	// For message 6M-10M
	blockNum16M    uint32
	messageSize16M uint64
	// For message 10M+
	blockNumMore    uint32
	messageSizeMore uint64

	ceilingMsgSize uint64
	blockBufSize   uint64
	blockNum       uint32
	managedShmSize uint64
}

func NewShmConf() *ShmConf {
	c := &ShmConf{
		extraSize:       1024 * 4,
		stateSize:       1024,
		blockSize:       1024,
		messageInfoSize: 1024,
		blockNum16K:     512,
		messageSize16K:  1024 * 16,
		blockNum128K:    128,
		messageSize128K: 1024 * 128,
		blockNum1M:      64,
		messageSize1M:   1024 * 1024,
		blockNum8M:      32,
		messageSize8M:   1024 * 1024 * 8,
		blockNum16M:     16,
		messageSize16M:  1024 * 1024 * 16,
		blockNumMore:    8,
		messageSizeMore: 1024 * 1024 * 32,
	}
	c.Update(c.messageSize16K)
	return c
}

func NewShmConfByRealMsgSize(realMsgSize uint64) *ShmConf {
	c := NewShmConf()
	c.Update(realMsgSize)
	return c
}

func (c *ShmConf) getCeilingMessageSize(realMsgSize uint64) uint64 {
	ceilingMsgSize := c.messageSize16K
	switch {
	case realMsgSize <= c.messageSize16K:
		ceilingMsgSize = c.messageSize16K
	case realMsgSize <= c.messageSize128K:
		ceilingMsgSize = c.messageSize128K
	case realMsgSize <= c.messageSize1M:
		ceilingMsgSize = c.messageSize1M
	case realMsgSize <= c.messageSize8M:
		ceilingMsgSize = c.messageSize8M
	case realMsgSize <= c.messageSize16M:
		ceilingMsgSize = c.messageSize16M
	default:
		ceilingMsgSize = c.messageSizeMore
	}
	return ceilingMsgSize
}

func (c *ShmConf) getBlockBufSize(ceilingMsgSize uint64) uint64 {
	return ceilingMsgSize + c.messageInfoSize
}

func (c *ShmConf) getBlockNum(ceilingMsgSize uint64) uint32 {
	var num uint32
	switch ceilingMsgSize {
	case c.messageSize16K:
		num = c.blockNum16K
	case c.messageSize128K:
		num = c.blockNum128K
	case c.messageSize1M:
		num = c.blockNum1M
	case c.messageSize8M:
		num = c.blockNum8M
	case c.messageSize16M:
		num = c.blockNum16M
	case c.messageSizeMore:
		num = c.blockNumMore
	default:
		log.Logger.Error(fmt.Sprintf("unknown ceiling_msg_size[%d]", ceilingMsgSize))
	}

	return num
}

func (c *ShmConf) Update(realMsgSize uint64) {
	c.ceilingMsgSize = c.getCeilingMessageSize(realMsgSize)
	c.blockBufSize = c.getBlockBufSize(c.ceilingMsgSize)
	c.blockNum = c.getBlockNum(c.ceilingMsgSize)
	c.managedShmSize = c.extraSize + c.stateSize + (c.blockSize+c.blockBufSize)*uint64(c.blockNum)
}

func (c *ShmConf) CeilingMsgSize() uint64 {
	return c.ceilingMsgSize
}

func (c *ShmConf) BlockBuffSize() uint64 {
	return c.blockBufSize
}

func (c *ShmConf) BlockNum() uint32 {
	return c.blockNum
}

func (c *ShmConf) ManagedShmSize() uint64 {
	return c.managedShmSize
}
