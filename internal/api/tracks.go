package api

import (
	"encoding/json"
	"net/http"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/ke1ruuu/lyvora-server/internal/models"
	"log"
	"fmt"
)

func GetTracks(c *gin.Context) {
	resp, err := http.Get("https://discoveryprovider.audius.co/v1/tracks/trending?app_name=lyvora")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
		return
	}
	defer resp.Body.Close()

	var data struct {
		Data []map[string]interface{} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid json"})
		return
	}

	tracks := []models.Track{}

	for _, obj := range data.Data {
		// Safely extract artwork URL
		artwork := ""
		if art, ok := obj["artwork"].(map[string]interface{}); ok {
			// pick highest resolution available
			if val, ok := art["150x150"].(string); ok {
				artwork = val
			} else if val, ok := art["480x480"].(string); ok {
				artwork = val
			} else if val, ok := art["1000x1000"].(string); ok {
				artwork = val
			}
		}

		title, _ := obj["title"].(string)
		id, _ := obj["id"].(string)

		// artist name
		artist := ""
		if user, ok := obj["user"].(map[string]interface{}); ok {
			if name, ok := user["name"].(string); ok {
				artist = name
			}
		}

		streamURL := "https://discoveryprovider.audius.co/v1/tracks/" + id + "/stream?app_name=lyvora"

		tracks = append(tracks, models.Track{
			ID:        id,
			Title:     title,
			Artist:    artist,
			Artwork:   artwork,
			StreamURL: streamURL,
		})
	}

	c.JSON(http.StatusOK, tracks)
}

// StreamTrack streams a track by ID (proxying from Audius or another source)
func StreamTrack(c *gin.Context) {
	id := c.Param("id")

	audioURL := fmt.Sprintf("https://discoveryprovider.audius.co/v1/tracks/%s/stream", id)

	resp, err := http.Get(audioURL)
	if err != nil {
		log.Printf("failed to fetch audio: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch audio"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("audio source returned %d: %s", resp.StatusCode, string(body))
		c.JSON(http.StatusBadGateway, gin.H{"error": "audio source returned error"})
		return
	}

	c.Header("Content-Type", "audio/mpeg")
	c.Status(http.StatusOK)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		log.Printf("streaming error: %v", err)
	}
}
