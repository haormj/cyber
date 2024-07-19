package shm

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"

	"github.com/haormj/cyber/log"

	"golang.org/x/sys/unix"
)

var _ Segment = &XSISegment{}

type XSISegment struct {
	init          bool
	conf          *ShmConf
	channelID     uint64
	state         *State
	blocks        []*Block
	managedShm    []byte
	blockBufLock  sync.Mutex
	blockBufAddrs map[uint32][]byte
	key           int
	shmID         int
}

func NewXSISegment(channelID uint64) *XSISegment {
	return &XSISegment{
		init:          false,
		conf:          NewShmConf(),
		channelID:     channelID,
		blockBufAddrs: make(map[uint32][]byte),
		key:           int(channelID),
	}
}

func (s *XSISegment) remap() bool {
	s.init = false
	log.Logger.Debug("before reset")
	s.Reset()
	log.Logger.Debug("after reset")
	return s.OpenOnly()
}

func (s *XSISegment) recreate(msgSize uint64) bool {
	s.init = false
	s.state.SetNeedRemap(true)
	s.Reset()
	s.conf.Update(msgSize)
	return s.OpenOrCreate()
}

func (s *XSISegment) getNextWritableBlockIndex() uint32 {
	blockNum := s.conf.BlockNum()
	for {
		tryIdx := s.state.FetchAddSeq(1) % blockNum
		if s.blocks[tryIdx].TryLockForWrite() {
			return tryIdx
		}
	}
}

func (s *XSISegment) AcquireBlockToWrite(msgSize uint64, writableBlock *WritableBlock) bool {
	if writableBlock == nil {
		return false
	}

	if !s.init && !s.OpenOrCreate() {
		log.Logger.Error("create shm failed, can't write now")
		return false
	}

	var result bool = true
	if s.state.NeedRemap() {
		result = s.remap()
	}

	if msgSize > s.conf.CeilingMsgSize() {
		log.Logger.Info("larger than current shm_buffer_size, need recreate",
			"msgSize", msgSize, "ceilingMsgSize", s.conf.CeilingMsgSize())
		result = s.recreate(msgSize)
	}

	if !result {
		log.Logger.Error("segment update failed")
		return false
	}

	index := s.getNextWritableBlockIndex()
	writableBlock.Index = index
	writableBlock.Block = s.blocks[index]
	writableBlock.Buf = s.blockBufAddrs[index]
	return true
}

func (s *XSISegment) ReleaseWrittenBlock(writableBlock *WritableBlock) {
	if writableBlock.Index >= s.conf.BlockNum() {
		return
	}
	s.blocks[writableBlock.Index].ReleaseWriteLock()
}

func (s *XSISegment) AcquireBlockToRead(readableBlock *ReadableBlock) bool {
	if readableBlock == nil {
		return false
	}

	if !s.init && !s.OpenOnly() {
		log.Logger.Error("failed to open shared memory, can't read now")
		return false
	}

	if readableBlock.Index > s.conf.BlockNum() {
		log.Logger.Error("invalid block_index", "index", readableBlock.Index,
			"blockNum", s.conf.BlockNum())
		return false
	}

	var result bool = true
	if s.state.NeedRemap() {
		result = s.remap()
	}

	if !result {
		log.Logger.Error("segment update failed")
		return false
	}

	if !s.blocks[readableBlock.Index].TryLockForRead() {
		return false
	}

	readableBlock.Block = s.blocks[readableBlock.Index]
	readableBlock.Buf = s.blockBufAddrs[readableBlock.Index]
	return true
}

func (s *XSISegment) ReleaseReadBlock(readableBlock *ReadableBlock) {
	if readableBlock.Index > s.conf.BlockNum() {
		return
	}
	s.blocks[readableBlock.Index].ReleaseReadLock()
}

func (s *XSISegment) Destroy() bool {
	if !s.init {
		return true
	}
	s.init = false

	s.state.DecreaseReferenceCounts()
	if s.state.ReferenceCounts() == 0 {
		return s.Remove()
	}

	log.Logger.Debug("destroy")

	return true
}

func (s *XSISegment) Type() string {
	return "xsi"
}

