package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"server/services"

	"google.golang.org/grpc"
)

func main() {
	var s *grpc.Server
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()
	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}

		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)

	fmt.Print("gRPC server listening from port 50051")
	if *tls {
		fmt.Println(" with TLS")
	}
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
