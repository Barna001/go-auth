package user

import (
	"strings"
)

// User has email, name and password
type User struct {
	Email, Password, Name string
}

// Deserialize the Users from plain string
func Deserialize(data string) []User {
	rows := strings.Split(data, "/n")
	var users []User
	for _, row := range rows {
		parts := strings.Split(row, ";")
		user := User{Email: parts[0], Password: parts[1], Name: parts[2]}
		users = append(users, user)
	}
	return users
}

// Serialize a User to a string
func (user *User) Serialize() string {
	return user.Email + ";" + user.Password + ";" + user.Name + "\n"
}
