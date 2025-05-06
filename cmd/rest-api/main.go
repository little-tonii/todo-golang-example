package main

import (
	"fmt"
	"todo-golang-example/internal/infrastructure/config"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	errors := config.LoadEnvironment()
	if errors != nil && len(errors) > 0 {
		panic(fmt.Sprintf("%v", errors))
	}
	error := config.InitializeDatabase()
	if error != nil {
		panic(error)
	}
	defer config.CloseDatabase()
	engine := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	engine.Use(cors.New(corsConfig))
	engine.Use(middleware.Recovery())
	engine.Use(middleware.ErrorHandler())
	engine.Run(":8080")
}
