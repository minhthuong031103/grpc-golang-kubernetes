// package server

// import (
// 	"context"
// 	pb "ordersvc/internal/generated/order"
// 	productpb "ordersvc/internal/generated/product"
// )

// func (s *Server) ListOrders(ctx context.Context, req *pb.Empty) (*pb.ListOrderDetailsResponse, error) {
// 	orders, err := s.OrderDAL.GetAllOrders()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var pbOrders []*pb.OrderDetails
// 	for _, order := range orders {
// 		pbOrder := &pb.OrderDetails{
// 			OrderId:         order.OrderId.String(),
// 			Products:        []*productpb.Product{},
// 			Quantities:      order.Quantities,
// 			Total:           order.Total,
// 			OrderDate:       order.OrderDate,
// 			Email:           order.Email,
// 			ShippingAddress: order.ShippingAddress,
// 			Status:          order.Status,
// 		}

// 		pbOrders = append(pbOrders, pbOrder)
// 	}
// 	return &pb.ListOrderDetailsResponse{
// 		Orders: pbOrders,
// 	}, nil
// }

package server

import (
	"context"
	pb "ordersvc/internal/generated/order"
	productpb "ordersvc/internal/generated/product"
	"time"
)

func (s *Server) ListOrders(ctx context.Context, req *pb.Empty) (*pb.ListOrderDetailsResponse, error) {
	orders, err := s.OrderDAL.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.OrderDetails
	for _, order := range orders {
		pbOrder := &pb.OrderDetails{
			OrderId:         order.OrderId.String(),
			Products:        []*productpb.Product{},
			Quantities:      order.Quantities,
			Total:           order.Total,
			OrderDate:       order.OrderDate,
			Email:           order.Email,
			ShippingAddress: order.ShippingAddress,
			Status:          order.Status,
		}

		productClient := productpb.NewProductServiceClient(s.ProductClient.GetConnection())
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		for _, productID := range order.ProductIDs {
			id := productID.String()
			getRes, err := productClient.GetProduct(ctx, &productpb.GetProductRequest{
				ProductId: id,
			})
			if err != nil {
				return nil, err
			}
			pbOrder.Products = append(pbOrder.Products, &productpb.Product{
				ProductId:   getRes.ProductId,
				ProductName: getRes.ProductName,
				Price:       getRes.Price,
				Description: getRes.Description,
				Quantity:    getRes.Quantity,
				Sold:        getRes.Sold,
			})
		}
		pbOrders = append(pbOrders, pbOrder)
	}
	return &pb.ListOrderDetailsResponse{
		Orders: pbOrders,
	}, nil
}
