package api

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgraph-io/badger/v4"
	"github.com/ke1ruuu/lyvora-server/internal/db"
	"github.com/ke1ruuu/lyvora-server/internal/models"
	"github.com/ke1ruuu/lyvora-server/internal/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}
	user.Password = hashed

	data, _ := json.Marshal(user)
	err = db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("user:"+user.Username), data)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

func Login(c *gin.Context) {
	var creds models.User
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user models.User
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("user:" + creds.Username))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &user)
		})
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if !utils.CheckPassword(user.Password, creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, _ := utils.GenerateToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
