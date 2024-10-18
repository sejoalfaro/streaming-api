package database

import (
	"github.com/sejoalfaro/streaming-api/models"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDataBase() {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DSN:        "libsql://gorm-test-sejoalfaro.turso.io/?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MjkyODg1MzYsImlkIjoiMTk4Mzk3MWUtNDI2MC00ZTc0LWE5YzYtZWIwNGViYzUzYjRkIn0.IyhCQjRPm2XsWb6MG3c046FcWm8qxNgYfueT8bf_mge7EopnRawE2nSYOK0N-UYte4BG8iPR7iIi1b_Gg9-XCA",
		DriverName: "libsql",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func init() {
	initDataBase()
	DB.AutoMigrate(&models.User{}, &models.Stream{})

}
