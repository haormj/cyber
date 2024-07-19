package main

import (
	"log"

	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/service"
	"google.golang.org/protobuf/proto"
)

func main() {
	svc := service.NewService[*pb.SimpleMessage, *pb.SimpleMessage]("service", "/test/service",
		func(request *pb.SimpleMessage, response **pb.SimpleMessage) {
			log.Println(request.GetInteger(), request.GetText())
			*response = &pb.SimpleMessage{
				Integer: proto.Int32(100),
				Text:    proto.String("service"),
			}
		})
	if !svc.Init() {
		log.Fatalln("service init failed")
	}

	select {}
}
