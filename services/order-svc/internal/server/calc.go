package server

import (
	"context"
	pb "ordersvc/internal/generated/order"
	productpb "ordersvc/internal/generated/product"
)

func calculateTotalPrice(s *Server, items []*pb.OrderItem) (float64, error) {
	var total float64 = 0
	for _, item := range items {
		resp, err := s.ProductClient.GetProduct(context.Background(), &productpb.GetProductRequest{ProductId: item.ProductId})
		if err != nil {
			return 0, err
		}

		item.Price = resp.Price
		total += float64(item.Quantity) * item.Price
	}
	return total, nil
}
