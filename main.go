package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ke1ruuu/lyvora-server/internal/api"
	"github.com/ke1ruuu/lyvora-server/internal/db"
)

func main() {
	db.InitBadger()
	defer db.CloseBadger()

	r := gin.Default()
	r.Use(corsMiddleware())

	// Auth
	r.POST("/api/register", api.Register)
	r.POST("/api/login", api.Login)

	// Tracks (public)
	r.GET("/api/tracks", api.GetTracks)
	r.GET("/api/stream/:id", api.StreamTrack)

	// Favorites (protected)
	auth := r.Group("/api")
	auth.Use(api.AuthMiddleware())
	{
		auth.GET("/favorites", api.GetFavorites)
		auth.POST("/favorites", api.AddFavorite)
		auth.DELETE("/favorites/:id", api.RemoveFavorite)
	}

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
