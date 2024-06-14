package main

import (
	"context"
	"fmt"
	"time"

	"customersvc/config"
	pb "customersvc/internal/generated/customer"

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
	client := pb.NewCustomerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Logging in from service " + address + "...")
	getRes, err := client.Login(ctx, &pb.LoginRequest{
		Email:    "thuong@gm.com",
		Password: "123456",
	})
	if err != nil {
		fmt.Printf("Could not login: %v\n", err)
		return
	}
	fmt.Printf("Token: %v\n", getRes.Token)

}