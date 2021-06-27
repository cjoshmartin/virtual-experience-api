package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func CreateExperience(chefCollection *database.ChefInstanceAccessor, experienceCollection *database.ExperienceInstanceAccessor) gin.HandlerFunc {
	return func (c *gin.Context) {
		var experience database.Experience

		if err := c.ShouldBindJSON(&experience); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		chefId := experience.ChefID
		_, err := chefCollection.FindChef(chefId.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Problem found with the chefid you have provided. Please check it and send again"})
			return
		}

		experience.Attendees = []primitive.ObjectID{}
		result, err := experienceCollection.CreateExperience(experience)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetExperienceByID(experienceCollection *database.ExperienceInstanceAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		experience, err := experienceCollection.FindExperience(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, experience)
	}
}

func UpdateExperience(experienceCollection *database.ExperienceInstanceAccessor) gin.HandlerFunc { // TODO: make this work
	return func(c *gin.Context) { // update record
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		var experience database.Experience

		if err := c.ShouldBindJSON(&experience); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, experience)
	}
}