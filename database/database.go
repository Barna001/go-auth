package database

import (
	"fmt"

	"github.com/Barna001/go-auth/user"
)

// Database includes getting a user by email and adding a new user
type Database interface {
	GetUser(string) user.User
	AddUser(user.User)
}

// TextDB is one implementation of Database, stores in plain text, in the given path
type TextDB struct {
	Location string
}

// GetUser returns the user with the given email
func (db TextDB) GetUser(email string) user.User {
	fmt.Println("getuser")
	return user.User{Email: "b@b.b", Name: "Barna", Password: "kektusmaximus"}
}

// AddUser adds a new user to the db
func (db TextDB) AddUser(user user.User) {
	fmt.Println("adduser")
}
