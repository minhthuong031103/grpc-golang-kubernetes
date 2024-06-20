package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/status"

	dal "ordersvc/internal/dal"
	pb "ordersvc/internal/generated/order"
	productpb "ordersvc/internal/generated/product"
	"ordersvc/internal/helper"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	// Generate a new UUID for the order
	customerId, err := gocql.ParseUUID(req.CustomerId)
	if err != nil {
		return nil, err
	}

	dalOrder := dal.Order{
		OrderId:    gocql.TimeUUID(),
		CustomerId: customerId,
		OrderDate:  helper.GetTimeNowInGMT7(),
		Status:     "pending",
		TotalPrice: 0,
		Products:   make([]dal.OrderItem, 0),
		CreatedAt:  helper.GetTimeNowInGMT7(),
		UpdatedAt:  helper.GetTimeNowInGMT7(),
		DeletedAt:  "",
	}

	for i, item := range req.Items {
		tmpProduct, err := s.ProductClient.GetProduct(ctx, &productpb.GetProductRequest{ProductId: item.ProductId})
		if err != nil {
			return nil, status.Errorf(http.StatusInternalServerError, "Failed to calc product: %v", err)
		}

		productId, err := gocql.ParseUUID(tmpProduct.ProductId)
		if err != nil {
			return nil, status.Errorf(http.StatusInternalServerError, "Failed to parse product id: %v", err)
		}

		dalOrder.TotalPrice += float64(item.Quantity) * tmpProduct.Price
		req.Items[i].Price = tmpProduct.Price
		req.Items[i].ProductName = tmpProduct.ProductName
		dalOrder.Products = append(dalOrder.Products, dal.OrderItem{
			ProductId:   productId,
			ProductName: tmpProduct.ProductName,
			Quantity:    item.Quantity,
			Price:       tmpProduct.Price,
		})
	}

	err = s.OrderDAL.CreateOrder(dalOrder)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		OrderId:    dalOrder.OrderId.String(),
		CustomerId: dalOrder.CustomerId.String(),
		OrderDate:  dalOrder.OrderDate,
		Status:     dalOrder.Status,
		TotalPrice: dalOrder.TotalPrice,
		Items:      req.Items,
		CreatedAt:  dalOrder.CreatedAt,
		UpdatedAt:  dalOrder.UpdatedAt,
	}, nil
}

func (s *Server) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	var orders []*pb.Order

	allOrder, err := s.OrderDAL.GetAllOrders()
	if err != nil {
		fmt.Print(err)
		return nil, status.Errorf(http.StatusInternalServerError, "Cant get data from db")
	}

	for _, order := range allOrder {
		orders = append(orders, &pb.Order{
			OrderId:    order.OrderId.String(),
			CustomerId: order.CustomerId.String(),
			Items:      make([]*pb.OrderItem, 0),
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			OrderDate:  order.OrderDate,
		})
	}
	return &pb.GetAllOrdersResponse{Orders: orders,
		Total: int32(len(orders)),
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	orderId, err := gocql.ParseUUID(req.OrderId)
	if err != nil {
		return nil, err
	}

	order, err := s.OrderDAL.GetOrder(orderId)
	if err != nil {
		return nil, err
	}

	items := make([]*pb.OrderItem, 0)
	for _, item := range order.Products {
		items = append(items, &pb.OrderItem{
			ProductId:   item.ProductId.String(),
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		})
	}

	return &pb.Order{
		OrderId:    order.OrderId.String(),
		CustomerId: order.CustomerId.String(),
		Items:      items,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		OrderDate:  order.OrderDate,
	}, nil
}

func (s *Server) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	orderId, err := gocql.ParseUUID(req.OrderId)
	if err != nil {
		return nil, status.Errorf(http.StatusInternalServerError, "Failed to parse order id: %v", err)
	}

	order, err := s.OrderDAL.GetOrder(orderId)
	if err != nil {
		return nil, status.Errorf(http.StatusInternalServerError, "Failed to get order: %v", err)
	}

	order.Status = req.Status
	order.UpdatedAt = helper.GetTimeNowInGMT7()

	err = s.OrderDAL.UpdateOrderStatus(orderId, req.Status)
	if err != nil {
		return nil, status.Errorf(http.StatusInternalServerError, "Failed to update order status: %v", err)
	}

	return &pb.UpdateOrderStatusResponse{
		OrderId: order.OrderId.String(),
		Status:  order.Status,
	}, nil
}
