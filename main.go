package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/mail"
	"time"

	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/cjoshmartin/virtual-experience-api/webserver"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("persondb").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)

	json.NewEncoder(response).Encode(result)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	webserver.RunnerRunner()
	webserver.RunnerRunnerRunner()
	mongoDatabase := database.StartDatebase()
	orderCollection := database.OrderInit(mongoDatabase)

	r := gin.Default()

	orders := r.Group("/order")
	{
		orders.POST("/create", func(c *gin.Context) {
			var order database.Order

			if err := c.ShouldBindJSON(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			order.PurchaseTime = primitive.Timestamp{T:uint32(time.Now().Unix())}

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
			result, err := orderCollection.CreateOrder(order)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, result)
		})

		orders.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			order, err := orderCollection.FindOrder(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, order)
		})
		orders.POST("/{id}/update", func(c *gin.Context) {
			var order database.Order

			if err := c.ShouldBindJSON(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, order)
		})
	}

	chefCollection := database.ChefInit(mongoDatabase)

	chefs := r.Group("/chef")
	{
		chefs.POST("/create", func(c *gin.Context) {
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
		})

		chefs.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			chef, err := chefCollection.FindChef(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, chef)
		})
		chefs.POST("/:id/add-experience", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		chefs.POST("/:id/update", func(c *gin.Context) {
			var chef database.Chef

			if err := c.ShouldBindJSON(&chef); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, chef)
		})

		chefs.GET("/all", func(c *gin.Context) {

			chef, err := chefCollection.FindAllChefs()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			if chef == nil {
				chef = []database.Chef{}
			}

			c.JSON(http.StatusOK, chef)
		})
	}

	users := r.Group("/user")
	{
		users.POST("/create", func(c *gin.Context) {
			var user database.User

			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, user)
		})
		users.GET("/{id}", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		users.POST("/{id}/update", func(c *gin.Context) {
			var user database.User

			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, user)
		})
	}

	experiences := r.Group("/experience")
	{
		experiences.POST("/create", func(c *gin.Context) {
			var experience database.Experience

			if err := c.ShouldBindJSON(&experience); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, experience)
		})
		experiences.GET("/{id}", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		experiences.POST("/{id}/update", func(c *gin.Context) { // update record
			var experience database.Experience

			if err := c.ShouldBindJSON(&experience); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, experience)
		})
	}

	r.Run()
}
