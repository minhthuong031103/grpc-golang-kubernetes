package dal

import (
	"log"

	"github.com/gocql/gocql"
)

type Product struct {
	ProductID   gocql.UUID
	ProductName string
	Price       float64
	Description string
	Quantity    int32
	Sold        int32
	ImageURL    string
}

type ProductDAL struct {
	Session *gocql.Session
}

func NewProductDAL(session *gocql.Session) *ProductDAL {
	return &ProductDAL{Session: session}
}

func (dal *ProductDAL) CreateProduct(product Product) error {
	err := dal.Session.Query(`INSERT INTO product (product_id, product_name, price, description, quantity, sold, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		product.ProductID, product.ProductName, product.Price, product.Description, product.Quantity, product.Sold, product.ImageURL).Exec()
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return err
	}
	return nil
}

func (dal *ProductDAL) GetProductByID(productID gocql.UUID) (*Product, error) {
	var product Product
	err := dal.Session.Query(`SELECT product_id, product_name, price, description, quantity, sold, image_url FROM product WHERE product_id = ?`,
		productID).Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Description, &product.Quantity, &product.Sold, &product.ImageURL)
	if err != nil {
		log.Printf("Failed to get product by ID: %v", err)
		return nil, err
	}
	return &product, nil
}

func (dal *ProductDAL) GetAllProducts() ([]Product, error) {
	var products []Product
	iter := dal.Session.Query(`SELECT product_id, product_name, price, description, quantity, sold, image_url FROM product`).Iter()
	defer iter.Close()

	var product Product
	for iter.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Description, &product.Quantity, &product.Sold, &product.ImageURL) {
		products = append(products, product)
	}
	if err := iter.Close(); err != nil {
		log.Printf("Failed to get all products: %v", err)
		return nil, err
	}
	return products, nil
}
