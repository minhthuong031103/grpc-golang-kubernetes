package server

import (
	"fmt"
	"log"
	"net"

	"fileuploadsvc/config"
	pb "fileuploadsvc/internal/generated/fileupload"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFileUploadServiceServer
	s3Client *s3.S3
	bucket   string
}

func StartGRPCServer(port int, s3cfg config.S3StorageConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:                        aws.String(s3cfg.Region),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Credentials:                   credentials.NewStaticCredentials(s3cfg.AccessKey, s3cfg.SecretKey, ""),
	}))
	s3Client := s3.New(sess)

	pb.RegisterFileUploadServiceServer(grpcServer, &Server{
		s3Client: s3Client,
		bucket:   s3cfg.Bucket,
	})

	log.Printf("fileupload-server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
