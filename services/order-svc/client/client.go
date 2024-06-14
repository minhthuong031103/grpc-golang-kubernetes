package main

import (
	"context"
	"fmt"
	"time"

	"ordersvc/config"
	pb "ordersvc/internal/generated/order"

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
	client := pb.NewOrdersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Getting all orders from service " + address + "...")
	getAllRes, err := client.ListOrders(ctx, &pb.Empty{})
	if err != nil {
		fmt.Printf("Could not get all orders: %v\n", err)
		return
	}
	for _, order := range getAllRes.GetOrders() {
		fmt.Printf("Order: %v\n", order)
	}

}
