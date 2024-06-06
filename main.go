package main

import (
	"log"

	"os"

	"dunky.com/eventbooking/db"

	"dunky.com/eventbooking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the dotenv file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db.InitDB(dataSourceName) // To initialize database and create tables
	//db.InitDB()

	server := gin.Default()

	// Registering routes
	routes.RegisterRoutes(server)

	server.Run(":9090") //localhost:9090

} // End of main()
