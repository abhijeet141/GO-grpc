package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	os := "windows"
	ctx = metadata.AppendToOutgoingContext(ctx, "os", os)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		log.Printf("failed: %v", err)
	} else {
		log.Printf("Response received: %v", reply)
	}
	return err
}
