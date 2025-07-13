package main

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/config"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	port := config.AppConfig.App.Port
	server.Run(fmt.Sprintf(":%d", port))

}
