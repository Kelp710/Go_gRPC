package main

import (
	"io"
	"log"

	pb "github.com/Kelp710/Go_gRPC/proto"
)

func (s *helloServer) HellowClientStream(stream pb.GreetService_HellowClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we've finished reading the client stream.
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Received message %s", req.Name)
		messages = append(messages, "hELOOO", req.Name)
	}
}
