package middleware

import (
	"encoding/json"
	"io"
	"net/http"
)

// TestFunc test functionality of server
func TestFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from TestFunc")
}

// ReadRequest func reads from request body
func ReadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type body struct {
		Key1 string
		Key2 string
	}
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	json.NewEncoder(w).Encode(b)
}
