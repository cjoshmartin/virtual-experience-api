package database

// inspecting databases
// "go.mongodb.org/mongo-driver/bson"

// databases, err := client.ListDatabaseNames(ctx, bson.M{})
// if err != nil {
// 	log.Fatal(err)
// }

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
// log.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))

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
// 	log.Println(podcast)
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
// 	log.Println(episode)
// }

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

type DatabaseAccessor interface {
	GetCollection() *mongo.Collection
	Create(order Order) (*mongo.InsertOneResult, error)
	Find(id string) (Order, error)
	Update(id string, data bson.D)
}

type DatabaseInstance struct {
	database   *mongo.Database
	collection *mongo.Collection
	ctx        context.Context
	client     *mongo.Client
}

func GetDatabaseURI() string {
	databaseName, exists := os.LookupEnv("DATABASE_CONTAINER_NAME")

	if !exists {
		databaseName = "0.0.0.0"
	}

	return "mongodb://" + databaseName

}

func (databaseInstance *DatabaseInstance) Connect() {
	mongodbURI := GetDatabaseURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databaseInstance.client = client
	databaseInstance.ctx = ctx
}

func (databaseInstance *DatabaseInstance) CloseConnection() {
	defer client.Disconnect(databaseInstance.ctx)
}

func (databaseInstance *DatabaseInstance) SetDatabase() {
	databaseName, exists := os.LookupEnv("MONGO_DB_NAME")

	if !exists {
		log.Fatalln("No database name found in environment.Please set `MONGO_DB_NAME`")
	} else {

		log.Println("Found  wanted database name in enviroment '" + databaseName + "'")
	}

	databaseInstance.database = databaseInstance.client.Database(databaseName)
}

func (databaseInstance *DatabaseInstance) ListDatabases() {
	databases, err := databaseInstance.client.ListDatabaseNames(databaseInstance.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Databases available for use: '" + strings.Join(databases, ",")+"'")
}

func StartDatebase() *DatabaseInstance {
	databaseInstance := DatabaseInstance{}

	log.Println("Starting database...")
	databaseInstance.Connect()
	databaseInstance.ListDatabases()
	databaseInstance.SetDatabase()
	log.Println("Database started successfully")

	return &databaseInstance
}
