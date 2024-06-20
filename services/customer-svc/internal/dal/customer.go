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
	Role       string
}

type CustomerDAL struct {
	Session *gocql.Session
}

func NewCustomerDAL(session *gocql.Session) *CustomerDAL {
	return &CustomerDAL{Session: session}
}

func (dal *CustomerDAL) CreateCustomer(user Customer) error {
	err := dal.Session.Query(`INSERT INTO "customer" (CustomerId, Name, Email, Password, tokenstr, role) VALUES (?, ?, ?, ?, ?, 'customer')`,
		gocql.TimeUUID(), user.Name, user.Email, user.Password, user.Token).Exec()
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return err
	}
	return nil
}

func (dal *CustomerDAL) GetCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	err := dal.Session.Query(`SELECT customerid, name, email, password, tokenstr, role FROM customer WHERE email = ? ALLOW FILTERING`, email).Scan(
		&customer.CustomerId, &customer.Name, &customer.Email, &customer.Password, &customer.Token, &customer.Role)
	if err != nil {
		log.Printf("Failed to get customer by email %s: %v", email, err)
		return nil, err
	}

	return &customer, nil
}

func (dal *CustomerDAL) GetCustomerByToken(token string) (*Customer, error) {
	var customer Customer
	err := dal.Session.Query(`SELECT CustomerId, Name, Email, Password, tokenstr, role FROM "customer" WHERE tokenstr = ?`, token).Scan(
		&customer.CustomerId, &customer.Name, &customer.Email, &customer.Password, &customer.Token, &customer.Role)
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

func (dal *CustomerDAL) GetAccountById(uuid string) (*Customer, error) {
	var customer Customer
	err := dal.Session.Query(`SELECT customerid, name, email, password, tokenstr, role FROM customer WHERE customerid = ?`, uuid).Scan(
		&customer.CustomerId, &customer.Name, &customer.Email, &customer.Password, &customer.Token, &customer.Role)
	if err != nil {
		log.Printf("Failed to get customer by id %s: %v", uuid, err)
		return nil, err
	}

	return &customer, nil
}

func (dal *CustomerDAL) UpdateRole(uuid string, role string) error {
	if role != "customer" && role != "admin" {
		log.Printf("Invalid role %s", role)
		return nil
	}

	err := dal.Session.Query(`UPDATE customer SET role = ? WHERE customerid = ?`, role, uuid).Exec()
	if err != nil {
		log.Printf("Failed to update customer role: %v", err)
		return err
	}
	return nil
}
