package database

import (
	"cinema-ticket-booking/backend/config"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPool(cfg config.DBConfig) *sql.DB{
	db, err := sql.Open("pgx", cfg.ToDSN())
	if err!=nil{
		log.Fatalf("Database config parsing failed: %v", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failure! Could not ping Postgres: %v", err)
	}

	log.Println("Database connection pool established successfully.")
	return db
}