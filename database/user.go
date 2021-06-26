package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserInstanceAccessor struct {
	instance *Instance
	collection *mongo.Collection
}

func UserInit(instance *Instance) *UserInstanceAccessor{
	userCollection := UserInstanceAccessor{instance: instance}
	userCollection.InitUserCollection()

	return &userCollection
}

func (userInstance *UserInstanceAccessor) InitUserCollection() {
	collection := "users"
	log.Println("accessing '"+ collection + "' collection")
	userInstance.collection = userInstance.instance.database.Collection(collection)
}

func (userInstance *UserInstanceAccessor) ConnectToUserCollection() (context.CancelFunc, *mongo.Client) {
	instance := userInstance.instance
	instance.Connect()
	userInstance.InitUserCollection()
	return instance.cancel, instance.client
}
