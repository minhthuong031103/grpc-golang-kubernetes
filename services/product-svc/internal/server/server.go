package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"productsvc/internal/dal"
	pb "productsvc/internal/generated/product/v1"

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

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productID := gocql.TimeUUID()
	product := dal.Product{
		ProductID:   productID,
		ProductName: req.ProductName,
		Price:       req.Price,
		Description: req.Description,
		Quantity:    req.Quantity,
		Sold:        req.Sold,
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
		Description: product.Description,
		Quantity:    product.Quantity,
		Sold:        product.Sold,
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
			Description: product.Description,
			Quantity:    product.Quantity,
			Sold:        product.Sold,
			ImageUrl:    product.ImageURL,
		})
	}
	return &pb.GetAllProductsResponse{Products: productResponses}, nil
}

func (s *Server) UpdateProductQuantityAndSold(ctx context.Context, req *pb.UpdateProductQuantityAndSoldRequest) (*pb.UpdateProductQuantityAndSoldResponse, error) {
	productID, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, err
	}
	err = s.ProductDAL.UpdateProductQuantityAndSold(productID, req.Quantity, req.Sold)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := s.ProductDAL.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductQuantityAndSoldResponse{
		ProductId: updatedProduct.ProductID.String(),
		Quantity:  updatedProduct.Quantity,
		Sold:      updatedProduct.Sold,
	}, nil
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
	pb.RegisterProductServiceServer(grpcServer, &Server{ProductDAL: productDAL})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
