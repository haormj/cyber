package main

import (
	"log"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	r, err := node.NewReader[*pb.SimpleMessage](&pb.RoleAttributes{
		HostName:    proto.String(common.GlobalDataInstance.HostName()),
		HostIp:      proto.String(common.GlobalDataInstance.HostIP()),
		ProcessId:   proto.Int32(int32(common.GlobalDataInstance.ProcessID())),
		NodeName:    proto.String("reader"),
		NodeId:      proto.Uint64(common.Hash([]byte("reader"))),
		ChannelName: proto.String("/test/writer"),
		ChannelId:   proto.Uint64(common.Hash([]byte("/test/writer"))),
	}, func(msg *pb.SimpleMessage) {
		log.Println(msg.GetInteger(), msg.GetText())
	}, 1)
	if err != nil {
		log.Fatalln(err)
	}

	_ = r

	select {}
}
