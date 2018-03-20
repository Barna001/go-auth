package database

import (
	"io/ioutil"
	"os"

	"github.com/Barna001/go-auth/errors"
	"github.com/Barna001/go-auth/user"
)

// Database includes getting a user by email and adding a new user
type Database interface {
	GetUsers() ([]user.User, error)
	GetUser(string) (user.User, error)
	AddUser(user.User) error
}

// TextDB is one implementation of Database, stores in plain text, in the given path
type TextDB struct {
	Location string
}

// GetUsers returns all users
func (db TextDB) GetUsers() ([]user.User, error) {
	prewUsersBytes, err := ioutil.ReadFile(db.Location)
	if err != nil {
		return []user.User{}, err
	}
	users := user.Deserialize(string(prewUsersBytes))
	return users, nil
}

// GetUser returns the user with the given email
func (db TextDB) GetUser(email string) (user.User, error) {
	prewUsersBytes, err := ioutil.ReadFile(db.Location)
	if err != nil {
		return user.User{}, err
	}
	users := user.Deserialize(string(prewUsersBytes))
	for _, curr := range users {
		if curr.Email == email {
			return curr, nil
		}
	}
	return user.User{}, errors.NoUserError{}
}

// AddUser adds a new user to the db
func (db TextDB) AddUser(user user.User) error {
	_, err := db.GetUser(user.Email)
	if err == nil {
		return errors.AlreadyAddedError{}
	}
	serialzedUser := []byte(user.Serialize())

	dbFile, err := os.OpenFile(db.Location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errors.CriticalHandling(err)
	defer dbFile.Close()

	_, err = dbFile.Write(serialzedUser)
	errors.CriticalHandling(err)

	return nil
}
