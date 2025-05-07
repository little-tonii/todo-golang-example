package main

import (
	"fmt"
	"io"
	"os"
	"todo-golang-example/internal/infrastructure/config"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if errors := config.LoadEnvironment(); errors != nil && len(errors) > 0 {
		panic(fmt.Sprintf("%v", errors))
	}

	if error := config.InitializeDatabase(); error != nil {
		panic(error)
	}
	defer config.CloseDatabase()

	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
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
