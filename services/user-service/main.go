package main

import (
	"log"
	"net"
	"user-service/internal"
	users "user-service/protogen/golang/user"

	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50052"

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server instance
	server := grpc.NewServer()

	// Create a user service instance with a reference to the db
	db := internal.NewDB()
	userService := internal.NewUserService(db)

	// Register the user service with the gRPC server
	users.RegisterUserServiceServer(server, userService)

	// Start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
