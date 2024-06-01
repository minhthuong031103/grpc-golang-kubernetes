package main

import (
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ttl = time.Second * 8
)

var (
	address = os.Getenv("POD_IP")
)

type Service struct {
	requestCount uint32
	serverName   string
}

func NewService(serverName string) *Service {

	return &Service{
		serverName: serverName,
	}
}

func (s *Service) Start() {
	address := os.Getenv("POD_IP")
	port := os.Getenv("SERVICE_PORT")
	router := gin.Default()
	router.GET("/service-a", func(c *gin.Context) {
		atomic.AddUint32(&s.requestCount, 1)
		c.JSON(http.StatusOK, gin.H{
			"message":       "Hello from service-a" + address + ":" + port,
			"server_name":   s.serverName,
			"request_count": atomic.LoadUint32(&s.requestCount),
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := NewService("service-a")
	s.Start()
}
