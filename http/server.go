package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Barna001/go-auth/database"
	"github.com/Barna001/go-auth/errors"
	"github.com/Barna001/go-auth/user"
)

// Server with port
type Server struct {
	Port       int
	Db         database.Database
	JwtSignKey string
}

// StartServer creates a DefaultServeMux server with the given port
func (server Server) StartServer() {
	http.HandleFunc("/user", server.handleUser)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.Port), nil))
}

func (server Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("get")
		claims, err := getClaimsFromToken(getJwtTokenFromHeader(r.Header), server.JwtSignKey)
		fmt.Println(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		handleGet(w, r, server.Db)
	case http.MethodPost:
		handlePost(w, r, server.Db)
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
}

func handlePost(w http.ResponseWriter, r *http.Request, db database.Database) {
	decoder := json.NewDecoder(r.Body)
	var user user.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Not valid user: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := db.AddUser(user); err != nil {
		http.Error(w, "Can not add user: "+err.Error(), http.StatusNotAcceptable)
		return
	}

	userJSON, _ := json.Marshal(user)
	fmt.Fprintf(w, string(userJSON), http.StatusCreated)
}
