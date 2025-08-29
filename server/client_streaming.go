package main

import (
	pb "github/sunil8777/go-gRPC/proto"
	"io"
	"log"
)

func (s *server) SayHelloClientStreaming(stream pb.GreetServices_SayHelloClientStreamingServer) error {
	var msg []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed in sending the data: %v", err)
		}
		log.Printf("Got the request with name: %v", req.Name)
		msg = append(msg, "Hello "+req.Name)
	}
	res := &pb.MessagesList{
		Messages: msg,
	}

	return stream.SendAndClose(res)
}
