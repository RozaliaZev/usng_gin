package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RozaliaZev/usng_gin.git/pkg/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/when/:year", handlers.GetNumberDays)
	router.Run(":8080")
}



  