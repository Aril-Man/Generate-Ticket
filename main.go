package main

import (
	"log"
	"ticket-backend-gh/controller"
	"ticket-backend-gh/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Middleware())
	apiGroup := router.Group("/api/wetalk")

	handler := &controller.Controller{}

	handler.MapEndpoints(apiGroup)

	router.Run(":8080")
}
