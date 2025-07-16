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

var outcomeSet = wire.NewSet(
	repository.NewOutcomeRepositoryImpl,
	wire.Bind(new(repository.OutcomeRepository), new(*repository.OutcomeRepositoryImpl)),
	service.NewOutcomeServiceImpl,
	wire.Bind(new(service.OutcomeService), new(*service.OutcomeServiceImpl)),
	controller.NewOutcomeControllerImpl,
	wire.Bind(new(controller.OutcomeController), new(*controller.OutcomeControllerImpl)),
)

var intermediateSet = wire.NewSet(
	repository.NewIntermediateRepositoryImpl,
	wire.Bind(new(repository.IntermediateRepository), new(*repository.IntermediateRepositoryImpl)),
	service.NewIntermediateServiceImpl,
	wire.Bind(new(service.IntermediateService), new(*service.IntermediateServiceImpl)),
	controller.NewIntermediateControllerImpl,
	wire.Bind(new(controller.IntermediateController), new(*controller.IntermediateControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		csfSet,
		outcomeSet,
		intermediateSet,
		app.NewRouter,
	)
	return nil
}
