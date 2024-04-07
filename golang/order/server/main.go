package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "test"

	"google.golang.org/grpc"
)

const port = 8000

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Name: fmt.Sprintf("Hello: %v!", in.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *server {
	return &server{}
}
