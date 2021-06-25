package main

import (
	"context"
	"encoding/json"

	"net/http"
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

// func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")

// 	params := mux.Vars(request)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])

// 	var person Person

// 	collection := client.Database("persondb").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + "}"))

// 		return
// 	}

// 	json.NewEncoder(response).Encode(person)
// }

// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")

// 	var people []Person

// 	collection := client.Database("persondb").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + "}"))

// 		return
// 	}

// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var person Person
// 		cursor.Decode(&person)
// 		people = append(people, person)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + "}"))

// 		return
// 	}

// 	json.NewEncoder(response).Encode(people)
// }
// func main() {
// 	fmt.Println("Starting the application...")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	databaseName, exists := os.LookupEnv("DATABASE_CONTAINER_NAME")

// 	if !exists {
// 		databaseName = "0.0.0.0"
// 	}
// 	mongodbURI := "mongodb://" + databaseName

// 	clientOptions := options.Client().ApplyURI(mongodbURI)

// 	client, _ = mongo.Connect(ctx, clientOptions)
// 	router := mux.NewRouter()

// 	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
// 	// router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
// 	// router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")

// 	fmt.Print("Listening on port: 12345")
// 	http.ListenAndServe(":12345", router)

// }

func main() {
	webserver.RunnerRunner()
	webserver.RunnerRunnerRunner()

	r := gin.Default()

	orders := r.Group("/order")
	{
		orders.POST("/create", func(c *gin.Context) {
			var order database.Order

			if err := c.ShouldBindJSON(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			// if order.PurchaseTime != nil {
			// set the current time
			// }

			c.JSON(http.StatusOK, order)
		})
		orders.GET("/{id}", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		orders.POST("/{id}/update", func(c *gin.Context) {
			var order database.Order

			if err := c.ShouldBindJSON(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			// if order.PurchaseTime != nil {
			// set the current time
			// }

			c.JSON(http.StatusOK, order)
		})
	}

	chefs := r.Group("/chef")
	{
		chefs.POST("/create", func(c *gin.Context) {
			var chef database.Chef

			if err := c.ShouldBindJSON(&chef); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
			}

			c.JSON(http.StatusOK, chef)
		})
		chefs.GET("/{id}", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		chefs.POST("/{id}/update", func(c *gin.Context) {
			var chef database.Chef

			if err := c.ShouldBindJSON(&chef); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
				return
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
