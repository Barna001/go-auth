package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Barna001/go-auth/database"
	"github.com/Barna001/go-auth/errors"
)

// Server with port
type Server struct {
	Port int
	Db   database.Database
}

// StartServer creates a DefaultServeMux server with the given port
func (server Server) StartServer() {
	http.HandleFunc("/user", server.handleUser)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.Port), nil))
}

func (server Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r, server.Db)
	case http.MethodPost:
		fmt.Fprintf(w, "post them")
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, db database.Database) {
	email := r.URL.Query()["email"]
	if len(email) == 0 {
		http.Error(w, "You have to give an email address", http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(email[0])
	if err != nil {
		switch err.(type) {
		case errors.NoUserError:
			http.Error(w, "No user with this email", http.StatusNotFound)
			return
		default:
			http.Error(w, "Can not read from DB", http.StatusInternalServerError)
			return
		}
	}
	userJSON, _ := json.Marshal(user)
	fmt.Fprintf(w, string(userJSON))
	return
}
