package main

import (
	"context"
	pb "github.com/Kelp710/Go_gRPC/proto"
)

func(s *helloServer) Hello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{Message: "Hello from the Server!"}, nil
}