package main

import (
	"context"
	pb "github/sunil8777/go-gRPC/proto"
	"log"
	"time"
)

func sayHelloCall(client pb.GreetServicesClient) {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v",err)
	}
	log.Printf("%s",res.Message)
}
