package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

type body struct {
	Key1 string
	Key2 string
}

// TestFunc test functionality of server
func TestFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from "+r.URL.Path[1:])
}

// TestPathParams tests getting path parameters
func TestPathParams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Getting path parameter")
	params := mux.Vars(r)
	log.Println(params)
	json.NewEncoder(w).Encode(params)
}

// TestPost is a dummy post request
func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("starting test post request")
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	log.Println(user)
	json.NewEncoder(w).Encode(user)
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
