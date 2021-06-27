package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func CreateChef(chefCollection *database.ChefInstanceAccessor) gin.HandlerFunc {

	return  func(c *gin.Context) {
		var chef database.Chef

		if err := c.ShouldBindJSON(&chef); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		if !isValidEmail(chef.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid email address"})
			return
		}

		if chef.Experiences == nil {
			chef.Experiences = []primitive.ObjectID{}
		}

		result, err := chefCollection.CreateChef(chef)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}


		c.JSON(http.StatusOK, result)
	}
}

func GetChefByID(chefCollection *database.ChefInstanceAccessor) gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		chef, err := chefCollection.FindChef(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, chef)
	}
}
// TODO
func AddExperienceToAChef(chefCollection *database.ChefInstanceAccessor) gin.HandlerFunc{
	return  func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}

func UpdateChef(chefCollection *database.ChefInstanceAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		var chef database.Chef

		if err := c.ShouldBindJSON(&chef); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, chef)
	}
}
func GetAllChefs(chefCollection *database.ChefInstanceAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		chef, err := chefCollection.FindAllChefs()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		if chef == nil {
			chef = []database.Chef{}
		}

		c.JSON(http.StatusOK, chef)
	}
}