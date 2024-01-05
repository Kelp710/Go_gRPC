package main

import (
	"log"
	"time"

	pb "github.com/Kelp710/Go_gRPC/proto"
)

func (s *helloServer) HelloServerStream(req *pb.NamesList, stream pb.GreetService_HelloServerStreamServer) error {
	log.Printf("Received name: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: "Hello " + name}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
