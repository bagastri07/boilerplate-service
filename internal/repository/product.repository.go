package repository

import (
	"context"

	"github.com/bagastri07/boilerplate-service/internal/model"
	"github.com/bagastri07/boilerplate-service/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) model.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	err := r.DB.WithContext(ctx).Save(product).Error
	if err != nil {
		logrus.WithContext(ctx).WithField("product", utils.Dump(product)).Error(err)
		return err
	}

	return nil
}
