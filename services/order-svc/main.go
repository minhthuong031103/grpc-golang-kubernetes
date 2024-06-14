package main

import (
	"fmt"
	"log"

	"ordersvc/config"
	productclient "ordersvc/internal/client/product"
	orderdal "ordersvc/internal/dal"
	"ordersvc/internal/server"

	"github.com/gocql/gocql"
)

func main() {
	log.Println("Starting ORDER service...")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cluster := gocql.NewCluster(config.Cassandra.Hosts...)
	cluster.Keyspace = config.Cassandra.Keyspace
	cluster.Consistency = gocql.ParseConsistency(config.Cassandra.Consistency)

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("unable to connect to Cassandra: %v", err)
	}
	defer session.Close()

	orderDAL := orderdal.NewOrderDAL(session)

	productAddress := fmt.Sprintf("%v:%v", config.ProductSvc.Host, config.ProductSvc.Port)
	productClientConn := productclient.NewProductClient(productAddress)
	if err := productClientConn.Connect(); err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	log.Println("Connected to product service at", productAddress)
	defer productClientConn.Disconnect()

	server.StartGRPCServer(config.Server.Port, productClientConn, orderDAL)
}
