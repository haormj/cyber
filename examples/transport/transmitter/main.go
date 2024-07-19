package main

import (
	"log"
	"time"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport"
	"google.golang.org/protobuf/proto"
)

func main() {
	transmitter, err := transport.CreateTransmitter[*pb.SimpleMessage](
		&pb.RoleAttributes{
			HostName:    proto.String(common.GlobalDataInstance.HostName()),
			HostIp:      proto.String(common.GlobalDataInstance.HostIP()),
			ProcessId:   proto.Int32(int32(common.GlobalDataInstance.ProcessID())),
			NodeName:    proto.String("transmiter"),
			NodeId:      proto.Uint64(common.Hash([]byte("transmiter"))),
			ChannelName: proto.String("/test/transmiter"),
			ChannelId:   proto.Uint64(common.Hash([]byte("/test/transmiter"))),
		},
		pb.OptionalMode_SHM,
	)
	if err != nil {
		log.Fatalln(err)
	}

	i := 0
	for {
		m := &pb.SimpleMessage{
			Integer: proto.Int32(int32(i)),
			Text:    proto.String("hello"),
		}
		if err := transmitter.Transmit(m); err != nil {
			log.Println("transmit failed", err)
		}
		i++
		log.Println(m.GetInteger(), m.GetText())
		time.Sleep(time.Second)
	}
}
