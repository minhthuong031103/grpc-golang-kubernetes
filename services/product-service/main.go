package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"productservice/config"
	pb "productservice/generated/product"

	"github.com/gocql/gocql"
	"google.golang.org/grpc"
)

const defaultImageURL = "https://cdn3d.iconscout.com/3d/premium/thumb/product-5806313-4863042.png?f=webp"

type server struct {
	pb.UnimplementedProductServiceServer
	session *gocql.Session
}

func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productID := gocql.TimeUUID()
	err := s.session.Query(`INSERT INTO product (product_id, product_name, price, quantity, image_url) VALUES (?, ?, ?, ?, ?)`,
		productID, req.ProductName, req.Price, req.Quantity, defaultImageURL).Exec()
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{ProductId: productID.String()}, nil
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var productID, productName, imageUrl string
	var price float64
	var quantity int32
	err := s.session.Query(`SELECT product_id, product_name, price, quantity, image_url FROM product WHERE product_id = ?`,
		req.ProductId).Scan(&productID, &productName, &price, &quantity, &imageUrl)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{
		ProductId:   productID,
		ProductName: productName,
		Price:       price,
		Quantity:    quantity,
		ImageUrl:    imageUrl,
	}, nil
}

func (s *server) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	iter := s.session.Query(`SELECT product_id, product_name, price, quantity, image_url FROM product`).Iter()
	defer iter.Close()

	var products []*pb.Product
	var productID, productName, imageUrl string
	var price float64
	var quantity int32

	for iter.Scan(&productID, &productName, &price, &quantity, &imageUrl) {
		products = append(products, &pb.Product{
			ProductId:   productID,
			ProductName: productName,
			Price:       price,
			Quantity:    quantity,
			ImageUrl:    imageUrl,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &pb.GetAllProductsResponse{Products: products}, nil
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cluster := gocql.NewCluster(config.Cassandra.Hosts...)
	cluster.Keyspace = config.Cassandra.Keyspace
	cluster.Consistency = gocql.ParseConsistency(config.Cassandra.Consistency)

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("unable to connect to Cassandra: %v", err)
	}
	defer session.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{session: session})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
