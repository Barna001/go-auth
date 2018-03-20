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

	"github.com/rs/cors"
)

// Server with port
type Server struct {
	Port       int
	Db         database.Database
	JwtSignKey string
}

// StartServer creates a DefaultServeMux server with the given port
func (server Server) StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", server.handleLogin)
	mux.HandleFunc("/user", server.handleUser)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"},
	})
	handlers := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(server.Port), handlers))
}

func (server Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		user, err := getUserFromBody(w, r)
		if err != nil {
			return
		}

		dbUser := handleGetUserFromDB(w, user.Email, server.Db)
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
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r, server)
	case http.MethodPost:
		handlePostUser(w, r, server.Db)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request, server Server) {
	claims, err := getClaimsFromToken(getJwtTokenFromHeader(r.Header), server.JwtSignKey)
	fmt.Println("claims", claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	emails := r.URL.Query()["email"]
	if len(emails) == 0 {
		users, err := server.Db.GetUsers()
		if err != nil {
			http.Error(w, "Can not read from DB", http.StatusInternalServerError)
		}
		usersJSON, _ := json.Marshal(users)
		fmt.Fprintf(w, string(usersJSON))
		return
	}
	user := handleGetUserFromDB(w, emails[0], server.Db)
	if user.Email != "" {
		userJSON, _ := json.Marshal(user)
		fmt.Fprintf(w, string(userJSON))
	}
}

func handleGetUserFromDB(w http.ResponseWriter, email string, db database.Database) user.User {
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

func handlePostUser(w http.ResponseWriter, r *http.Request, db database.Database) {
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
