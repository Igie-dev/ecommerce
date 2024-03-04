package database

import (
	"fmt"
	"log"

	"github.com/laptop-shop.api/config"
	"github.com/laptop-shop.api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var defaultDB = &DB{}

func (db *DB) connection(config *config.DBConfig) (err error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Manila", config.Host, config.User, config.Password, config.DbName, config.Port, config.SSLMode)

	db.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Database connection error: %v", err)
		return err
	}

	db.AutoMigrate(&model.Customer{})

	return nil

}

// GetDB returns db instance
func GetDB() *DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration
func ConnectDB() error {
	return defaultDB.connection(config.GetDbConfig())
}
