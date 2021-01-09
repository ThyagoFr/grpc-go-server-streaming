package main

import (
	"context"
	"fmt"
	"io"
	"log"

	primenumber "github.com/thyagofr/grpc-go-server-streaming/prime"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Running grpc client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error creating grpc client : %v", err)
	}
	defer cc.Close()
	client := primenumber.NewPrimeNumberStreamingServiceClient(cc)
	req := &primenumber.PrimeNumberStreamingRequest{
		Request: &primenumber.PrimeNumberStreaming{
			Number: 120,
		},
	}
	stream, err := client.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.GetResponse())
	}

}
