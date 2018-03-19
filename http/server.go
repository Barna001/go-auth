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
	http.HandleFunc("/login", server.handleLogin)
	http.HandleFunc("/user", server.handleUser)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.Port), nil))
}

func (server Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodPost:
		user, err := getUserFromBody(w, r)
		if err != nil {
			return
		}

		dbUser := handleGetUser(w, user.Email, server.Db)
		if user.Password == dbUser.Password {
			token := createTokenForEndpoints(server.JwtSignKey, user.Email)
			fmt.Fprintf(w, token)
		} else {
			http.Error(w, "Invalid email or password", http.StatusForbidden)
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (server Server) handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		claims, err := getClaimsFromToken(getJwtTokenFromHeader(r.Header), server.JwtSignKey)
		fmt.Println(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		emails := r.URL.Query()["email"]
		if len(emails) == 0 {
			http.Error(w, "You have to give an email address", http.StatusBadRequest)
		}
		user := handleGetUser(w, emails[0], server.Db)
		if user.Email != "" {
			userJSON, _ := json.Marshal(user)
			fmt.Fprintf(w, string(userJSON))
		}
	case http.MethodPost:
		handlePost(w, r, server.Db)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGetUser(w http.ResponseWriter, email string, db database.Database) user.User {
	user, err := db.GetUser(email)
	if err != nil {
		switch err.(type) {
		case errors.NoUserError:
			http.Error(w, "No user with this email", http.StatusNotFound)
		default:
			http.Error(w, "Can not read from DB", http.StatusInternalServerError)
		}
	}
	return user
}

func handlePost(w http.ResponseWriter, r *http.Request, db database.Database) {
	user, err := getUserFromBody(w, r)
	if err != nil {
		return
	}

	if err := db.AddUser(user); err != nil {
		http.Error(w, "Can not add user: "+err.Error(), http.StatusNotAcceptable)
		return
	}

	userJSON, _ := json.Marshal(user)
	fmt.Fprintf(w, string(userJSON), http.StatusCreated)
}

func getUserFromBody(w http.ResponseWriter, r *http.Request) (user.User, error) {
	decoder := json.NewDecoder(r.Body)
	var parsedUser user.User
	if err := decoder.Decode(&parsedUser); err != nil {
		http.Error(w, "Not valid user: "+err.Error(), http.StatusBadRequest)
		return user.User{}, err
	}
	defer r.Body.Close()

	return parsedUser, nil
}
