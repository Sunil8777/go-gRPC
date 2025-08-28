package main

import (
	"log"
	"net"
	pb "github/sunil8777/go-gRPC/proto"
	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreetServicesServer
}

const (
	port = ":8000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServicesServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v",err)
	}
	log.Fatal("server started")
}