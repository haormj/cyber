package main

import (
	"log"
	"time"

	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/service"
	"google.golang.org/protobuf/proto"
)

func main() {
	client := service.NewClient[*pb.SimpleMessage, *pb.SimpleMessage]("client", "/test/service")
	if !client.Init() {
		log.Fatalln("client init failed")
	}

	if err := client.AsyncSendRequest(&pb.SimpleMessage{
		Integer: proto.Int32(2),
		Text:    proto.String("client async"),
	}, func(response *pb.SimpleMessage) {
		log.Println("async", response.GetInteger(), response.GetText())
	}); err != nil {
		log.Fatalln(err)
	}

	response, err := client.SendRequest(&pb.SimpleMessage{
		Integer: proto.Int32(1),
		Text:    proto.String("client"),
	}, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response.GetInteger(), response.GetText())
}
