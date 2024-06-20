package server

import (
	"context"
	dataaccesslayer "customersvc/internal/dal"
	pb "customersvc/internal/generated/customer"
	"customersvc/internal/helper"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Token, error) {
	log.Println("Login request received for ", req.Email)

	// Check if the user exists
	customer, err := s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err != nil {
		log.Println("Failed to retrieve customer:", err)
		return nil, status.Errorf(http.StatusNotFound, "Failed to retrieve user: %v", err)
	}

	if !helper.CheckPasswordsMatch(customer.Password, req.Password) {
		log.Println("Invalid credentials for", req.Email)
		return nil, status.Errorf(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate a JWT token
	token, err := s.JWTTool.GenerateToken(customer.CustomerId.String())
	if err != nil {
		log.Println("Failed to generate token:", err)
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	err = s.CustomerDAL.UpdateCustomerToken(customer.CustomerId.String(), token)
	if err != nil {
		log.Println("Failed to update customer token:", err)
		return nil, status.Errorf(codes.Internal, "Failed to update token: %v", err)
	}

	log.Println("Login successful for ", req.Email)
	// Return a successful response with the token
	return &pb.Token{
		Token: token,
	}, nil
}

func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.Authorized, error) {
	{ // validateion
		// validate email
		if !helper.ValidateEmail(req.Email) {
			return nil, status.Error(http.StatusBadRequest, "Invalid email")
		}

		// validate password
		if !helper.ValidatePassword(req.Password) {
			return nil, status.Error(http.StatusBadRequest, "Invalid password")
		}

		// validate name
		if !helper.ValidateNotEmptyStr(req.Name) {
			return nil, status.Error(http.StatusBadRequest, "Invalid name")
		}
	}

	// Check if the user already exists
	customer, err := s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err == nil && customer != nil {
		return nil, status.Error(http.StatusConflict, "User already exists")
	}

	// Create a new user
	err = s.CustomerDAL.CreateCustomer(dataaccesslayer.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	// Retrieve the newly created customer for further processing
	customer, err = s.CustomerDAL.GetCustomerByEmail(req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve user after creation")
	}

	// Generate a JWT token
	token, err := s.JWTTool.GenerateToken(customer.CustomerId.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate token")
	}

	err = s.CustomerDAL.UpdateCustomerToken(customer.CustomerId.String(), token)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update user token")
	}

	// Return a successful response with the token
	return &pb.Authorized{
		Authorized: true,
		CustomerId: customer.CustomerId.String(),
		Name:       customer.Name,
		Email:      customer.Email,
		Token:      token,
		Role:       customer.Role,
	}, nil
}

func (s *Server) Authorize(ctx context.Context, req *pb.Token) (*pb.Authorized, error) {
	log.Println("Authorization request received")

	// Validate the JWT token
	claims, err := s.JWTTool.ValidateToken(req.Token)
	if err != nil {
		log.Println("Token validation failed:", err)
		return nil, status.Errorf(http.StatusUnauthorized, "Invalid token: %v", err)
	}

	// Retrieve the customer associated with the token
	customer, err := s.CustomerDAL.GetCustomerByToken(req.Token)
	if err != nil {
		log.Println("Failed to retrieve customer by token:", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve customer: %v", err)
	}

	// Check if the customer ID from the token matches the retrieved customer
	if customer.CustomerId.String() != claims["id"].(string) { // Assuming 'id' is the correct claim, and it is a string
		log.Println("Token customer ID does not match customer record")
		return nil, status.Errorf(http.StatusUnauthorized, "Token does not match the expected user")
	}

	log.Println("Authorization successful for ", customer.Email)
	// Return a successful response
	return &pb.Authorized{
		Authorized: true,
		CustomerId: customer.CustomerId.String(),
		Name:       customer.Name,
		Email:      customer.Email,
		Token:      req.Token,
		Role:       customer.Role,
	}, nil
}

func (s *Server) SetRole(ctx context.Context, req *pb.SetRoleRequest) (*pb.SetRoleResponse, error) {
	log.Println("Role update request received")

	// Get the requestor's role
	requestor, err := s.CustomerDAL.GetCustomerByToken(req.Token)
	if err != nil {
		log.Println("Failed to retrieve requestor:", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve requestor: %v", err)
	}
	if requestor.Role != "admin" {
		log.Println("Unauthorized requestor")
		return nil, status.Error(http.StatusUnauthorized, "Unauthorized")
	}

	// Update the role
	err = s.CustomerDAL.UpdateRole(req.AccountId, req.Role)
	if err != nil {
		log.Println("Failed to update role:", err)
		return nil, status.Errorf(codes.Internal, "Failed to update role: %v", err)
	}

	// Return a successful response
	return &pb.SetRoleResponse{
		Success:   true,
		AccountId: req.AccountId,
		Role:      req.Role,
		Message:   "Role updated successfully",
	}, nil
}
