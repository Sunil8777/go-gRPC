package main

import (
	"log"
	pb "github/sunil8777/go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServicesClient(conn)

	name := &pb.NameList{
		Names: []string{"ab","cd","ef"},
	}

	// callSayHelloCall(client)
	// callSayHelloServerStreaming(client, name)
	// callSayHelloClientStreaming(client, name)
	callHelloBidirectional(client,name)
}
