package main

import (
	"fmt"

	"github.com/Barna001/go-auth/database"
)

func main() {
	fmt.Println("hello there general kenobi")
	db := database.TextDB{Location: "users.txt"}
	db.GetUser("b@b.b")
}
