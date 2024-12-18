package main

import (
	"context"
	pb "grpc-server/proto"
	"log"

	"google.golang.org/grpc/metadata"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No Metadata")
	} else {
		os := data["os"]
		log.Println("Metadata: ", os[0])
	}
	return &pb.HelloResponse{
		Message: "Hello Abhijeet",
	}, nil
}
