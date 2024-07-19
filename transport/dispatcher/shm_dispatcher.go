package dispatcher

import (
	"math"
	"sync"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/shm"
)

var ShmDispatcherInstance = NewShmDispatcher()

type ShmDispatcher struct {
	*BaseDispatcher[*shm.ReadableBlock]
	hostID          uint64
	segments        sync.Map
	previousIndexes sync.Map
	notifier        shm.Notifier
	wg              sync.WaitGroup
}

func NewShmDispatcher() *ShmDispatcher {
	d := &ShmDispatcher{
		BaseDispatcher: NewBaseDispatcher[*shm.ReadableBlock](),
	}

	d.init()

	return d
}

func (d *ShmDispatcher) addSegment(selfAttr *pb.RoleAttributes) {
	channelID := selfAttr.GetChannelId()

	if _, ok := d.segments.Load(channelID); ok {
		return
	}

	d.segments.Store(channelID, shm.NewSegment(channelID))
	d.previousIndexes.Store(channelID, uint32(math.MaxUint32))
}

func (d *ShmDispatcher) readMessage(channelID uint64, blockIndex uint32) {
	log.Logger.Debug("reading sharedmem message from block", "channelID", channelID, "blockIndex", blockIndex)

	readableBlock := &shm.ReadableBlock{
		Index: blockIndex,
	}

	v, ok := d.segments.Load(channelID)
	if !ok {
		log.Logger.Warn("segment not find", "channelID", channelID)
		return
	}
	segment, ok := v.(shm.Segment)
	if !ok {
		log.Logger.Error("convert value to segment failed")
		return
	}

	if !segment.AcquireBlockToRead(readableBlock) {
		log.Logger.Warn("fail to acquire block",
			"channel", common.GlobalDataInstance.GetChannelByID(channelID), "index", blockIndex)
		return
	}

	msgInfo := message.NewMessageInfo()
	if msgInfo.Deserialize(readableBlock.Buf[readableBlock.Block.MsgSize() : readableBlock.Block.MsgSize()+readableBlock.Block.MsgInfoSize()]) {
		d.onMessage(channelID, readableBlock, msgInfo)
	} else {
		log.Logger.Error("msg info deserialize failed", "channel", common.GlobalDataInstance.GetChannelByID(channelID))
	}
	segment.ReleaseReadBlock(readableBlock)
}

func (d *ShmDispatcher) onMessage(channelID uint64, readableBlock *shm.ReadableBlock, msgInfo *message.MessageInfo) {
	if d.isShutdown.Load() {
		return
	}

	v, ok := d.msgListeners.Load(channelID)
	if ok {
		listenerHandler := v.(*message.ListenerHandler[*shm.ReadableBlock])
		listenerHandler.Run(readableBlock, msgInfo)
	} else {
		log.Logger.Error("cannot find handler", "channel", common.GlobalDataInstance.GetChannelByID(channelID), "channelID", channelID)
	}
}

func (d *ShmDispatcher) threadFunc() {
	var info shm.ReadableInfo
	for {
		if d.isShutdown.Load() {
			return
		}

		if !d.notifier.Listen(100, &info) {
			log.Logger.Debug("listen failed")
			continue
		}

		if info.HostID != d.hostID {
			log.Logger.Debug("shm readable info from other host")
			continue
		}

		if _, ok := d.segments.Load(info.ChannelID); !ok {
			continue
		}

		i, ok := d.previousIndexes.Load(info.ChannelID)
		if !ok {
			log.Logger.Debug("previous index not find", "channelID", info.ChannelID)
			continue
		}

		previousIndex, ok := i.(uint32)
		if !ok {
			log.Logger.Error("value convert to previous index failed")
			continue
		}

		if info.BlockIndex != 0 && info.BlockIndex != math.MaxUint32 {
			switch {
			case info.BlockIndex == previousIndex:
				log.Logger.Debug("receive same index", "blockIndex", info.BlockIndex, "channelID", info.ChannelID)
			case info.BlockIndex < previousIndex:
				log.Logger.Debug("receive previous message", "previousIndex", previousIndex, "blockIndex", info.BlockIndex)
			case info.BlockIndex-previousIndex > 1:
				log.Logger.Debug("receive jump message", "previousIndex", previousIndex, "blockIndex", info.BlockIndex)
			}
		}
		d.previousIndexes.Store(info.ChannelID, info.BlockIndex)

		d.readMessage(info.ChannelID, info.BlockIndex)
	}
}

func (d *ShmDispatcher) init() bool {
	d.hostID = common.Hash([]byte(common.GlobalDataInstance.HostIP()))
	d.notifier = shm.NewNotifier()
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		d.threadFunc()
	}()
	return true
}

func (d *ShmDispatcher) AddListener(selfAttr *pb.RoleAttributes, listener MessageListener[*shm.ReadableBlock]) {
	d.BaseDispatcher.AddListener(selfAttr, listener)
	d.addSegment(selfAttr)
}

func (d *ShmDispatcher) AddOppositeListener(selfAttr, oppoAttr *pb.RoleAttributes, listener MessageListener[*shm.ReadableBlock]) {
	d.BaseDispatcher.AddListener(selfAttr, listener)
	d.addSegment(selfAttr)
}

func (d *ShmDispatcher) Shutdown() {
	if d.isShutdown.Swap(true) {
		return
	}

	d.wg.Wait()
}
