package main

import (
	"context"
	pb "github/sunil8777/go-gRPC/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectional(client pb.GreetServicesClient, names *pb.NameList ){
	log.Println("Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	wait := make(chan int)
	
	go func(){
		for {
			message, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil {
				log.Fatalf("Error in receiving: %v",err)
			}
			log.Println(message.Message)
		}
		close(wait)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("unable to send name: %v",err)
		}
		time.Sleep(time.Second*2)
	}
	stream.CloseSend()
	<-wait
	log.Fatalln("Streaming finished")
	
}