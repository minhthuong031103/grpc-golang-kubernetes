package main

import (
	"context"
	"fmt"
	"time"

	"productservice/config"
	pb "productservice/generated/product"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		return
	}
	address := fmt.Sprintf("localhost:%d", config.Server.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Did not connect: %v\n", err)
		return
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// // Create a new product
	// fmt.Println("Creating a new product...")
	// createRes, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
	// 	ProductName: "New Product",
	// 	Price:       19.99,
	// 	Description: "A brand new product.",
	// 	Quantity:    10,
	// 	Sold:        0,
	// })
	// if err != nil {
	// 	fmt.Printf("Could not create product: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Product created with ID: %s\n", createRes.GetProductId())

	// // Get the product
	// fmt.Printf("Getting product with ID: %s\n", createRes.GetProductId())
	// getRes, err := client.GetProduct(ctx, &pb.GetProductRequest{
	// 	ProductId: createRes.GetProductId(),
	// })
	// if err != nil {
	// 	fmt.Printf("Could not get product: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Product: %v\n", getRes)

	// Get all products
	fmt.Println("Getting all products...")
	getAllRes, err := client.GetAllProducts(ctx, &pb.GetAllProductsRequest{})
	if err != nil {
		fmt.Printf("Could not get all products: %v\n", err)
		return
	}
	for _, product := range getAllRes.GetProducts() {
		fmt.Printf("Product ID: %s, Name: %s, Price: %.2f, Description: %s, Quantity: %d, Sold: %d, Image URL: %s\n",
			product.ProductId, product.ProductName, product.Price, product.Description, product.Quantity, product.Sold, product.ImageUrl)
	}
}
