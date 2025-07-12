package main

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":5050")

}
