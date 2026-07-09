package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadDBConfig() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("Invalid POSTGRES_PORT value: %v", err)
	}



	return DBConfig{
		Host:     os.Getenv("POSTGRES_HOST" ),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER" ),
		Password: os.Getenv( "POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB" ),
	}
}

func (c DBConfig) ToDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}