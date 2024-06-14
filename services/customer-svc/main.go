package main

import (
	"log"

	"customersvc/config"
	dataaccesslayer "customersvc/internal/dal"
	myjwt "customersvc/internal/jwt"
	"customersvc/internal/server"

	"github.com/gocql/gocql"
)

func main() {
	log.Println("Starting CUSTOMER service...")
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

	customerDAL := dataaccesslayer.NewCustomerDAL(session)

	jwtTool := myjwt.NewJWT(config.SecretKey)
	server.StartGRPCServer(config.Server.Port, jwtTool, customerDAL)
}
