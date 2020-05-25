package router

import (
	"github.com/gorilla/mux"
	"github.com/projmanserver/middleware"
	"github.com/projmanserver/middleware/auth"
	"github.com/projmanserver/middleware/projects"
)

// NewRouter creates a router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Test apis
	router.HandleFunc("/test", middleware.TestFunc)
	router.HandleFunc("/test/readReq", middleware.ReadRequest)
	router.HandleFunc("/auth/testpost", middleware.TestPost).Methods("POST", "OPTIONS")
	router.Handle("/test/authReq", auth.JwtVerify(middleware.AuthRequest))
	router.HandleFunc("/test/path/{id}", middleware.TestPathParams).Methods("GET", "OPTIONS")

	// Auth apis
	router.HandleFunc("/auth/signup", auth.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/login", auth.Login)

	// Project apis
	router.HandleFunc("/projects/add", projects.AddProject).Methods("POST", "OPTIONS")
	router.HandleFunc("/projects/getAll/{email}", projects.AddProject).Methods("GET", "OPTIONS")
	return router
}
