package usecase

import (
	"context"

	"github.com/bagastri07/boilerplate-service/internal/model"
	"github.com/bagastri07/boilerplate-service/utils"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo model.ProductRepository
}

func NewProductUsecase(productRepo model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) Create(ctx context.Context, product *model.Product) error {
	err := u.productRepo.Create(ctx, product)
	if err != nil {
		logrus.WithContext(ctx).WithField("product", utils.Dump(product)).Error(err)
		return err
	}

	return nil
}