func (s *XSISegment) OpenOrCreate() bool {
	if s.init {
		return true
	}

	// create managed shm
	var retry int
	var err error
	for retry < 2 {
		s.shmID, err = unix.SysvShmGet(s.key, int(s.conf.ManagedShmSize()), 0644|unix.IPC_CREAT|unix.IPC_EXCL)
		if err == nil {
			break
		}

		if errors.Is(err, syscall.EINVAL) {
			log.Logger.Info("need larger space, recreate")
			s.Reset()
			s.Remove()
			retry++
			continue
		}

		if errors.Is(err, syscall.EEXIST) {
			log.Logger.Debug("shm already exist, open only")
			return s.OpenOnly()
		}

		break
	}

	if err != nil {
		log.Logger.Error("create shm failed", "err", err)
		return false
	}

	s.managedShm, err = unix.SysvShmAttach(s.shmID, 0, 0)
	if err != nil {
		log.Logger.Error("attach shm failed", "err", err)
		unix.SysvShmCtl(s.shmID, unix.IPC_RMID, nil)
		return false
	}

	if !s.castState() {
		return false
	}
	s.state.ceilingMsgSize.Store(s.conf.CeilingMsgSize())
	s.conf.Update(s.state.CeilingMsgSize())

	if !s.castBlocks() {
		return false
	}

	if !s.castBlockBufs() {
		return false
	}

	s.state.IncreaseReferenceCounts()
	s.init = true
	log.Logger.Debug("open or create true")
	return true
}

func (s *XSISegment) OpenOnly() bool {
	if s.init {
		return true
	}

	var err error
	s.shmID, err = unix.SysvShmGet(s.key, 0, 0644)
	if err != nil {
		log.Logger.Error("get shm failed", "err", err)
		return false
	}

	s.managedShm, err = unix.SysvShmAttach(s.shmID, 0, 0)
	if err != nil {
		log.Logger.Error("attach shm failed", "err", err)
		return false
	}

	if !s.castState() {
		return false
	}
	s.conf.Update(s.state.CeilingMsgSize())

	if !s.castBlocks() {
		return false
	}

	if !s.castBlockBufs() {
		return false
	}

	s.state.IncreaseReferenceCounts()
	s.init = true
	log.Logger.Debug("open only true")

	return true
}

func (s *XSISegment) castState() bool {
	state, err := NewStateFromShm(s.managedShm)
	if err != nil {
		log.Logger.Error("create state failed", "err", err)
		unix.SysvShmDetach(s.managedShm)
		unix.SysvShmCtl(s.shmID, unix.IPC_RMID, nil)
		return false
	}
	s.state = state

	return true
}

func (s *XSISegment) castBlocks() bool {
	var offset uint = StateSize
	var i uint32
	for i = 0; i < s.conf.BlockNum(); i++ {
		block, err := NewBlockFromShm(s.managedShm[offset : offset+BlockSize])
		if err != nil {
			log.Logger.Error("create block failed", "err", err)
			unix.SysvShmDetach(s.managedShm)
			unix.SysvShmCtl(s.shmID, unix.IPC_RMID, nil)
			return false
		}
		s.blocks = append(s.blocks, block)
		offset += BlockSize
	}

	return true
}

func (s *XSISegment) castBlockBufs() bool {
	offset := uint(StateSize + uint(s.conf.BlockNum())*uint(BlockSize))
	var i uint32
	for i = 0; i < s.conf.BlockNum(); i++ {
		blockBuf := unsafe.Slice(&s.managedShm[offset], s.conf.BlockBuffSize())
		if blockBuf == nil {
			break
		}
		s.blockBufLock.Lock()
		s.blockBufAddrs[i] = blockBuf
		s.blockBufLock.Unlock()
		offset += uint(s.conf.BlockBuffSize())
	}

	if i != s.conf.BlockNum() {
		log.Logger.Error("create block buf failed")
		s.blockBufLock.Lock()
		s.blockBufAddrs = make(map[uint32][]byte)
		s.blockBufLock.Unlock()
		unix.SysvShmDetach(s.managedShm)
		unix.SysvShmCtl(s.shmID, unix.IPC_RMID, nil)
		return false
	}

	return true
}

func (s *XSISegment) Remove() bool {
	var err error
	s.shmID, err = unix.SysvShmGet(s.key, 0, 0644)
	if err != nil {
		log.Logger.Error("remove shm failed", "err", err)
		return false
	}

	if _, err := unix.SysvShmCtl(s.shmID, unix.IPC_RMID, nil); err != nil {
		log.Logger.Error("remove shm failed", "err", err)
		return false
	}

	log.Logger.Debug("remove success")
	return true
}

func (s *XSISegment) Reset() {
	s.state = nil
	s.blocks = nil
	s.blockBufLock.Lock()
	s.blockBufAddrs = make(map[uint32][]byte)
	s.blockBufLock.Unlock()

	if s.managedShm != nil {
		unix.SysvShmDetach(s.managedShm)
		s.managedShm = nil
	}
}
