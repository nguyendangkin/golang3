package database

import (
	"fmt"
	"log"
	"natasha/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo *gorm.DB

func Connect() {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Data.Database.Host,
		config.Data.Database.Port,
		config.Data.Database.User,
		config.Data.Database.Password,
		config.Data.Database.Name,
		config.Data.Database.SSLMode,
	)

	var err error

	Repo, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Database connected")

}
