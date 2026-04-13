package http

import (
	"net/http"

	"github.com/sdsuy/content-delivery-api/internal/usecase"
)

func NewRouter(articleUC *usecase.ArticleUseCase) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)

	handler := NewArticleHandler(articleUC)

	mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.Create(w, r)
			return
		}
		if r.Method == http.MethodGet {
			handler.GetAll(w, r)
			return
		}
	})

	return mux
}
