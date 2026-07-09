
package main

import (
	"cinema-ticket-booking/backend/config"
	"cinema-ticket-booking/backend/database"
	"cinema-ticket-booking/backend/routers"
	"cinema-ticket-booking/backend/services"
)

func main() {
	dbConfig := config.LoadDBConfig()
	dbPool := database.InitPool(dbConfig)
	defer dbPool.Close()

	authService := &services.Register{
		DB: dbPool,
	}

	routers.Routers(authService)
}