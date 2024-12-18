package main

import (
	"context"
	pb "grpc-server/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bi-Directional Streaming Started")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	ch := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming: %v", err)
			}
			log.Println(message.Message)
		}
		close(ch)
	}()
	for _, name := range names.Names {
		err := stream.Send(&pb.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent Name: %v", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-ch
	log.Println("Bi-Directional Streaming Finished")
}
