package main

import (
	"log"
	"time"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	w, err := node.NewWriter[*pb.SimpleMessage](&pb.RoleAttributes{
		HostName:    proto.String(common.GlobalDataInstance.HostName()),
		HostIp:      proto.String(common.GlobalDataInstance.HostIP()),
		ProcessId:   proto.Int32(int32(common.GlobalDataInstance.ProcessID())),
		NodeName:    proto.String("writer"),
		NodeId:      proto.Uint64(common.Hash([]byte("writer"))),
		ChannelName: proto.String("/test/writer"),
		ChannelId:   proto.Uint64(common.Hash([]byte("/test/writer"))),
	})
	if err != nil {
		log.Fatalln(err)
	}

	i := 0
	for {
		m := &pb.SimpleMessage{
			Integer: proto.Int32(int32(i)),
			Text:    proto.String("hello"),
		}
		if err := w.Write(m); err != nil {
			log.Println("write failed", err)
		}
		i++
		log.Println(m.GetInteger(), m.GetText())
		time.Sleep(time.Second)
	}
}
