package main

import (
	"context"
	"log"
	"time"

	pb "test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "localhost:8000"

func main() {
	// var opts []grpc.DialOption
	// conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error: %v", err)
	}
	defer conn.Close()

	gc := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	gr, err := gc.SayHello(ctx, &pb.HelloRequest{Name: "Ram"})
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("%v", gr.Name)
}
