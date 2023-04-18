package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"final-project-hacktiv8/models"
	"final-project-hacktiv8/routers"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
	)

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func main() {

	PORT := os.Getenv("DB_PORT")

	routers.StartServer(db).Run(":" + PORT)
}
