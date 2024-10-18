package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sejoalfaro/streaming-api/database"
)

func HealthCheck(c *gin.Context) {
	// Verifica la conexión a la base de datos
	sqlDB, err := db.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Can't connect to the database"})
		return
	}

	// Ping a la base de datos para asegurarse de que está conectada
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Can't connect to the database"})
		return
	}

	// Si todo está bien, devolver el estado "healthy"
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "message": "OK!"})
}
