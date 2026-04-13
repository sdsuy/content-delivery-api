package http

import (
	"encoding/json"
	"net/http"

	"github.com/sdsuy/content-delivery-api/internal/usecase"
)

type ArticleHandler struct {
	usecase *usecase.ArticleUseCase
}

func NewArticleHandler(u *usecase.ArticleUseCase) *ArticleHandler {
	return &ArticleHandler{usecase: u}
}

func (h *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string
		Content string
	}

	json.NewDecoder(r.Body).Decode(&req)

	article, err := h.usecase.Create(req.Title, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	articles, _ := h.usecase.GetAll()
	json.NewEncoder(w).Encode(articles)
}
