package server

import (
	"bytes"
	"context"
	fileuploadpb "gateway/internal/generated/fileupload"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func uploadFileHandler(fileuploadConn *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			c.Abort()
			return
		}

		// abort if file size > 10MB
		if file.Size > 10<<20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB"})
			c.Abort()
			return
		}

		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
			c.Abort()
			return
		}
		defer openedFile.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(openedFile)

		client := fileuploadpb.NewFileUploadServiceClient(fileuploadConn)
		res, err := client.UploadFile(context.Background(), &fileuploadpb.UploadRequest{
			Filename: file.Filename,
			Data:     buf.Bytes(),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file: " + err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"url": res.Url})
	}
}
