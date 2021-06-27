package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ExperienceInstanceAccessor struct {
	instance *Instance
	collection *mongo.Collection
}


func (experienceInstance *ExperienceInstanceAccessor) InitExperienceCollection() {
	collection := "experiences"
	log.Println("accessing '"+ collection + "' collection")
	experienceInstance.collection = experienceInstance.instance.database.Collection(collection)
}

func ExperienceInit(instance *Instance) *ExperienceInstanceAccessor{
	experienceCollection := ExperienceInstanceAccessor{instance: instance}
	experienceCollection.InitExperienceCollection()

	return &experienceCollection
}

func (experienceInstance *ExperienceInstanceAccessor) ConnectToExperienceCollection() (context.CancelFunc, *mongo.Client){
	instance := experienceInstance.instance
	instance.Connect()
	experienceInstance.InitExperienceCollection()
	return instance.cancel, instance.client
}

func (experienceInstance *ExperienceInstanceAccessor) CreateExperience(experience Experience) (*mongo.InsertOneResult, error){
	instance := experienceInstance.instance
	cancel, client := experienceInstance.ConnectToExperienceCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	experience.ID = primitive.NewObjectID()

	return experienceInstance.collection.InsertOne(instance.ctx, experience)
}

func (experienceInstance *ExperienceInstanceAccessor) FindExperience(hexString string)(Experience, error){
		id := GetID(hexString)

		var experience Experience

		instance := experienceInstance.instance
		cancel, client := experienceInstance.ConnectToExperienceCollection()
		defer cancel()
		defer client.Disconnect(instance.ctx)

		err := experienceInstance.collection.FindOne(instance.ctx, bson.M{"_id": id}).Decode(&experience)

		return experience, err
}