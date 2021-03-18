package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping-app/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on port 8181...")

	log.Fatal(http.ListenAndServe(":8181", r))
}
