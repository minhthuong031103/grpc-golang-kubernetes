package order

import (
	"log"
	"ordersvc/internal/helper"

	"github.com/gocql/gocql"
)

type Order struct {
	OrderId    gocql.UUID
	CustomerId gocql.UUID
	OrderDate  string
	Status     string
	TotalPrice float64
	Products   []OrderItem
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}

type OrderItem struct {
	ProductId   gocql.UUID
	ProductName string
	Quantity    int32
	Price       float64
}

type OrderDAL struct {
	Session *gocql.Session
}

func NewOrderDAL(session *gocql.Session) *OrderDAL {
	return &OrderDAL{Session: session}
}

func (dal *OrderDAL) CreateOrder(order Order) error {
	// Convert OrderItems to a list of maps
	var products []map[string]interface{}
	for _, item := range order.Products {
		products = append(products, map[string]interface{}{
			"product_id":   item.ProductId,
			"product_name": item.ProductName,
			"quantity":     item.Quantity,
			"price":        item.Price,
		})
	}

	// Insert order into the orders table
	err := dal.Session.Query(`INSERT INTO orders (order_id, customer_id, order_date, status, total_price, products, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		order.OrderId, order.CustomerId, order.OrderDate, order.Status, order.TotalPrice, products, order.CreatedAt, order.UpdatedAt, order.DeletedAt).Exec()
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return err
	}

	// Update the stock of the products
	// for _, item := range order.Products {
	// 	// update quantity of the product
	// 	err = dal.Session.Query(`UPDATE product SET quantity = quantity - ? WHERE product_id = ?`, item.Quantity, item.ProductId).Exec()
	// 	if err != nil {
	// 		log.Printf("Failed to update product stock: %v", err)
	// 		return err
	// 	}
	// 	// update sold of the product
	// 	err = dal.Session.Query(`UPDATE product SET sold = sold + ? WHERE product_id = ?`, item.Quantity, item.ProductId).Exec()
	// 	if err != nil {
	// 		log.Printf("Failed to update product sold: %v", err)
	// 		return err
	// 	}
	// }

	return nil
}

func (dal *OrderDAL) GetOrder(orderId gocql.UUID) (*Order, error) {
	// Retrieve the order
	var order Order
	var products []map[string]interface{}
	err := dal.Session.Query(`SELECT order_id, customer_id, order_date, status, total_price, products, created_at, updated_at, deleted_at
	FROM orders WHERE order_id = ?`,
		orderId).Scan(&order.OrderId, &order.CustomerId, &order.OrderDate, &order.Status, &order.TotalPrice, &products, &order.CreatedAt, &order.UpdatedAt, &order.DeletedAt)
	if err != nil {
		log.Printf("Failed to get order: %v", err)
		return nil, err
	}

	// Convert products to OrderItem structs
	for _, p := range products {
		order.Products = append(order.Products, OrderItem{
			ProductId:   p["product_id"].(gocql.UUID),
			ProductName: p["product_name"].(string),
			Quantity:    int32(p["quantity"].(int)),
			Price:       p["price"].(float64),
		})
	}

	return &order, nil
}

func (dal *OrderDAL) GetAllOrders() ([]Order, error) {
	// Retrieve all orders
	var orders []Order
	iter := dal.Session.Query(`SELECT order_id, customer_id, order_date, status, total_price, products, created_at, updated_at, deleted_at FROM orders`).Iter()
	for {
		var order Order
		var products []map[string]interface{}
		if !iter.Scan(&order.OrderId, &order.CustomerId, &order.OrderDate, &order.Status, &order.TotalPrice, &products, &order.CreatedAt, &order.UpdatedAt, &order.DeletedAt) {
			break
		}

		// Convert products to OrderItem structs
		for _, p := range products {
			order.Products = append(order.Products, OrderItem{
				ProductId:   p["product_id"].(gocql.UUID),
				ProductName: p["product_name"].(string),
				Quantity:    int32(p["quantity"].(int)),
				Price:       p["price"].(float64),
			})
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (dal *OrderDAL) UpdateOrderStatus(orderId gocql.UUID, status string) error {
	err := dal.Session.Query(`UPDATE orders SET status = ? updated_at = ?	 WHERE order_id = ?`, status, helper.GetTimeNowInGMT7(), orderId).Exec()
	if err != nil {
		log.Printf("Failed to update order status: %v", err)
		return err
	}

	return nil
}
