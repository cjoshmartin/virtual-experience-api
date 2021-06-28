package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(userCollection *database.UserInstanceAccessor) gin.HandlerFunc{
	return func(c *gin.Context) {
		var user database.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		if !isValidEmail(user.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid email address"})
			return
		}

		result, err := userCollection.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetUserByID(userCollection *database.UserInstanceAccessor) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		user, err := userCollection.FindUser(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}}

// TODO Make this work
func UpdateUser(userCollection *database.UserInstanceAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_ = id
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		var user database.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}