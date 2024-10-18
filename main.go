package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sejoalfaro/streaming-api/database"
	"github.com/sejoalfaro/streaming-api/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
