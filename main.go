package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

type Song struct {
	ID     primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	Name   string             `json:"name, omitempty" bson:"name, omitempty`
	Album  string             `json:"album, omitempty" bson:"album, omitempty`
	Artist string             `json:"artist, omitempty" bson:"artist, omitempty`
	Genres []string           `json:"genres, omitempty" bson:"genres, omitempty`
}

func  main()  {
	
	var mongodbURI = "mongodb://db"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
}