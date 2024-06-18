package server

import (
	"fmt"
	"log"
	"net"

	grpcclientconn "ordersvc/internal/connection"
	dal "ordersvc/internal/dal"
	pb "ordersvc/internal/generated/order"
	productpb "ordersvc/internal/generated/product"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	OrderDAL      *dal.OrderDAL
	ProductClient productpb.ProductServiceClient
}

func StartGRPCServer(port int, productClientConn *grpcclientconn.GRPCClient, orderDAL *dal.OrderDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)

	productClient := productpb.NewProductServiceClient(productClientConn.GetConnection())
	pb.RegisterOrderServiceServer(grpcServer, &Server{
		OrderDAL:      orderDAL,
		ProductClient: productClient})

	log.Printf("order-server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
