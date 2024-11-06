package api

import (
	"database/sql"
	"log"
	"net/http"
	"test-project/cmd/services/products"
	"test-project/cmd/services/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// ! User
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// ! Products
	productStore := products.NewProductStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

// https://www.youtube.com/watch?v=7VLmLOiQ3ck	- 1:10:29
