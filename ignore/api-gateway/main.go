package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

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
	serviceCache map[string][]*api.ServiceEntry
	cacheMutex   sync.RWMutex
}

func NewService(serverName string) *Service {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://consul-server:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	s := &Service{
		serverName:   serverName,
		consulClient: consulClient,
		serviceCache: make(map[string][]*api.ServiceEntry),
	}

	// Start a goroutine to update the service cache
	go s.updateServiceCache()

	return s
}

func (s *Service) updateServiceCache() {
	for {
		services, _, err := s.consulClient.Catalog().Services(nil)
		if err != nil {
			log.Println("Error fetching services from Consul:", err)
			time.Sleep(time.Second * 5)
			continue
		}

		for serviceName := range services {
			serviceEntries, _, err := s.consulClient.Health().Service(serviceName, "", true, nil)
			if err != nil {
				log.Println("Error fetching service entries from Consul:", err)
				continue
			}

			s.cacheMutex.Lock()
			s.serviceCache[serviceName] = serviceEntries
			s.cacheMutex.Unlock()
		}

		// Use a blocking query to wait for changes
		time.Sleep(time.Minute)
	}
}

func (s *Service) getServiceInstances(serviceName string) []*api.ServiceEntry {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()
	return s.serviceCache[serviceName]
}

func (s *Service) Start() {
	router := gin.Default()

	router.GET("/api/:service/*path", func(c *gin.Context) {
		atomic.AddUint32(&s.requestCount, 1)
		serviceName := c.Param("service")
		serviceInstances := s.getServiceInstances(serviceName)
		if len(serviceInstances) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "service not found or not healthy"})
			return
		}

		targetInstance := serviceInstances[0]
		targetURL := "http://" + targetInstance.Service.Address + ":" + strconv.Itoa(targetInstance.Service.Port)
		proxyURL, err := url.Parse(targetURL)
		fmt.Println("proxyURL", proxyURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(proxyURL)
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/api/"+serviceName)
		fmt.Println("c.Request.URL.Path", c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server_name":   s.serverName,
			"request_count": atomic.LoadUint32(&s.requestCount),
		})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if port == "" {
		port = "8080"
	}
	s := NewService("api-gateway")
	s.Start()
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"sync/atomic"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	address = os.Getenv("POD_IP")
// 	port    = os.Getenv("SERVICE_PORT")
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
// 	router := gin.Default()

// 	router.GET("/api/service-a", func(c *gin.Context) {
// 		atomic.AddUint32(&s.requestCount, 1)
// 		proxyURL, err := url.Parse("http://service-a:8080")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		proxy := httputil.NewSingleHostReverseProxy(proxyURL)
// 		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/api")
// 		proxy.ServeHTTP(c.Writer, c.Request)
// 	})

// 	router.GET("/api/metrics", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"server_name":   s.serverName,
// 			"request_count": atomic.LoadUint32(&s.requestCount),
// 		})
// 	})

// 	if err := router.Run(":" + port); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	if port == "" {
// 		port = "8080"
// 	}
// 	s := NewService("api-gateway")
// 	s.Start()
// }
