package main

import (
	"os"
	"uuid-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(os.Getenv("PORT"))
}