package main

import (
	"log"
	"net"
	"order-service/internal"
	orders "order-service/protogen/golang/orders"
	users "order-service/protogen/golang/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const addr = "0.0.0.0:50051"
	const userServiceAddr = "0.0.0.0:50052"

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server instance
	server := grpc.NewServer()

	// Create a order service instance with a reference to the db
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	// Register the order service with the gRPC server
	orders.RegisterOrdersServer(server, &orderService)

	// Connect to the user service
	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer conn.Close()
	userClient := users.NewUserServiceClient(conn)
	orderService.SetUserClient(userClient)

	// Start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
