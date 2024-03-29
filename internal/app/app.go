package app

import (
	"context"
	"github.com/core-go/health"
	mgo "github.com/core-go/health/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handler"
	"go-service/internal/service"
)

type ApplicationContext struct {
	Health *health.Handler
	User   *handler.UserHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(conf.Mongo.Database)

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	mongoChecker := mgo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
