package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type OrderResult struct {
	ID primitive.ObjectID `json:"id"`
	Date time.Time `json:"date"`
	Experience database.Experience `json:"experience"`
	Chef database.Chef `json:"chef"`
	HeadCount int `json:"head_count"`
	SubTotal float32 `json:"sub_total"`
	Taxes float32 `json:"taxes"`
	Tip float32 `json:"tip"`
	Total float32 `json:"total"`
}

func CreateOrder(orderCollection *database.OrderInstanceAccessor, experienceCollection *database.ExperienceInstanceAccessor, chefCollection *database.ChefInstanceAccessor, userCollection *database.UserInstanceAccessor) gin.HandlerFunc{
	return  func(c *gin.Context) {
		var order database.Order

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		order.DateAndTime = time.Now()
		taxes := order.Taxes

		if taxes > 1  || taxes < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "taxes amount should be between 0 and 1"})
			return
		}

		subTotal := order.SubTotal
		if subTotal < 0.01 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Subtotal has to be greater then 0"})
			return
		}

		taxes = order.SubTotal * order.Taxes
		total := order.SubTotal + taxes + order.Tip

		order.Total = total

		experienceId := order.ExperienceId
		experience, err := experienceCollection.FindExperience(experienceId.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid experienceId"})
			return
		}

		user, err := userCollection.FindUser(order.UserID.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Invalid User ID"})
			return
		}

		_, err = experienceCollection.AddAttendee(user.ID.Hex(), experience.ID.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		chefId := order.ChefId
		chef, err := chefCollection.FindChef(chefId.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Problem found with the chefid you have provided. Please check it and send again"})
			return
		}

		order.ID = primitive.NewObjectID()

		_, err = orderCollection.CreateOrder(order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		output := OrderResult{
			ID: order.ID,
			Date:order.DateAndTime,
			Experience : experience,
			Chef: chef,
			HeadCount: len(experience.Attendees),
			SubTotal: order.SubTotal,
			Tip: order.Tip,
			Taxes: taxes,
			Total: total,
		}

		c.JSON(http.StatusOK, output)
	}
}

func GetOrderByID(orderCollection *database.OrderInstanceAccessor) gin.HandlerFunc  {

	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "No Id provided"})
			return
		}

		order, err := orderCollection.FindOrder(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}