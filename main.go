package main

import (
	"fibunacci/apis"
	"github.com/gin-gonic/gin"
	"log"

	_ "fibunacci/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Fibonacci API
// @version 1.0
// @description Swagger API for Golang Project Fibonacci.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email contact@gmail.com
// @license.name APACHE LICENSE, VERSION 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0
// @BasePath /api/v1
func main() {

	log.Println("Started application...")

	// Create a instance of GIN
	r := gin.New()

	// Route for swagger
	// TODO secure swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Route for the api
	// TODO secure api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/fseries/:count", apis.GetFibonacciSeries)
		v1.GET("/fzero/", apis.GetAllFibonacciSeries)
		v1.GET("/fnumber/:ordinal", apis.GetFibonacciNumberForOrdinal)
	}

	// All routes configured. Bare bones set up without any other middleware components
	r.Run()
}
