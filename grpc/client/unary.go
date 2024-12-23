package main

import (
	"context"
	pb "grpc-server/proto"
	"log"
	"time"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%v", res)
}
