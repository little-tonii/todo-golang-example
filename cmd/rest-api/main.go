package main

import (
	"io"
	"log"
	"os"

	"todo-golang-example/internal/infrastructure/config"
	"todo-golang-example/internal/interface/router"
	sharedConfig "todo-golang-example/internal/shared/config"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if errors := sharedConfig.LoadEnvironment(); errors != nil && len(errors) > 0 {
		log.Fatalf("Không thể tải biến môi trường: %v", errors)
	}

	if error := config.InitializeDatabase(); error != nil {
		log.Fatal(error)
	}
	defer config.CloseDatabase()

	logFile, error := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if error != nil {
		log.Fatal(error)
	}
	defer logFile.Close()

	gin.ForceConsoleColor()

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

	router.InitializeUserRouter(engine)

	if error := engine.Run(":8080"); error != nil {
		log.Fatalf("Khởi động server thất bại: %v", error)
	}
}
