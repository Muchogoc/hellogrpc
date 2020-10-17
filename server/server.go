package main

import (
	"log"
	"net"

	"github.com/Muchogoc/hellogrpc/api"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := api.ImplementedPingServer{}

	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server grpc server over port 9000: %v", err)
	}
}
