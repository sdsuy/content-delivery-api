package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/sdsuy/content-delivery-api/internal/delivery/http"
)

func main() {
	router := httpDelivery.NewRouter()

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
