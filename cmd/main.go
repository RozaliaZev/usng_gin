package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RozaliaZev/usng_gin.git/pkg/handlers"
	"github.com/RozaliaZev/usng_gin.git/pkg/middlewares"
)

func main() {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	router.Use(middlewares.SetHeader())

	router.GET("/when/:year", handlers.GetNumberDays)
	router.Run(":8080")
}



  