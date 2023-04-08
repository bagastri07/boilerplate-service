package bootstrap

import (
	"context"
	"fmt"
	"net"

	grpcTransport "github.com/bagastri07/boilerplate-service/internal/delivery/grpc"
	pb "github.com/bagastri07/boilerplate-service/pb/boilerplate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/bagastri07/boilerplate-service/internal/config"
	"github.com/bagastri07/boilerplate-service/internal/constant"
	"github.com/bagastri07/boilerplate-service/internal/database"
	"github.com/bagastri07/boilerplate-service/internal/repository"
	"github.com/bagastri07/boilerplate-service/internal/usecase"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	database.InitializePostgresConn()

	pgDB, err := database.PostgreSQL.DB()
	continueOrFatal(err)

	// init repositories
	productRepository := repository.NewProductRepository(database.PostgreSQL)

	// init usecases
	productUsecase := usecase.NewProductUsecase(productRepository)

	// init grpc
	grpcDelivery := grpcTransport.NewGRPCServer()
	err = grpcDelivery.InjectProductUsecase(productUsecase)
	continueOrFatal(err)

	productGRPCServer := grpc.NewServer()

	pb.RegisterProductServiceServer(productGRPCServer, grpcDelivery)
	if config.Env() == constant.EnvDevelopment {
		reflection.Register(productGRPCServer)
	}

	lis, err := net.Listen("tcp", ":"+config.GRPCPort())
	continueOrFatal(err)

	go func() {
		err = productGRPCServer.Serve(lis)
		continueOrFatal(err)
	}()

	startingMessage()

	wait := gracefulShutdown(context.Background(), config.GracefulShutdownTimeOut(), map[string]operation{
		"postgressql connection": func(ctx context.Context) error {
			return pgDB.Close()
		},
	})
	<-wait
}

func startingMessage() {
	logrus.Info(fmt.Sprintf("%s@%s is starting", config.ServiceName(), config.ServiceVersion()))
	logrus.Info(fmt.Sprintf("grpc server started on :%s", config.GRPCPort()))
}
