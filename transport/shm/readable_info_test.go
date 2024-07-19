package shm

import "testing"

func TestReadableInfo(t *testing.T) {
	info1 := ReadableInfo{
		HostID:     1,
		BlockIndex: 2,
		ChannelID:  3,
	}
	info2 := ReadableInfo{}
	if !info2.Deserialize(info1.Serialize()) {
		t.Fatal("info deserialize failed")
	}

	if info1.HostID != info2.HostID ||
		info1.BlockIndex != info2.BlockIndex ||
		info1.ChannelID != info2.ChannelID {
		t.Fatal("info1", "hostID", info1.HostID, "blockIndex", info1.BlockIndex, "channelID", info1.ChannelID,
			"info2", "hostID", info2.HostID, "blockIndex", info2.BlockIndex, "channelID", info2.ChannelID,
		)
	}
}
