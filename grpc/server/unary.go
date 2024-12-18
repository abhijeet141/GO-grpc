package main

import (
	"context"
	pb "grpc-server/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello Abhijeet",
	}, nil
}
