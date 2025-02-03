package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/funnyfoxd/docker_pinger/models"
)

var DB *gorm.DB

func InitDB() error {
	dbUrl := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Fail to connect to DB: %v", err)
		return err
	}

	DB.AutoMigrate(&models.ContainerPing{})
	return nil
}