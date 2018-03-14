package user

import (
	"strings"
)

// User has email, name and password
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Deserialize the Users from plain string
func Deserialize(data string) []User {
	rows := strings.Split(data, "\n")
	rows = rows[0 : len(rows)-1]
	var users []User
	for _, row := range rows {
		row = strings.Replace(row, "\n", "", -1)
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
