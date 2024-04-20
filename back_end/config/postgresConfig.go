package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DbName   string
	SSLMode  string
}

func GetDbConfig() *DBConfig {
	var c = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return &c
}
