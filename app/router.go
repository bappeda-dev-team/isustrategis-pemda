package app

import (
	"isustrategisService/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(csfController controller.CsfController, outcomeController controller.OutcomeController, intermediateController controller.IntermediateController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/csf", csfController.Create)
	e.PUT("/csf/:id", csfController.Update)
	e.DELETE("/csf/:idPohon", csfController.Delete)
	e.GET("/csf/detail/:csfId", csfController.FindById)
	e.GET("/csf/:tahun", csfController.FindAll)

	e.POST("/outcome", outcomeController.Create)
	e.PUT("/outcome/:id", outcomeController.Update)
	e.DELETE("/outcome/:id", outcomeController.Delete)
	e.GET("/outcome/detail/:id", outcomeController.FindById)
	e.GET("/outcome/:tahun", outcomeController.FindAll)

	e.POST("/intermediate", intermediateController.Create)
	e.PUT("/intermediate/:id", intermediateController.Update)
	e.DELETE("/intermediate/:pohon_id", intermediateController.Delete)
	e.GET("/intermediate/detail/:id", intermediateController.FindById)
	e.GET("/intermediate/:tahun", intermediateController.FindAll)

	return e
}
