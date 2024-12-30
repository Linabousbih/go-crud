package main

import (
	"gin-gorm-crud/config"
	"gin-gorm-crud/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	//Connect to MongoDB
	config.ConnectDatabase()

	//Initialize the Gin router

	r := gin.Default()
	routes.SetupRoutes(r)

	//Start the server
	r.Run(":" + port)
	log.Println("Server is running on http://localhost:8080")

}
