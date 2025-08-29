package main

import (
	"context"
	pb "github/sunil8777/go-gRPC/proto"
	"log"
	"time"
)

func callSayHelloClientStreaming(client pb.GreetServicesClient, names *pb.NameList){
	log.Printf("client streaimg started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("unable to send name: %v",names)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("unable to send name: %v",err)
		}
		time.Sleep(time.Second*2)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Streaming finished")
	if err != nil {
		log.Fatalf("error while receiving:  %v",err)
	}
	log.Printf("%v",res.Messages)
	
}