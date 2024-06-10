package internal

import "errors"

type User struct {
	ID    uint64
	Name  string
	Email string
}

type DB struct {
	users map[uint64]*User
}

func NewDB() *DB {
	return &DB{users: make(map[uint64]*User)}
}

func (db *DB) GetUser(userID uint64) (*User, error) {
	user, exists := db.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Add other database interaction methods as needed...
