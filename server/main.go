package main

import (
	"fmt"
	"log"
	"net"

	"server/services"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	fmt.Println("gRPC server listening from port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
