package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

var (
	address = os.Getenv("POD_IP")
	port    = os.Getenv("SERVICE_PORT")
)

type Service struct {
	requestCount uint32
	serverName   string
	consulClient *api.Client
	serviceID    string
}

func NewService(serverName string) *Service {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://consul-server:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		serverName:   serverName,
		consulClient: consulClient,
		serviceID:    serverName + "-" + address + "-" + port,
	}
}

func (s *Service) registerService() {
	registration := &api.AgentServiceRegistration{
		ID:      s.serviceID,
		Name:    s.serverName,
		Address: s.serverName,
		Port:    8080,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + address + ":" + port + "/health",
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	err := s.consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}
	log.Println("Service registered with Consul")
}

func (s *Service) deregisterService() {
	err := s.consulClient.Agent().ServiceDeregister(s.serviceID)
	if err != nil {
		log.Println("Failed to deregister service:", err)
	}
	log.Println("Service deregistered from Consul")
}

func (s *Service) Start() {
	address := os.Getenv("POD_IP")
	port := os.Getenv("SERVICE_PORT")
	router := gin.Default()

	router.GET("/service-a", func(c *gin.Context) {
		atomic.AddUint32(&s.requestCount, 1)
		c.JSON(http.StatusOK, gin.H{
			"message":       "Hello from service-a " + address + ":" + port,
			"server_name":   s.serverName,
			"request_count": atomic.LoadUint32(&s.requestCount),
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	go func() {
		if err := router.Run(":" + port); err != nil {
			log.Fatal(err)
		}
	}()

	s.registerService()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	s.deregisterService()
}

func main() {
	if port == "" {
		port = "8080"
	}
	s := NewService("service-a")
	s.Start()
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"sync/atomic"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// const (
// 	ttl = time.Second * 8
// )

// var (
// 	address = os.Getenv("POD_IP")
// )

// type Service struct {
// 	requestCount uint32
// 	serverName   string
// }

// func NewService(serverName string) *Service {

// 	return &Service{
// 		serverName: serverName,
// 	}
// }

// func (s *Service) Start() {
// 	address := os.Getenv("POD_IP")
// 	port := os.Getenv("SERVICE_PORT")
// 	router := gin.Default()
// 	router.GET("/service-a", func(c *gin.Context) {
// 		atomic.AddUint32(&s.requestCount, 1)
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":       "Hello from service-a" + address + ":" + port,
// 			"server_name":   s.serverName,
// 			"request_count": atomic.LoadUint32(&s.requestCount),
// 		})
// 	})

// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	s := NewService("service-a")
// 	s.Start()
// }
