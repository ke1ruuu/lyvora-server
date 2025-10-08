package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ke1ruuu/lyvora-server/internal/api"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/tracks", api.GetTracks)
		apiGroup.GET("/stream/:id", api.StreamTrack)
	}
}
