package order

import (
	"log"

	"github.com/gocql/gocql"
)

type Order struct {
	OrderId         gocql.UUID
	ProductIDs      []gocql.UUID
	Quantities      []int32
	Total           float64
	OrderDate       string
	Email           string
	ShippingAddress string
	Status          string
}

type OrderDAL struct {
	Session *gocql.Session
}

func NewOrderDAL(session *gocql.Session) *OrderDAL {
	return &OrderDAL{Session: session}
}

// CreateOrder inserts a new order into the order table
func (dal *OrderDAL) CreateOrder(order Order) error {
	err := dal.Session.Query(`INSERT INTO "order" (OrderId, ProductIDs, Quantities, Total, OrderDate, Email, ShippingAddress, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		order.OrderId, order.ProductIDs, order.Quantities, order.Total, order.OrderDate, order.Email, order.ShippingAddress, order.Status).Exec()
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return err
	}
	return nil
}

// GetAllOrders retrieves all orders from the order table
func (dal *OrderDAL) GetAllOrders() ([]Order, error) {
	var orders []Order
	iter := dal.Session.Query(`SELECT OrderId, ProductIDs, Quantities, Total, OrderDate, Email, ShippingAddress, Status FROM "order"`).Iter()
	defer iter.Close()

	var order Order
	for iter.Scan(&order.OrderId, &order.ProductIDs, &order.Quantities, &order.Total, &order.OrderDate, &order.Email, &order.ShippingAddress, &order.Status) {
		orders = append(orders, order)
	}
	if err := iter.Close(); err != nil {
		log.Printf("Failed to get all orders: %v", err)
		return nil, err
	}

	return orders, nil
}
