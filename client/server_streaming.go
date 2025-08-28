package main

import (
	"context"
	pb "github/sunil8777/go-gRPC/proto"
	"io"
	"log"
)

func sayHelloServerStreaming(client pb.GreetServicesClient, name *pb.NameList) {
	log.Println("Streaming started")

	res, err := client.SayHelloServerStreaming(context.Background(), name)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		message, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
