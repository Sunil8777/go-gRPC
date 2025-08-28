package main

import (
	"context"
	pb "github/sunil8777/go-gRPC/proto"
)

func (s *server) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	},nil
}
