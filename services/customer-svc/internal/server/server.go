package server

import (
	"fmt"
	"log"
	"net"

	dal "customersvc/internal/dal"
	pb "customersvc/internal/generated/customer"

	myjwt "customersvc/internal/jwt"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCustomerServiceServer
	CustomerDAL *dal.CustomerDAL
	JWTTool     *myjwt.JWT
}

func StartGRPCServer(port int, JWTTool *myjwt.JWT, customerDAL *dal.CustomerDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterCustomerServiceServer(grpcServer, &Server{
		CustomerDAL: customerDAL,
		JWTTool:     JWTTool,
	})

	log.Printf("customer-server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
