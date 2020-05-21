package router

import (
	"net/http"

	"github.com/projmanserver/middleware"
	"github.com/projmanserver/middleware/authentication"
)

// NewRouter creates a router
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", middleware.TestFunc)
	mux.HandleFunc("/test/readReq", middleware.ReadRequest)
	mux.HandleFunc("/auth/signup", authentication.CreateUser)
	mux.HandleFunc("/auth/login", authentication.Login)
	mux.HandleFunc("/auth/verifyJWT", authentication.JwtVerify)
	return mux
}
