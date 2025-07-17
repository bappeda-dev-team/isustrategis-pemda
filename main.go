package main

import (
	"fmt"
	"isustrategisService/app"
	"isustrategisService/docs"
	"os"

	"isustrategisService/helper"

	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) *echo.Echo {
	return e
}

// @title Isu Strategis Pemda
// @version 1.0
// @description API For Isu Strategis Pemda
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host ${HOST}:${PORT}
// @BasePath /

func main() {

	// app.RunFlyway()

	server := InitializedServer()
	host := os.Getenv("host")
	port := os.Getenv("port")

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)

	addr := fmt.Sprintf("%s:%s", host, port)

	err := server.Start(addr)
	helper.PanicIfError(err)
}
