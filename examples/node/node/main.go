package main

import (
	"log"
	"time"

	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	c := node.NewChannelCreator[*pb.SimpleMessage](node.NewChannelNode("n1"))
	if _, err := c.CreateReader("/test/hello", func(msg *pb.SimpleMessage) {
		log.Println("reader1", msg.GetInteger(), msg.GetText())
	}); err != nil {
		log.Fatalln(err)
	}

	if _, err := c.CreateReader("/test/hello", func(msg *pb.SimpleMessage) {
		log.Println("reader2", msg.GetInteger(), msg.GetText())
	}); err != nil {
		log.Fatalln(err)
	}

	w, err := c.CreateWriter("/test/hello")
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
		log.Println("writer", m.GetInteger(), m.GetText())
		time.Sleep(time.Second)
	}
}
