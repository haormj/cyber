package shm

import "encoding/binary"

const ReadableInfoSize = 32

type ReadableInfo struct {
	padding    uint64
	HostID     uint64
	BlockIndex uint32
	ChannelID  uint64
}

func (r *ReadableInfo) Serialize() []byte {
	var b []byte
	b = binary.LittleEndian.AppendUint64(b, r.padding)
	b = binary.LittleEndian.AppendUint64(b, r.HostID)
	b = binary.LittleEndian.AppendUint32(b, r.BlockIndex)
	b = binary.LittleEndian.AppendUint32(b, 0)
	b = binary.LittleEndian.AppendUint64(b, r.ChannelID)
	return b
}

func (r *ReadableInfo) Deserialize(b []byte) bool {
	if len(b) != ReadableInfoSize {
		return false
	}

	r.padding = binary.LittleEndian.Uint64(b[:8])
	r.HostID = binary.LittleEndian.Uint64(b[8:16])
	r.BlockIndex = binary.LittleEndian.Uint32(b[16:20])
	r.ChannelID = binary.LittleEndian.Uint64(b[24:])

	return true
}
