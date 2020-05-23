package middleware

import (
	"encoding/json"
	"io"
	"net/http"
)

type body struct {
	Key1 string
	Key2 string
}

// TestFunc test functionality of server
func TestFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from "+r.URL.Path[1:])
}

// ReadRequest func reads from request body
func ReadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	json.NewEncoder(w).Encode(b)
}

// AuthRequest only works if user is authenticated
func AuthRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	json.NewEncoder(w).Encode(b)
}
