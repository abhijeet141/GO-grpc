package main

import (
	pb "grpc-server/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	CallSayHello(client)
	names := &pb.NamesList{
		Names: []string{"Abhi", "Ajay", "Aditya"},
	}
	callSayHelloServerSideStream(client, names)
	callSayHelloClientSideStream(client, names)
}
