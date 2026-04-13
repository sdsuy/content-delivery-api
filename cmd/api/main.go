package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/sdsuy/content-delivery-api/internal/delivery/http"
	"github.com/sdsuy/content-delivery-api/internal/repository/memory"
	"github.com/sdsuy/content-delivery-api/internal/usecase"
)

func main() {
	repo := memory.NewArticleMemoryRepository()
	usecase := usecase.NewArticleUseCase(repo)

	router := httpDelivery.NewRouter(usecase)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
