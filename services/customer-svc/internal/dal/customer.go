package dataaccesslayer

import (
	"log"

	"github.com/gocql/gocql"
)

type Customer struct {
	CustomerId gocql.UUID
	Name       string
	Email      string
	Password   string
	Token      string
}

type CustomerDAL struct {
	Session *gocql.Session
}

func NewCustomerDAL(session *gocql.Session) *CustomerDAL {
	return &CustomerDAL{Session: session}
}

func (dal *CustomerDAL) CreateCustomer(user Customer) error {
	err := dal.Session.Query(`INSERT INTO "customer" (CustomerId, Name, Email, Password, tokenstr) VALUES (?, ?, ?, ?, ?)`,
		user.CustomerId, user.Name, user.Email, user.Password, user.Token).Exec()
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return err
	}
	return nil
}

func (dal *CustomerDAL) GetCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	err := dal.Session.Query(`SELECT customerid, name, email, password, tokenstr FROM customer WHERE email = ? ALLOW FILTERING`, email).Scan(
		&customer.CustomerId, &customer.Name, &customer.Email, &customer.Password, &customer.Token)
	if err != nil {
		log.Printf("Failed to get customer by email: %v", err)
		return nil, err
	}

	return &customer, nil
}

func (dal *CustomerDAL) GetCustomerByToken(token string) (*Customer, error) {
	var customer Customer
	err := dal.Session.Query(`SELECT CustomerId, Name, Email, Password, tokenstr FROM "customer" WHERE Token = ?`, token).Scan(
		&customer.CustomerId, &customer.Name, &customer.Email, &customer.Password, &customer.Token)
	if err != nil {
		log.Printf("Failed to get customer by token: %v", err)
		return nil, err
	}

	return &customer, nil
}

func (dal *CustomerDAL) UpdateCustomerToken(uuid string, token string) error {
	err := dal.Session.Query(`UPDATE customer SET tokenstr = ? WHERE customerid = ?`, token, uuid).Exec()
	if err != nil {
		log.Printf("Failed to update customer token: %v", err)
		return err
	}
	return nil
}

func (dal *CustomerDAL) DeleteCustomerToken(token string) error {
	err := dal.Session.Query(`UPDATE "customer" SET tokenstr = ? WHERE tokenstr = ?`, "", token).Exec()
	if err != nil {
		log.Printf("Failed to delete customer token: %v", err)
		return err
	}
	return nil
}
