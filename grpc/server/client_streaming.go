package main

import (
	pb "grpc-server/proto"
	"io"
	"log"
)

func (s *HelloServer) SayHelloClientSideStreaming(stream pb.GreetService_SayHelloClientSideStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, req.Name)
	}
}
