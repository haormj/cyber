package shm

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/haormj/cyber/log"

	"github.com/haormj/cyber/common"
)

var MulticastNotifierInstance Notifier = NewMulticastNotifier()

type MulticastNotifier struct {
	isShutdown atomic.Bool
	notifyConn *net.UDPConn
	listenConn *net.UDPConn
}

func NewMulticastNotifier() *MulticastNotifier {
	n, err := NewMulticastNotifierE()
	if err != nil {
		panic(err)
	}

	return n
}

func NewMulticastNotifierE() (*MulticastNotifier, error) {
	n := &MulticastNotifier{}
	if err := n.init(); err != nil {
		n.Shutdown()
		return nil, err
	}

	return n, nil
}

func (n *MulticastNotifier) init() error {
	mcastIP := "239.255.0.100"
	var mcastPort uint32 = 8888

	config := common.GlobalDataInstance.Config()
	if config != nil && config.TransportConf != nil &&
		config.TransportConf.ShmConf != nil &&
		config.TransportConf.ShmConf.ShmLocator != nil {
		mcastIP = config.TransportConf.ShmConf.ShmLocator.GetIp()
		mcastPort = config.TransportConf.ShmConf.ShmLocator.GetPort()
	}

	notifyAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", mcastIP, mcastPort))
	if err != nil {
		return err
	}

	n.notifyConn, err = net.DialUDP("udp", nil, notifyAddr)
	if err != nil {
		return err
	}

	listenAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", mcastIP, mcastPort))
	if err != nil {
		return err
	}

	n.listenConn, err = net.ListenMulticastUDP("udp", nil, listenAddr)
	if err != nil {
		return err
	}

	return nil
}

func (n *MulticastNotifier) Type() string {
	return "multicast"
}

func (n *MulticastNotifier) Notify(info *ReadableInfo) bool {
	if n.isShutdown.Load() {
		return false
	}

	if info == nil {
		log.Logger.Error("info nil")
		return false
	}

	data := info.Serialize()
	if _, err := n.notifyConn.Write(data); err != nil {
		return false
	}

	return true
}

func (n *MulticastNotifier) Listen(timeoutMs int, info *ReadableInfo) bool {
	if n.isShutdown.Load() {
		return false
	}

	if info == nil {
		log.Logger.Error("info nil")
		return false
	}

	if err := n.listenConn.SetReadDeadline(time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)); err != nil {
		log.Logger.Error("listenConn.SetReadDeadline error", "err", err)
		return false
	}
	data := make([]byte, 32)
	if _, err := n.listenConn.Read(data); err != nil {
		if opError, ok := err.(*net.OpError); ok && opError.Timeout() {
			log.Logger.Debug("timeout, no readableinfo")
		} else {
			log.Logger.Error("read error", "err", err)
		}
		return false
	}

	return info.Deserialize(data)
}

func (n *MulticastNotifier) Shutdown() {
	if n.isShutdown.Swap(true) {
		return
	}

	if n.notifyConn != nil {
		n.notifyConn.Close()
	}

	if n.listenConn != nil {
		n.listenConn.Close()
	}
}
