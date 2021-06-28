package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)
func IsBetween9and5(StartTime int) bool {
	NineAm := 9
	FivePM := 17
	EndTime := StartTime + 1

	DoesStartAfterNineAM := StartTime >= NineAm
	DoesStartBeforeFivePM := StartTime < FivePM
	DoesEndTimeEndBeforeFivePM := EndTime  < FivePM

	return  DoesStartAfterNineAM &&  DoesStartBeforeFivePM && DoesEndTimeEndBeforeFivePM
}

	func CreateExperience(chefCollection *database.ChefInstanceAccessor, experienceCollection *database.ExperienceInstanceAccessor) gin.HandlerFunc {
	return func (c *gin.Context) {
		var experience database.Experience

		if err := c.ShouldBindJSON(&experience); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}


		today := time.Now()
		dateOfExperience := experience.DateAndTime
		if today.Equal(dateOfExperience) || today.After(dateOfExperience) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Experiences have to be scheduled in advance"})
			return
		}

		if !IsBetween9and5(dateOfExperience.Hour()){
			c.JSON(http.StatusBadRequest, gin.H{"status": "Events must be scheduled after 9am, and before 5pm GMT. Event also have to end before 5pm GMT"})
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

type AddAttendeeBody struct {
	UserID primitive.ObjectID `json:"user_id"`
	ExperienceID primitive.ObjectID `json:"experience_id"`
}

func AddAttendeeToExperience(userCollection *database.UserInstanceAccessor, experienceCollection *database.ExperienceInstanceAccessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var addAttendeeBody AddAttendeeBody

		if err := c.ShouldBindJSON(&addAttendeeBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		_, err := userCollection.FindUser(addAttendeeBody.UserID.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid User ID"})
			return
		}
		result, err := experienceCollection.AddAttendee(addAttendeeBody.UserID.Hex(), addAttendeeBody.ExperienceID.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}

}
