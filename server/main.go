package main

import (
	"log"
	"net"

	pb "github.com/robertojrojas/grpc-go-example/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &helloserver{})
	log.Printf("server started at %v" ,lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
