package main

import (
	pb "grpc-server/proto"

	"io"
	"log"
	"time"
)

func (s *HelloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		err = stream.Send(&pb.HelloResponse{
			Message: req.Name,
		})
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
