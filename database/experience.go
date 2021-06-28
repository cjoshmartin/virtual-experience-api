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
		id, err := GetID(hexString)
		if err != nil {
			return Experience{}, err
		}

		var experience Experience

		instance := experienceInstance.instance
		cancel, client := experienceInstance.ConnectToExperienceCollection()
		defer cancel()
		defer client.Disconnect(instance.ctx)

		err = experienceInstance.collection.FindOne(instance.ctx, bson.M{"_id": id}).Decode(&experience)

		return experience, err
}

func (experienceInstance *ExperienceInstanceAccessor) AddAttendee(userHexString string, experienceHexString string) (*mongo.UpdateResult, error)  {
	userID, err := GetID(userHexString)
	if err != nil {
		return nil, err
	}
	experienceID, err := GetID(experienceHexString)
	if err != nil {
		return nil, err
	}
	instance := experienceInstance.instance
	cancel, client := experienceInstance.ConnectToExperienceCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	result, err := experienceInstance.collection.UpdateOne(instance.ctx, bson.M{"_id": experienceID}, bson.M{ "$push": bson.M{"attendees": userID}})
	return result, err

}

func (experienceInstance *ExperienceInstanceAccessor) GetExperienceByChefID(chefHexString string) ([]Experience, error){
	chefID, err := GetID(chefHexString)
	if err != nil {
		return []Experience{}, err
	}
	var experiences []Experience

	instance := experienceInstance.instance
	cancel, client := experienceInstance.ConnectToExperienceCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	cursor, err := experienceInstance.collection.Find(instance.ctx, bson.M{"chefid": chefID})

	defer cursor.Close(instance.ctx)
	for cursor.Next(instance.ctx) {
		var experience Experience
		if err = cursor.Decode(&experience); err != nil {
			return experiences,err
		}
		experiences = append(experiences, experience)
	}

	return experiences, err
}

func (experienceInstance *ExperienceInstanceAccessor) FindAllExperiences() ([]Experience, error) {
	var experiences []Experience

	instance := experienceInstance.instance
	cancel, client := experienceInstance.ConnectToExperienceCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	cursor, err := experienceInstance.collection.Find(instance.ctx, bson.M{})

	defer cursor.Close(instance.ctx)
	for cursor.Next(instance.ctx) {
		var experience Experience
		if err = cursor.Decode(&experience); err != nil {
			return experiences,err
		}
		experiences = append(experiences, experience)
	}

	return experiences, err
}