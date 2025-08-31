package main

import (
	"context"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/config"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Static("/uploads", "./uploads") // Serve files in the uploads folder at /uploads URL path

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConfig.App.Port),
		Handler: server,
	}

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()

	log.Printf("✅ Server is running on port %d\n", config.AppConfig.App.Port)

	// Listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown HTTP server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close DB (important)
	sqlDB, err := db.DB.DB()
	if err == nil {
		_ = sqlDB.Close()
	}

	log.Println("🚀 Server exited properly")

}
