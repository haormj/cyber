package transmitter

import (
	"errors"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/endpoint"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/shm"
	"google.golang.org/protobuf/proto"
)

type ShmTransmitter[M proto.Message] struct {
	*endpoint.Endpoint
	segment   shm.Segment
	channelID uint64
	hostID    uint64
	notifier  shm.Notifier
	seqNum    uint64
	msgInfo   *message.MessageInfo
}

func NewShmTransmitter[M proto.Message](attr *pb.RoleAttributes) (*ShmTransmitter[M], error) {
	if attr == nil {
		return nil, errors.New("attr nil")
	}

	t := &ShmTransmitter[M]{
		Endpoint:  endpoint.NewEndpoint(attr),
		channelID: attr.GetChannelId(),
		hostID:    common.Hash([]byte(attr.GetHostIp())),
		msgInfo:   message.NewMessageInfo(),
	}

	t.msgInfo.SenderID = t.ID()
	t.msgInfo.SeqNum = t.seqNum

	return t, nil
}

func (t *ShmTransmitter[M]) SeqNum() uint64 {
	return t.seqNum
}

func (t *ShmTransmitter[M]) NextSeqNum() uint64 {
	t.seqNum++
	return t.seqNum
}

func (t *ShmTransmitter[M]) Enable() {
	if t.Enabled {
		return
	}

	t.segment = shm.NewSegment(t.channelID)
	t.notifier = shm.NewNotifier()
	t.Enabled = true
}

func (t *ShmTransmitter[M]) Disable() {
	if t.Enabled {
		t.segment = nil
		t.notifier = nil
		t.Enabled = false
	}
}

func (t *ShmTransmitter[M]) Transmit(msg M) error {
	t.msgInfo.SeqNum = t.NextSeqNum()
	return t.TransmitWithMessageInfo(msg, t.msgInfo)
}

func (t *ShmTransmitter[M]) TransmitWithMessageInfo(msg M, msgInfo *message.MessageInfo) error {
	if !t.Enabled {
		return errors.New("transmiter not enable")
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	var msgSize uint64 = uint64(len(buf))

	var wb shm.WritableBlock
	if !t.segment.AcquireBlockToWrite(msgSize, &wb) {
		return errors.New("acquire block to write failed")
	}

	copy(wb.Buf[:msgSize], buf)
	wb.Block.SetMsgSize(msgSize)

	msgInfoData := msgInfo.Serialize()
	copy(wb.Buf[msgSize:], msgInfoData)
	wb.Block.SetMsgInfoSize(uint64(message.MessageInfoSize))
	t.segment.ReleaseWrittenBlock(&wb)

	readableInfo := &shm.ReadableInfo{
		HostID:     t.hostID,
		BlockIndex: wb.Index,
		ChannelID:  t.channelID,
	}

	if !t.notifier.Notify(readableInfo) {
		return errors.New("notify failed")
	}

	log.Logger.Debug("writing sharedmem message", "channelID", t.channelID, "blockIndex", wb.Index)

	return nil
}
