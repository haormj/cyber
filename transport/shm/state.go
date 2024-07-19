package shm

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

const StateSize = 32

type State struct {
	needRemap      *atomic.Bool
	seq            *atomic.Uint32
	referenceCount *atomic.Uint32
	ceilingMsgSize *atomic.Uint64
}

func NewStateFromShm(buf []byte) (*State, error) {
	state := &State{}
	state.needRemap = (*atomic.Bool)(unsafe.Pointer(&buf[8]))
	state.seq = (*atomic.Uint32)(unsafe.Pointer(&buf[12]))
	state.referenceCount = (*atomic.Uint32)(unsafe.Pointer(&buf[16]))
	state.ceilingMsgSize = (*atomic.Uint64)(unsafe.Pointer(&buf[24]))
	if state.needRemap == nil || state.seq == nil || state.referenceCount == nil || state.ceilingMsgSize == nil {
		return nil, errors.New("create state failed")
	}

	return state, nil
}

func (s *State) DecreaseReferenceCounts() {
	for {
		currentReferenceCount := s.referenceCount.Load()
		if currentReferenceCount == 0 {
			return
		}

		if s.referenceCount.CompareAndSwap(currentReferenceCount, currentReferenceCount-1) {
			break
		}
	}
}

func (s *State) IncreaseReferenceCounts() {
	s.referenceCount.Add(1)
}

func (s *State) FetchAddSeq(diff uint32) uint32 {
	return s.seq.Add(diff) - diff
}

func (s *State) Seq() uint32 {
	return s.seq.Load()
}

func (s *State) SetNeedRemap(need bool) {
	s.needRemap.Store(need)
}

func (s *State) NeedRemap() bool {
	return s.needRemap.Load()
}

func (s *State) CeilingMsgSize() uint64 {
	return s.ceilingMsgSize.Load()
}

func (s *State) ReferenceCounts() uint32 {
	return s.referenceCount.Load()
}
