package grpcclient

import (
	"google.golang.org/grpc"
)

// GrpcClient defines the common operations for any gRPC client
type GrpcClient interface {
	// Connect establishes a connection to the gRPC service
	Connect() error
	// Disconnect closes the connection to the gRPC service
	Disconnect() error
	// GetConnection returns the underlying gRPC connection
	GetConnection() *grpc.ClientConn
}
