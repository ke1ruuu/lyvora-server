package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/ke1ruuu/lyvora-server/internal/db"
	"github.com/ke1ruuu/lyvora-server/internal/models"
	"github.com/ke1ruuu/lyvora-server/internal/utils"
)

// Middleware to validate JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		token := authHeader[len("Bearer "):]
		username, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}

// Add a favorite
func AddFavorite(c *gin.Context) {
	username := c.GetString("username")

	var track models.Track
	if err := c.BindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	key := []byte("favorites:" + username)

	var favorites []models.Track
	db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return nil // No favorites yet
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &favorites)
		})
	})

	// Add if not already in list
	for _, f := range favorites {
		if f.ID == track.ID {
			c.JSON(http.StatusOK, gin.H{"message": "already in favorites"})
			return
		}
	}

	favorites = append(favorites, track)
	data, _ := json.Marshal(favorites)

	db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})

	c.JSON(http.StatusOK, gin.H{"message": "track added to favorites"})
}

// Get favorites
func GetFavorites(c *gin.Context) {
	username := c.GetString("username")
	key := []byte("favorites:" + username)

	var favorites []models.Track
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &favorites)
		})
	})

	if err != nil {
		c.JSON(http.StatusOK, []models.Track{})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// Remove favorite
func RemoveFavorite(c *gin.Context) {
	username := c.GetString("username")
	id := c.Param("id")
	key := []byte("favorites:" + username)

	var favorites []models.Track
	db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return nil
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &favorites)
		})
	})

	// Filter out
	newFavs := []models.Track{}
	for _, f := range favorites {
		if f.ID != id {
			newFavs = append(newFavs, f)
		}
	}

	data, _ := json.Marshal(newFavs)
	db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})

	c.JSON(http.StatusOK, gin.H{"message": "removed from favorites"})
}
