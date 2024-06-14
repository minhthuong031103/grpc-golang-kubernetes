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

	// fmt.Println("Logging in from service " + address + "...")
	// getRes, err := client.Login(ctx, &pb.LoginRequest{
	// 	Email:    "thuong@gm.com",
	// 	Password: "123456",
	// })
	// if err != nil {
	// 	fmt.Printf("Could not login: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Token: %v\n", getRes.Token)

	// fmt.Println("Trying to register from service " + address + "...")
	// registerRes, err := client.Register(ctx, &pb.RegisterRequest{
	// 	Name:     "Thuong",
	// 	Email:    "a@gm.com",
	// 	Password: "123456",
	// })
	// if err != nil {
	// 	fmt.Printf("Could not register: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Authorized: %v\n", registerRes.Authorized)

	fmt.Println("Trying to authenticate using token from service " + address + "...")
	authRes, err := client.Authorize(ctx, &pb.Token{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCJ9.ZQ8DPJjhN1t7oc7QPcYvqwK-Ze3o3xLefMVQCQPsN-g",
	})
	if err != nil {
		fmt.Printf("Could not authenticate: %v\n", err)
		return
	}
	fmt.Printf("Authorized: %v\n", authRes.Authorized)
	fmt.Printf("Email: %v\n", authRes.Email)
	fmt.Printf("Name: %v\n", authRes.Name)
	fmt.Printf("Token: %v\n", authRes.Token)
	fmt.Printf("CustomerId: %v\n", authRes.CustomerId)

}
