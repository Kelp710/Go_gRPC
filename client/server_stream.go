package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Kelp710/Go_gRPC/proto"
)

func callHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.HelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("error while calling HelloServerStream RPC: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream.
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from HelloServerStream: %s", message)
	}
	log.Printf("streaming finished")
}
