package main

import (
	"context"
	"log"
	"net"

	common "github.com/vule96/commons"
	"google.golang.org/grpc"
)

var (
	// serviceName = "orders"
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())

	log.Println("GRPC Server started at: ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
