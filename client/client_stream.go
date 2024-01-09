package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kelp710/Go_gRPC/proto"
)

func callHellowClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.HellowClientStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Could not send a name: %v", err)
		}
		log.Printf("Sent %s\n", name)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client stream finished")
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("%v", res.Messages)

}
