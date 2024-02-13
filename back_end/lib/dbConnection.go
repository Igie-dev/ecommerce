package lib

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var localhost string
var user string
var password string
var dbname string
var port string
var DB *gorm.DB

func DbConnection() {
	localhost = os.Getenv("DB_LOCALHOST")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Manila", localhost, user, password, dbname, port)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error", err)
	}
	DB = db
}

func CloseDatabase() {
	if DB != nil {
		db, _ := DB.DB()
		db.Close()
	}
}
