package main

import (
	"context"
	"log"
	"os"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

type Experience struct {
	ID               primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name             string               `json:"name, omitempty" bson:"name, omitempty"`
	Description      string               `json:"description, omitempty" bson:"description, omitempty"[`
	Attendees        []primitive.ObjectID `json:"attendees, omitempty" bson:"attendees, omitempty"`
	DateOfExperience primitive.DateTime   `json:"dateofexperience, omitempty" bson:"dateofexperience, omitempty"`
	Price            float32              `json:"price, omitempty bson:"price, omitempty"`
}
type User struct {
	ID          primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name        string               `json:"name, omitempty" bson:"name, omitempty"`
	Email       string               `json:"email, omitempty" bson:"email, omitempty"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}
type Chef struct {
	ID          primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name        string               `json:"name, omitempty" bson:"name, omitempty"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}

type Order struct {
	ID           primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	ExperienceId primitive.ObjectID `json:"experienceid, omitempty" bson:"experienceid, omitempty"`
	ChefId       primitive.ObjectID `json:"chefid, omitempty" bson:"chefid, omitempty"`
	PurchaseTime primitive.DateTime `json:"purchasetime, omitempty" bson:"purchasetime, omitempty"`
	SubTotal     float32            `json:"subtotal, omitempty" bson:"subtotal, omitempty"`
	Tip          float32            `json:"tip, omitempty" bson:"tip, omitempty"`
	Taxes        float32            `json:"taxes, omitempty" bson:"taxes, omitempty"`
	// Total, calculate it on the fly
}

func main() {

	// = "mongodb://0.0.0.0"
	// = "mongodb://db"

	databaseName, exists := os.LookupEnv("DATABASE_CONTAINER_NAME")
	if !exists {
		databaseName = "0.0.0.0"
	}
	mongodbURI := "mongodb://" + databaseName

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

}

// inspecting databases
// "go.mongodb.org/mongo-driver/bson"

// databases, err := client.ListDatabaseNames(ctx, bson.M{})
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(databases)

// inserting records into database

// quickstartDatabase := client.Database("quickstart")
// podcastsCollection := quickstartDatabase.Collection("podcasts")
// episodesCollection := quickstartDatabase.Collection("episodes")

// podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
// 	{Key: "title", Value: "The Polyglot Developer Podcast"},
// 	{Key: "author", Value: "Nic Raboy"},
// })

// episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
// 	bson.D{
// 		{"podcast", podcastResult.InsertedID},
// 		{"title", "GraphQL for API Development"},
// 		{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
// 		{"duration", 25},
// 	},
// 	bson.D{
// 		{"podcast", podcastResult.InsertedID},
// 		{"title", "Progressive Web Application Development"},
// 		{"description", "Learn about PWA development with Tara Manicsic."},
// 		{"duration", 32},
// 	},
// })

// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))

// Looking for records
// quickstartDatabase := client.Database("quickstart")
// podcastsCollection := quickstartDatabase.Collection("podcasts")
// episodesCollection := quickstartDatabase.Collection("episodes")

// cursor, err := podcastsCollection.Find(ctx, bson.M{})
// if err != nil {
// 	log.Fatal(err)
// }
// defer cursor.Close(ctx)

// for cursor.Next(ctx) {
// 	var podcast bson.M
// 	if err = cursor.Decode(&podcast); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(podcast)
// }

// cursor, err = episodesCollection.Find(ctx, bson.M{})

// if err != nil {
// 	log.Fatal(err)
// }
// defer cursor.Close(ctx)

// for cursor.Next(ctx) {
// 	var episode bson.M
// 	if err = cursor.Decode(&episode); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(episode)
// }
