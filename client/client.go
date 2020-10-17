package main

import (
	"context"
	"log"

	"github.com/Muchogoc/hellogrpc/api"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	message := api.PingMessage{
		Greeting: "Hello from client",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response)
}
