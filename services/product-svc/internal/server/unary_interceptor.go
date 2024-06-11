package server

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryInterceptor is a server interceptor for logging requests.
func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	method := info.FullMethod
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Received request: method=%s, metadata=%v", method, md)

	// Calls the handler
	resp, err := handler(ctx, req)

	duration := time.Since(start)
	statusCode := status.Code(err)
	log.Printf("Completed request: method=%s, duration=%s, status=%s", method, duration, statusCode)
	return resp, err
}
