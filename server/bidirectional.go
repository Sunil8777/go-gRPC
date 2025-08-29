package main

import (
	pb "github/sunil8777/go-gRPC/proto"
	"io"
	"log"
)

func (s *server) SayHelloBidirectionalStreaming(stream pb.GreetServices_SayHelloBidirectionalStreamingServer) error {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return  nil
		}
		if err != nil {
			log.Fatalf("Unable to recieve the names: %v",err)
		}
		log.Println("Got the request with name: ",res.Name)

		result := &pb.HelloResponse{
			Message: "Hello " + res.Name,
		}
		if err := stream.Send(result); err != nil {
			return err
		}
	}
}