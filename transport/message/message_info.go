package message

import (
	"encoding/binary"

	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/transport/identity"
)

const MessageInfoSize = 2*identity.ID_SIZE + 8

type MessageInfo struct {
	SenderID  *identity.Identity
	ChannelID uint64
	SeqNum    uint64
	SpareID   *identity.Identity
}

func NewMessageInfo() *MessageInfo {
	return &MessageInfo{
		SenderID: identity.NewIdentity(false),
		SpareID:  identity.NewIdentity(false),
	}
}

func CloneMessageInfo(msgInfo *MessageInfo) *MessageInfo {
	return &MessageInfo{
		SenderID:  identity.CloneIdentity(msgInfo.SenderID),
		ChannelID: msgInfo.ChannelID,
		SeqNum:    msgInfo.SeqNum,
		SpareID:   identity.CloneIdentity(msgInfo.SenderID),
	}
}

func (i *MessageInfo) Serialize() []byte {
	var data []byte
	data = append(data, i.SenderID.Data()...)
	data = binary.LittleEndian.AppendUint64(data, i.SeqNum)
	data = append(data, i.SpareID.Data()...)

	return data
}

func (i *MessageInfo) Deserialize(data []byte) bool {
	if len(data) != int(MessageInfoSize) {
		log.Logger.Warn("size mismatch", "given", len(data), "target", MessageInfoSize)
		return false
	}

	i.SenderID.SetData(data[:8])
	i.SeqNum = binary.LittleEndian.Uint64(data[8:16])
	i.SpareID.SetData(data[16:])
	return true
}
