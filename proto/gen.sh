protoc calculator.proto --go_out=../server --go-grpc_out=../server
protoc predict-liveness.proto --go_out=../internal/app/grpc/client --go-grpc_out=../internal/app/grpc/client
