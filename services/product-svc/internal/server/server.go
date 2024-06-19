package server

import (
	"fmt"
	"log"
	"net"

	"productsvc/internal/dal"
	pb "productsvc/internal/generated/product"

	"google.golang.org/grpc"
)

const defaultImageURL = "https://cdn3d.iconscout.com/3d/premium/thumb/product-5806313-4863042.png?f=webp"

type Server struct {
	pb.UnimplementedProductServiceServer
	ProductDAL *dal.ProductDAL
}

func StartGRPCServer(port int, productDAL *dal.ProductDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)

	pb.RegisterProductServiceServer(grpcServer, &Server{ProductDAL: productDAL})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
