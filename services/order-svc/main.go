package main

import (
	"log"

	"ordersvc/config"
	orderdal "ordersvc/internal/dal"
	"ordersvc/internal/server"

	"github.com/gocql/gocql"
)

func main() {
	log.Println("Starting product service...")
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

	server.StartGRPCServer(config.Server.Port, orderDAL)
}
