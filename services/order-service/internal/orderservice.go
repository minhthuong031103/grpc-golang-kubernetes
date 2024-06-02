package internal

import (
	"context"
	"fmt"
	"log"
	orders "order-service/protogen/golang/orders"
	"order-service/protogen/golang/user"
)

// OrderService should implement the OrdersServer interface generated from grpc.
//
// UnimplementedOrdersServer must be embedded to have forwarded compatible implementations.
type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
	userClient user.UserServiceClient
}

// NewOrderService creates a new OrderService
func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

// SetUserClient sets the user client for the order service
func (s *OrderService) SetUserClient(userClient user.UserServiceClient) {
	s.userClient = userClient
}

// AddOrder implements the AddOrder method of the grpc OrdersServer interface to add a new order
func (s *OrderService) AddOrder(ctx context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	// Example: Fetch user details from UserService
	userResp, err := s.userClient.GetUser(ctx, &user.GetUserRequest{UserId: req.Order.CustomerId})
	if err != nil {
		return nil, err
	}

	// Use user details for further processing
	log.Printf("Fetched user details: %v", userResp.User)

	// Add order logic
	return &orders.Empty{}, s.db.AddOrder(req.GetOrder())
}

// GetOrder implements the GetOrder method of the grpc OrdersServer interface to fetch an order for a given orderID
func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Received get order request")
	order := o.db.GetOrderByID(req.GetOrderId())
	if order == nil {
		return nil, fmt.Errorf("order not found for orderID: %d", req.GetOrderId())
	}
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

// UpdateOrder implements the UpdateOrder method of the grpc OrdersServer interface to update an order
func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an update order request")
	o.db.UpdateOrder(req.GetOrder())
	return &orders.Empty{}, nil
}

// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")
	o.db.RemoveOrder(req.GetOrderId())
	return &orders.Empty{}, nil
}

// GetAllOrders implements the GetAllOrders method of the grpc OrdersServer interface to fetch all orders
func (o *OrderService) ListOrders(_ context.Context, _ *orders.Empty) (*orders.PayloadWithMultipleOrders, error) {
	log.Printf("Received get all orders request")
	return &orders.PayloadWithMultipleOrders{Orders: o.db.GetAllOrders()}, nil
}

// ./internal/db.go
