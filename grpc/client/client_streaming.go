package main

import (
	"context"
	pb "grpc-server/proto"
	"log"
	"time"
)

func callSayHelloClientSideStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client Side Streaming Started")
	stream, err := client.SayHelloClientSideStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for _, name := range names.Names {
		err := stream.Send(&pb.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent name: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while receiving %v", err)
		return
	}
	log.Println("Client Side Streaming Finished")
	log.Printf("Message: %v", res.Messages)
}
