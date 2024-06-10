package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "productservice/generated/product"
	"productservice/internal/dal"

	"github.com/gocql/gocql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const defaultImageURL = "https://cdn3d.iconscout.com/3d/premium/thumb/product-5806313-4863042.png?f=webp"

type Server struct {
	pb.UnimplementedProductServiceServer
	ProductDAL *dal.ProductDAL
}

func NewServer(productDAL *dal.ProductDAL) *Server {
	return &Server{ProductDAL: productDAL}
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productID := gocql.TimeUUID()
	product := dal.Product{
		ProductID:   productID,
		ProductName: req.ProductName,
		Price:       req.Price,
		Quantity:    req.Quantity,
		ImageURL:    defaultImageURL,
	}
	err := s.ProductDAL.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{ProductId: productID.String()}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	productID, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, err
	}
	product, err := s.ProductDAL.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{
		ProductId:   product.ProductID.String(),
		ProductName: product.ProductName,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageUrl:    product.ImageURL,
	}, nil
}

func (s *Server) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	products, err := s.ProductDAL.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var productResponses []*pb.Product
	for _, product := range products {
		productResponses = append(productResponses, &pb.Product{
			ProductId:   product.ProductID.String(),
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    product.Quantity,
			ImageUrl:    product.ImageURL,
		})
	}
	return &pb.GetAllProductsResponse{Products: productResponses}, nil
}

// UnaryInterceptor is a server interceptor for logging requests.
func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	method := info.FullMethod
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Received request: method=%s, metadata=%v", method, md)

	// Calls the handler
	resp, err := handler(ctx, req)

	duration := time.Since(start)
	statusCode := status.Code(err)
	log.Printf("Completed request: method=%s, duration=%s, status=%s", method, duration, statusCode)
	return resp, err
}

func StartGRPCServer(port int, productDAL *dal.ProductDAL) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterProductServiceServer(grpcServer, NewServer(productDAL))
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
