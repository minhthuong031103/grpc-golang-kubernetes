package server

import (
	"context"
	pb "ordersvc/internal/generated/order"
)

func (s *Server) ListOrders(ctx context.Context, req *pb.Empty) (*pb.PayloadWithMultipleOrders, error) {
	orders, err := s.OrderDAL.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrder := &pb.Order{
			OrderId:         order.OrderId.String(),
			ProductIds:      []string{},
			Quantities:      order.Quantities,
			Total:           order.Total,
			OrderDate:       order.OrderDate,
			Email:           order.Email,
			ShippingAddress: order.ShippingAddress,
			Status:          order.Status,
		}
		for _, productID := range order.ProductIDs {
			pbOrder.ProductIds = append(pbOrder.ProductIds, productID.String())
		}
		pbOrders = append(pbOrders, pbOrder)
	}
	return &pb.PayloadWithMultipleOrders{
		Orders: pbOrders,
	}, nil
}
