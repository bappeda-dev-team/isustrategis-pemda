package main

import (
	"fmt"
	"isustrategisService/app"
	"isustrategisService/docs"
	"os"

	"isustrategisService/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(e *echo.Echo) *echo.Echo {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Atau tentukan domain spesifik, misal: []string{"http://localhost:3000", "https://yourdomain.com"}
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	return e
}

// @title Permasalahan & Isu Strategis Service API
// @version 1.0
// @description API For Permasalahan & Isu Strategis Services
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host ${HOST}:${PORT}
// @BasePath /

func main() {

	app.RunFlyway()

	server := InitializedServer()
	host := os.Getenv("host")
	port := os.Getenv("port")

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)

	addr := fmt.Sprintf("%s:%s", host, port)

	err := server.Start(addr)
	helper.PanicIfError(err)
}
