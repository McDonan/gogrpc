package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {

}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func (s calculatorServer) Average(stream Calculator_AverageServer) error {
	sum, count := 0.0, 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}

	res := AverageResponse{Result: sum / count}
	return stream.SendAndClose(&res)
}

func (s calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{Result: sum}
		if err := stream.Send(&res); err != nil {
			return err
		}
	}

	return nil
}

func (s calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fibonacci(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fibonacci(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
