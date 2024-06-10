package main

import (
	"context"
	"log"
	"time"

	pb "productservice/generated/product"

	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// // Create a new product
	// log.Println("Creating a new product...")
	// createRes, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
	// 	ProductName: "New Product",
	// 	Price:       19.99,
	// 	Quantity:    10,
	// })
	// if err != nil {
	// 	log.Fatalf("Could not create product: %v", err)
	// }
	// log.Printf("Product created with ID: %s", createRes.GetProductId())

	// // Get the product
	// log.Printf("Getting product with ID: %s", createRes.GetProductId())
	// getRes, err := client.GetProduct(ctx, &pb.GetProductRequest{
	// 	ProductId: createRes.GetProductId(),
	// })
	// if err != nil {
	// 	log.Fatalf("Could not get product: %v", err)
	// }
	// log.Printf("Product: %v", getRes)

	// Get all products
	log.Println("Getting all products...")
	getAllRes, err := client.GetAllProducts(ctx, &pb.GetAllProductsRequest{})
	if err != nil {
		log.Fatalf("Could not get all products: %v", err)
	}
	for _, product := range getAllRes.GetProducts() {
		log.Printf("Product ID: %s, Name: %s, Price: %.2f, Quantity: %d, Image URL: %s",
			product.ProductId, product.ProductName, product.Price, product.Quantity, product.ImageUrl)
	}
}
