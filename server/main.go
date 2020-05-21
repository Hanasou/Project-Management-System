package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/projmanserver/router"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
