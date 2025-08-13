package main

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/config"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/middlewares"
	"github.com/Reza-Rayan/twitter-like-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server, db.DB)

	// Apply Prometheus middleware globally
	server.Use(middlewares.PrometheusMiddleware())

	// Running CLI
	Execute()

	// Conditionally expose /metrics based on config
	if config.AppConfig.Monitoring.Enabled {
		server.GET(config.AppConfig.Monitoring.Path, middlewares.PrometheusHandler())
	}

	server.Static("/uploads", "./uploads") // Serve files in the uploads folder at /uploads URL path

	port := config.AppConfig.App.Port
	server.Run(fmt.Sprintf(":%d", port))

}
