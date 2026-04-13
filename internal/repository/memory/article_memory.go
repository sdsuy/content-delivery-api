package memory

import (
	"errors"
	"sync"

	"github.com/sdsuy/content-delivery-api/internal/domain"
)

type ArticleMemoryRepository struct {
	data map[string]*domain.Article
	mu   sync.RWMutex
}

func NewArticleMemoryRepository() *ArticleMemoryRepository {
	return &ArticleMemoryRepository{
		data: make(map[string]*domain.Article),
	}
}

func (r *ArticleMemoryRepository) Create(article *domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[article.ID] = article
	return nil
}

func (r *ArticleMemoryRepository) GetByID(id string) (*domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	article, ok := r.data[id]
	if !ok {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (r *ArticleMemoryRepository) GetAll() ([]*domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Article
	for _, a := range r.data {
		result = append(result, a)
	}
	return result, nil
}

func (r *ArticleMemoryRepository) Update(article *domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[article.ID] = article
	return nil
}

func (r *ArticleMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, id)
	return nil
}
