package grpc

import (
	"context"

	"github.com/bagastri07/boilerplate-service/internal/model"
	pb "github.com/bagastri07/boilerplate-service/pb/boilerplate"
	"github.com/bagastri07/boilerplate-service/utils"
	"github.com/sirupsen/logrus"
)

func (t *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductRespose, error) {
	err := t.productUC.Create(ctx, &model.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
	})

	if err != nil {
		logrus.WithContext(ctx).WithField("req", utils.Dump(req)).Error(err)
		return nil, err
	}

	return &pb.CreateProductRespose{
		Message: "created",
	}, nil
}
