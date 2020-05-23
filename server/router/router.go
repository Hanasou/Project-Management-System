package router

import (
	"net/http"

	"github.com/projmanserver/middleware"
	"github.com/projmanserver/middleware/auth"
)

// NewRouter creates a router
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", middleware.TestFunc)
	mux.HandleFunc("/test/readReq", middleware.ReadRequest)
	mux.Handle("/test/authReq", auth.JwtVerify(middleware.AuthRequest))
	mux.HandleFunc("/auth/signup", auth.CreateUser)
	mux.HandleFunc("/auth/login", auth.Login)
	return mux
}
