package server

import (
	"fmt"
	"log"
	"net"

	dal "ordersvc/internal/dal"
	pb "ordersvc/internal/generated/order"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOrdersServer
	OrderDAL *dal.OrderDAL
}

func StartGRPCServer(port int, orderDAL *dal.OrderDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterOrdersServer(grpcServer, &Server{OrderDAL: orderDAL})

	log.Printf("order-server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
