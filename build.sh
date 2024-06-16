#!/bin/bash

# build the services

echo "Working directory: $(pwd)"
echo "Building services..."

echo "Building PRODUCT service..."
cd services/product-svc
go build -o ../../build/product main.go
cd ../..

echo "Building CUSTOMER service..."
cd services/customer-svc
go build -o ../../build/customer main.go
cd ../..

echo "Building ORDER service..."
cd services/order-svc
go build -o ../../build/order main.go
cd ../..

echo "Building GATEWAY service..."
cd grpc-gateway
go build -o ../build/gateway main.go
cd ..

echo "Build complete."