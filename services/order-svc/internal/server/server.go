package server

import (
	"fmt"
	"log"
	"net"

	productclient "ordersvc/internal/client/product"
	dal "ordersvc/internal/dal"
	pb "ordersvc/internal/generated/order"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOrdersServer
	OrderDAL      *dal.OrderDAL
	ProductClient *productclient.ProductClient
}

func StartGRPCServer(port int, productClient *productclient.ProductClient, orderDAL *dal.OrderDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterOrdersServer(grpcServer, &Server{
		OrderDAL:      orderDAL,
		ProductClient: productClient})

	log.Printf("order-server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
