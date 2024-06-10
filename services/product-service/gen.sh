protoc --go_out=. --go-grpc_out=. proto/product.proto
protoc --go_out=./client --go-grpc_out=./client proto/product.proto