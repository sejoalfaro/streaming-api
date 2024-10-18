package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sejoalfaro/streaming-api/database"
	"github.com/sejoalfaro/streaming-api/routes"
)

func main() {
	router := gin.Default()
	router.GET("/health", routes.HealthCheck)
	router.Run(":8080")
}
