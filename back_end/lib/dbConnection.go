package lib

import (
	"fmt"
	"log"

	"github.com/laptop-shop.api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(config *config.DBConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Manila", config.Host, config.User, config.Password, config.DbName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Database connection error: %v", err)
		return nil, err
	}
	return db, nil

}
