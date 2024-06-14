package server

import (
	"context"
	pb "customersvc/internal/generated/customer"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Token, error) {
	// Check if the user exists
	customer, err := s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	// Check if the password is correct
	if customer.Password != req.Password {
		return nil, nil
	}

	// Generate a JWT token
	token, err := s.JWTTool.GenerateToken(customer.Email)
	if err != nil {
		return nil, err
	}

	err = s.CustomerDAL.UpdateCustomerToken(customer.CustomerId.String(), token)
	if err != nil {
		return nil, err
	}

	return &pb.Token{
		Token: token,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.Authorized, error) {

	return &pb.Authorized{}, nil
}

func (s *Server) Authorized(ctx context.Context, req *pb.Token) (*pb.Authorized, error) {

	return &pb.Authorized{}, nil
}
