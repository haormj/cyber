package state

import (
	"os"
	"sync/atomic"
	"time"
)

type State uint8

const (
	STATE_UNINITIALIZED State = iota
	STATE_INITIALIZED
	STATE_SHUTTING_DOWN
	STATE_SHUTDOWN
)

var cyberState atomic.Uint32

func GetState() State {
	return State(cyberState.Load())
}

func SetState(s State) {
	cyberState.Store(uint32(s))
}

func OK() bool {
	return GetState() == STATE_INITIALIZED
}

func IsShutdown() bool {
	return GetState() == STATE_SHUTTING_DOWN || GetState() == STATE_SHUTDOWN
}

func WaitForShutdown() {
	for !IsShutdown() {
		time.Sleep(200 * time.Millisecond)
	}
}

func AsyncShutdown() {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		panic(err)
	}

	p.Signal(os.Interrupt)
}
