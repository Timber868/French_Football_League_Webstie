package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Timber868/french-league/service/user"
	"github.com/gorilla/mux"
)

// Stuct that defines the API server it has both its address and also a poinbter to the database
type APIServer struct {
	addr string
	db   *sql.DB
}

// Function that creates a new API server struct based on inputs
func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	//Api v1 subrouter
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoute(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
