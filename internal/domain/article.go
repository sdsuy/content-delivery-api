package domain

import "time"

type Article struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
