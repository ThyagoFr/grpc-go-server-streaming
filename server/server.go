package main

import (
	"fmt"
	"log"
	"net"

	primenumber "github.com/thyagofr/grpc-go-server-streaming/prime"
	"github.com/thyagofr/grpc-go-server-streaming/server/calc"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) PrimeNumberDecomposition(request *primenumber.PrimeNumberStreamingRequest, stream primenumber.PrimeNumberStreamingService_PrimeNumberDecompositionServer) error {
	number := request.Request.GetNumber()
	channel := make(chan int64)
	go calc.CalculatePrime(number, channel)
	for v := range channel {
		stream.Send(&primenumber.PrimeNumberStreamingResponse{
			Response: v,
		})
	}
	return nil
}

func main() {

	fmt.Println("Running server...")
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	serv := grpc.NewServer()
	primenumber.RegisterPrimeNumberStreamingServiceServer(serv, &server{})
	if err := serv.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Println("Running...")
}
