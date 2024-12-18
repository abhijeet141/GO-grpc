package main

import (
	"context"
	pb "grpc-server/proto"
	"io"
	"log"
)

func callSayHelloServerSideStream(client pb.GreetServiceClient, names *pb.NamesList) {
	res, err := client.SayHelloServerSideStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for {
		message, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message.Message)
	}
	log.Printf("Streaming Finished")
}
