package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/sdsuy/content-delivery-api/internal/domain"
	"github.com/sdsuy/content-delivery-api/internal/repository"
)

type ArticleUseCase struct {
	repo repository.ArticleRepository
}

func NewArticleUseCase(r repository.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{repo: r}
}

func (u *ArticleUseCase) Create(title, content string) (*domain.Article, error) {
	article := &domain.Article{
		ID:        uuid.NewString(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.repo.Create(article)
	return article, err
}

func (u *ArticleUseCase) GetByID(id string) (*domain.Article, error) {
	return u.repo.GetByID(id)
}

func (u *ArticleUseCase) GetAll() ([]*domain.Article, error) {
	return u.repo.GetAll()
}

func (u *ArticleUseCase) Update(id, title, content string) (*domain.Article, error) {
	article, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Content = content
	article.UpdatedAt = time.Now()

	err = u.repo.Update(article)
	return article, err
}

func (u *ArticleUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}
