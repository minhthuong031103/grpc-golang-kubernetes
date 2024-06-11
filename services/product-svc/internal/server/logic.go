package server

import (
	"context"
	"productsvc/internal/dal"
	pb "productsvc/internal/generated/product/v1"

	"github.com/gocql/gocql"
)

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
