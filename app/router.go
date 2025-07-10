package app

import (
	"isustrategisService/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(csfController controller.CsfController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/csf", csfController.Create)
	e.PUT("/csf/:id", csfController.Update)
	e.DELETE("/csf/:id", csfController.Delete)
	e.GET("/csf/detail/:csfId", csfController.FindById)
	e.GET("/csf/:pohonId", csfController.FindAll)

	return e
}
