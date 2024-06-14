package server

import (
	"context"
	dataaccesslayer "customersvc/internal/dal"
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
	token, err := s.JWTTool.GenerateToken(customer.CustomerId.String())
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
	// Check if the user already exists
	_, err := s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err == nil {
		return nil, nil
	}

	// Create a new user

	err = s.CustomerDAL.CreateCustomer(dataaccesslayer.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	customer, err := s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	// Generate a JWT token
	token, err := s.JWTTool.GenerateToken(customer.CustomerId.String())
	if err != nil {
		return nil, err
	}

	err = s.CustomerDAL.UpdateCustomerToken(customer.CustomerId.String(), token)
	if err != nil {
		return nil, err
	}

	return &pb.Authorized{
		Authorized: true,
		CustomerId: customer.CustomerId.String(),
		Name:       customer.Name,
		Email:      customer.Email,
		Token:      token,
	}, nil
}

func (s *Server) Authorize(ctx context.Context, req *pb.Token) (*pb.Authorized, error) {
	claims, err := s.JWTTool.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}

	customer, err := s.CustomerDAL.GetCustomerByToken(req.Token)
	if err != nil {
		return nil, err
	}

	if customer.CustomerId.String() != claims["id"] {
		return nil, nil
	}

	return &pb.Authorized{
		Authorized: true,
		CustomerId: customer.CustomerId.String(),
		Name:       customer.Name,
		Email:      customer.Email,
		Token:      req.Token,
	}, nil
}
