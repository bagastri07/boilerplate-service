package model

import (
	"context"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
}

type ProductUsecase interface {
	Create(ctx context.Context, product *Product) error
}
