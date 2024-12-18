package main

import (
	pb "grpc-server/proto"
	"log"
	"time"
)

func (s *HelloServer) SayHelloServerSideStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerSideStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + name,
		})
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
