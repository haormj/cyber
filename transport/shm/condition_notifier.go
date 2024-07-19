package shm

import (
	"errors"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/haormj/cyber/log"

	"github.com/haormj/cyber/common"
	"golang.org/x/sys/unix"
)

const kBufLength = 4096
const IndicatorSize = 8 + kBufLength*ReadableInfoSize + 8*kBufLength

var ConditionNotifierInstance Notifier = NewConditionNotifier()

type Indicator struct {
	nextSeq *atomic.Uint64
	infos   [kBufLength][]byte
	seqs    [kBufLength]*uint64
}

type ConditionNotifier struct {
	key        int
	managedShm []byte
	shmSize    uint64
	indicator  *Indicator
	nextSeq    uint64
	isShutdown atomic.Bool
}

func NewConditionNotifier() *ConditionNotifier {
	n, err := NewConditionNotifierE()
	if err != nil {
		panic(err)
	}

	return n
}

func NewConditionNotifierE() (*ConditionNotifier, error) {
	n := &ConditionNotifier{}
	n.key = int(common.Hash([]byte("/apollo/cyber/transport/shm/notifier")))
	log.Logger.Debug("condition notifier key", "key", n.key)
	n.shmSize = IndicatorSize

	if !n.init() {
		n.isShutdown.Store(true)
		return nil, errors.New("fail to init condition notifier")
	}
	n.nextSeq = n.indicator.nextSeq.Load()
	log.Logger.Debug("next_seq", "nextSeq", n.nextSeq)

	return n, nil
}

func (n *ConditionNotifier) init() bool {
	return n.openOrCreate()
}

func (n *ConditionNotifier) openOrCreate() bool {
	// create managed shm
	var retry int
	var shmID int
	var err error
	for retry < 2 {
		shmID, err = unix.SysvShmGet(n.key, int(n.shmSize), 0644|unix.IPC_CREAT|unix.IPC_EXCL)
		if err == nil {
			break
		}

		if errors.Is(err, syscall.EINVAL) {
			log.Logger.Info("need larger space, recreate")
			n.reset()
			n.remove()
			retry++
			continue
		}

		if errors.Is(err, syscall.EEXIST) {
			log.Logger.Debug("shm already exist, open only")
			if n.openOnly() {
				return true
			}
			retry++
			continue
		}

		break
	}

	if err != nil {
		log.Logger.Error("create shm failed", "retry", retry, "err", err)
		return false
	}

	n.managedShm, err = unix.SysvShmAttach(shmID, 0, 0)
	if err != nil {
		log.Logger.Error("attach shm failed")
		unix.SysvShmCtl(shmID, unix.IPC_RMID, nil)
		return false
	}

	if !n.castIndicator() {
		return false
	}

	log.Logger.Debug("open or create true")
	return true
}

func (n *ConditionNotifier) openOnly() bool {
	// get managed shm
	shmID, err := unix.SysvShmGet(n.key, 0, 0644)
	if err != nil {
		log.Logger.Error("get shm failed", "err", err)
		return false
	}

	n.managedShm, err = unix.SysvShmAttach(shmID, 0, 0)
	if err != nil {
		log.Logger.Error("attach shm failed", "err", err)
		unix.SysvShmCtl(shmID, unix.IPC_RMID, nil)
		return false
	}

	if uint64(len(n.managedShm)) != n.shmSize {
		log.Logger.Error("shm size", "got", len(n.managedShm), "expected", n.shmSize)
		unix.SysvShmDetach(n.managedShm)
		unix.SysvShmCtl(shmID, unix.IPC_RMID, nil)
		return false
	}

	if !n.castIndicator() {
		return false
	}

	log.Logger.Debug("open only true")
	return true
}

func (n *ConditionNotifier) castIndicator() bool {
	n.indicator = &Indicator{}
	offset := 0
	n.indicator.nextSeq = (*atomic.Uint64)(unsafe.Pointer(&n.managedShm[offset]))
	offset += 8

	for i := 0; i < kBufLength; i++ {
		n.indicator.infos[i] = unsafe.Slice(&n.managedShm[offset], ReadableInfoSize)
		offset += ReadableInfoSize
	}

	for i := 0; i < kBufLength; i++ {
		n.indicator.seqs[i] = (*uint64)(unsafe.Pointer(&n.managedShm[offset]))
		offset += 8
	}

	if n.indicator == nil {
		log.Logger.Error("get indicator failed")
		unix.SysvShmDetach(n.managedShm)
		return false
	}

	return true
}

func (n *ConditionNotifier) remove() bool {
	shmID, err := unix.SysvShmGet(n.key, 0, 0644)
	if err != nil {
		log.Logger.Error("remove shm failed", "err", err)
		return false
	}

	if _, err := unix.SysvShmCtl(shmID, unix.IPC_RMID, nil); err != nil {
		log.Logger.Error("remove shm failed", "err", err)
		return false
	}

	log.Logger.Debug("remove success")
	return true
}

func (n *ConditionNotifier) reset() {
	n.indicator = nil
	if n.managedShm != nil {
		unix.SysvShmDetach(n.managedShm)
		n.managedShm = nil
	}
}

func (n *ConditionNotifier) Type() string {
	return "condition"
}

func (n *ConditionNotifier) Shutdown() {
	if n.isShutdown.Swap(true) {
		return
	}

	time.Sleep(100 * time.Millisecond)
	n.reset()
}

func (n *ConditionNotifier) Notify(info *ReadableInfo) bool {
	if info == nil {
		log.Logger.Error("info nil")
		return false
	}

	if n.isShutdown.Load() {
		log.Logger.Debug("notifier is shutdown")
		return false
	}

	currentSeq := n.indicator.nextSeq.Add(1) - 1
	idx := currentSeq % kBufLength
	copy(n.indicator.infos[idx][:], info.Serialize())
	*n.indicator.seqs[idx] = currentSeq

	return true
}

func (n *ConditionNotifier) Listen(timeoutMs int, info *ReadableInfo) bool {
	if info == nil {
		log.Logger.Error("info nil")
		return false
	}

	if n.isShutdown.Load() {
		log.Logger.Debug("notifier is shutdown")
		return false
	}

	timeoutUs := timeoutMs * 1000
	for !n.isShutdown.Load() {
		seq := n.indicator.nextSeq.Load()
		if seq != n.nextSeq {
			idx := n.nextSeq % kBufLength
			actualSeq := *n.indicator.seqs[idx]
			if actualSeq >= n.nextSeq {
				n.nextSeq = actualSeq
				info.Deserialize(n.indicator.infos[idx][:])
				n.nextSeq++
				return true
			} else {
				// log.Logger.Debug("seq is writing, can not read now", "nextSeq", n.nextSeq, "seq", seq, "actualSeq", actualSeq)
			}

		}

		if timeoutUs > 0 {
			time.Sleep(50 * time.Microsecond)
			timeoutUs -= 50
		} else {
			return false
		}
	}

	return false
}
