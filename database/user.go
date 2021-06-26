package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (userInstance *UserInstanceAccessor) CreateUser(user User) (*mongo.InsertOneResult, error) {
	instance := userInstance.instance
	cancel, client := userInstance.ConnectToUserCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	user.ID = primitive.NewObjectID()

	return  userInstance.collection.InsertOne(instance.ctx, user)
}

func (userInstance *UserInstanceAccessor) FindUser(hexString string) (User, error){
	id := GetID(hexString)

	var user User

	instance := userInstance.instance
	cancel, client := userInstance.ConnectToUserCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	err := userInstance.collection.FindOne(instance.ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}