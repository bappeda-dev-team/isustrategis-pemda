//go:build wireinject
// +build wireinject

package main

import (
	"isustrategisService/app"
	"isustrategisService/controller"
	"isustrategisService/repository"
	"isustrategisService/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var csfSet = wire.NewSet(
	repository.NewCsfRepositoryImpl,
	wire.Bind(new(repository.CsfRepository), new(*repository.CsfRepositoryImpl)),
	service.NewCsfServiceImpl,
	wire.Bind(new(service.CsfService), new(*service.CsfServiceImpl)),
	controller.NewCsfControllerImpl,
	wire.Bind(new(controller.CsfController), new(*controller.CsfControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		csfSet,
		app.NewRouter,
	)
	return nil
}
