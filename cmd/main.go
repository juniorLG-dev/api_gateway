package main

import (
	"gateway/internal/configuration/database"
	"gateway/internal/configuration/http_launcher"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	db, err := database.SetupDB()
	if err != nil {
		log.Panic(err)
	}

	router := gin.Default()

	http_launcher.InitRegisterRoutes(db, &router.RouterGroup)
	http_launcher.InitDynamicRouting(db, &router.RouterGroup)

	router.Run(":8080")
}
