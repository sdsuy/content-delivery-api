package repository

import "github.com/sdsuy/content-delivery-api/internal/domain"

type ArticleRepository interface {
	Create(article *domain.Article) error
	GetByID(id string) (*domain.Article, error)
	GetAll() ([]*domain.Article, error)
	Update(article *domain.Article) error
	Delete(id string) error
}
