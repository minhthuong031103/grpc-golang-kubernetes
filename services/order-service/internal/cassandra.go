package internal

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("77.37.47.87")
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Keyspace = "system" // Connect to system keyspace to create a new keyspace
	cluster.Consistency = gocql.LocalOne

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Create your own keyspace
	err = Session.Query("CREATE KEYSPACE IF NOT EXISTS order_data WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 2}").Exec()
	if err != nil {
		panic(err)
	}

	// Use your own keyspace
	cluster.Keyspace = "order_data"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Create the orders table
	err = Session.Query(`
		CREATE TABLE IF NOT EXISTS orders (
			order_id UUID PRIMARY KEY,
			total DECIMAL,
		)
	`).Exec()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("Cassandra initialized successfully")
}
