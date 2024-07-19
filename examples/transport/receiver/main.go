package main

import (
	"log"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport"
	"github.com/haormj/cyber/transport/message"
	"google.golang.org/protobuf/proto"
)

func main() {
	_, err := transport.CreateReceiver[*pb.SimpleMessage](
		&pb.RoleAttributes{
			HostName:    proto.String(common.GlobalDataInstance.HostName()),
			HostIp:      proto.String(common.GlobalDataInstance.HostIP()),
			ProcessId:   proto.Int32(int32(common.GlobalDataInstance.ProcessID())),
			NodeName:    proto.String("receiver"),
			NodeId:      proto.Uint64(common.Hash([]byte("receiver"))),
			ChannelName: proto.String("/test/transmiter"),
			ChannelId:   proto.Uint64(common.Hash([]byte("/test/transmiter"))),
		},
		pb.OptionalMode_SHM,
		func(msg *pb.SimpleMessage, msgInfo *message.MessageInfo, attr *pb.RoleAttributes) {
			log.Println(msg.GetInteger(), msg.GetText())
		})
	if err != nil {
		log.Fatalln(err)
	}

	select {}
}
