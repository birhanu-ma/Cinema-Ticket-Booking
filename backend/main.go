package main

import (
	"cinema-ticket-booking/backend/config"
	"cinema-ticket-booking/backend/database"
	"cinema-ticket-booking/backend/routers"
	"cinema-ticket-booking/backend/services"
	"log"
	"os"
)

func main() {
	dbConfig := config.LoadDBConfig()
		
log.Printf("%+v\n", dbConfig)
log.Println(dbConfig.ToDSN())
	dbPool := database.InitPool(dbConfig)
	defer dbPool.Close()



	authService := &services.Register{
		DB: dbPool,
	}
	paymentService := &services.Payment{
		DB:          dbPool,
		ChapaSecret: os.Getenv("CHAPA_SECRET_KEY"),
	}

	routers.Routers(paymentService, authService)
}