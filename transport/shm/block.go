package shm

import (
	"errors"
	"sync/atomic"
	"unsafe"

	"github.com/haormj/cyber/log"
)

const (
	BlockSize              = 32
	kRWLockFree      int32 = 0
	kWriteExclusive  int32 = -1
	kMaxTryLockTimes int32 = 5
)

type Block struct {
	lockNum     *atomic.Int32
	msgSize     *uint64
	msgInfoSize *uint64
}

func NewBlockFromShm(buf []byte) (*Block, error) {
	block := &Block{}
	block.lockNum = (*atomic.Int32)(unsafe.Pointer(&buf[8]))
	block.msgSize = (*uint64)(unsafe.Pointer(&buf[16]))
	block.msgInfoSize = (*uint64)(unsafe.Pointer(&buf[24]))
	if block.lockNum == nil || block.msgSize == nil || block.msgInfoSize == nil {
		return nil, errors.New("create block failed")
	}

	return block, nil
}

func (b *Block) MsgSize() uint64 {
	return *b.msgSize
}

func (b *Block) SetMsgSize(i uint64) {
	*b.msgSize = i
}

func (b *Block) MsgInfoSize() uint64 {
	return *b.msgInfoSize
}

func (b *Block) SetMsgInfoSize(i uint64) {
	*b.msgInfoSize = i
}

func (b *Block) TryLockForWrite() bool {
	if !b.lockNum.CompareAndSwap(kRWLockFree, kWriteExclusive) {
		log.Logger.Debug("lock num", "lockNum", b.lockNum.Load())
		return false
	}
	return true
}

func (b *Block) TryLockForRead() bool {
	var tryTimes int32
	var lockNum int32
	for {
		if tryTimes > kMaxTryLockTimes {
			log.Logger.Info("fail to add read lock num", "lockNum", lockNum)
			return false
		}

		lockNum = b.lockNum.Load()
		if lockNum < kRWLockFree {
			log.Logger.Info("block is being written")
			return false
		}

		if b.lockNum.CompareAndSwap(lockNum, lockNum+1) {
			return true
		}

		tryTimes++
	}
}

func (b *Block) ReleaseWriteLock() {
	b.lockNum.Add(1)
}

func (b *Block) ReleaseReadLock() {
	b.lockNum.Add(-1)
}
