package main

import (
	"fmt"

	"github.com/Barna001/go-auth/database"
	"github.com/Barna001/go-auth/http"
	"github.com/Barna001/go-auth/user"
)

func main() {
	fmt.Println("hello there general kenobi")
	db := database.TextDB{Location: "users.txt"}
	db.AddUser(user.User{Email: "b@b.b", Name: "Barna", Password: "kek"})
	user, err := db.GetUser("b@b.b")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	webServer := http.Server{Port: 8080, Db: db}
	webServer.StartServer()
}
