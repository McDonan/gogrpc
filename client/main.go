package main

import (
	"client/services"
	"flag"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var cc *grpc.ClientConn
	var creds credentials.TransportCredentials
	var err error

	host := flag.String("host", "localhost:50051", "gRPC Server host")
	tls := flag.Bool("tls", false, "use secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}

	} else {
		creds = insecure.NewCredentials()
	}
	defer cc.Close()

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	if err := calculatorService.Hello("Bond"); err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}
	//if err := calculatorService.Fibonacci(6); err != nil {
	//	log.Fatal(err)
	//}
	//if err := calculatorService.Average(1, 2, 3, 4, 5, 6); err != nil {
	//	log.Fatal(err)
	//}
	//if err := calculatorService.Sum(1, 2, 3, 4, 5, 6); err != nil {
	//	log.Fatal(err)
	//}
}
