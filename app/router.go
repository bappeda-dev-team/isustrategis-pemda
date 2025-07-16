package app

import (
	"isustrategisService/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(csfController controller.CsfController, outcomeController controller.OutcomeController, intermediateController controller.IntermediateController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/csf", csfController.Create)
	e.PUT("/csf/:id", csfController.Update)
	e.DELETE("/csf/:id", csfController.Delete)
	e.GET("/csf/detail/:csfId", csfController.FindById)
	e.GET("/csf/:pohonId", csfController.FindAll)

	e.POST("/outcome", outcomeController.Create)
	e.PUT("/outcome/:id", outcomeController.Update)
	e.DELETE("/outcome/:id", outcomeController.Delete)
	e.GET("/outcome/detail/:id", outcomeController.FindById)
	e.GET("/outcome/:pohonId", outcomeController.FindAll)

	e.POST("/intermediate", intermediateController.Create)
	e.PUT("/intermediate/:id", intermediateController.Update)
	e.DELETE("/intermediate/:pohon_id", intermediateController.Delete)
	e.GET("/intermediate/detail/:id", intermediateController.FindById)
	e.GET("/intermediate/:pohon_id", intermediateController.FindAll)

	return e
}
