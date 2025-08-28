package main

import (
	pb "github/sunil8777/go-gRPC/proto"
	"log"
	"time"
)

func (s *server) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetServices_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error in sending data; %v", err)
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
