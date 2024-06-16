package server

import (
	"bytes"
	"context"
	"fmt"

	pb "fileuploadsvc/internal/generated/fileupload"
	"fileuploadsvc/internal/helper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *Server) UploadFile(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploaderWithClient(s.s3Client)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fmt.Sprintf("f%v-%v", helper.GenerateID(), req.Filename)),
		Body:   bytes.NewReader(req.Data),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	// Return the URL of the new object
	return &pb.UploadResponse{
		Url: result.Location,
	}, nil
}
