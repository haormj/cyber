package base

import "sync"

type Callback func(...any)

type Signal struct {
	slots []*Slot
	mutex sync.Mutex
}

func NewSignal() *Signal {
	return &Signal{}
}

func (s *Signal) Call(args ...any) {
	var local []*Slot
	s.mutex.Lock()
	for _, slot := range s.slots {
		local = append(local, NewSlot(slot.connected, slot.callback))
	}
	s.mutex.Unlock()

	for _, slot := range local {
		slot.Call(args...)
	}

	s.ClearDisconnectedSlots()
}

func (s *Signal) Connect(callback Callback) *Connection {
	slot := NewSlot(true, callback)
	s.mutex.Lock()
	s.slots = append(s.slots, slot)
	s.mutex.Unlock()

	return NewConnection(slot, s)
}

func (s *Signal) Disconnect(conn Connection) bool {
	var find bool = false
	s.mutex.Lock()
	for _, slot := range s.slots {
		if conn.HasSlot(slot) {
			find = true
			slot.Disconnect()
		}
	}
	s.mutex.Unlock()

	if find {
		s.ClearDisconnectedSlots()
	}

	return find
}

func (s *Signal) ClearDisconnectedSlots() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var newSlots []*Slot
	for _, slot := range s.slots {
		if slot.connected {
			newSlots = append(newSlots, slot)
		}
	}

	s.slots = newSlots
}

func (s *Signal) DisconnectAllSlots() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, slot := range s.slots {
		slot.Disconnect()
	}

	s.slots = nil
}

type Connection struct {
	slot   *Slot
	signal *Signal
}

func NewConnection(slot *Slot, signal *Signal) *Connection {
	return &Connection{
		slot:   slot,
		signal: signal,
	}
}

func (c *Connection) HasSlot(slot *Slot) bool {
	if c.slot != nil && slot != nil {
		return c.slot == slot
	}
	return false
}

func (c *Connection) IsConnected() bool {
	if c.slot != nil {
		return c.slot.connected
	}
	return false
}

func (c *Connection) Disconnect() bool {
	if c.signal != nil && c.slot != nil {
		return c.signal.Disconnect(*c)
	}

	return false
}

type Slot struct {
	callback  Callback
	connected bool
}

func NewSlot(connected bool, callback Callback) *Slot {
	return &Slot{
		connected: connected,
		callback:  callback,
	}
}

func (s *Slot) Disconnect() {
	s.connected = false
}

func (s *Slot) Connected() bool {
	return s.connected
}

func (s *Slot) Call(args ...any) {
	s.callback(args...)
}
