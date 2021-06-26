package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)
type ChefInstanceAccessor struct {
	instance   *Instance
	collection *mongo.Collection
}

func ChefInit(instance *Instance) *ChefInstanceAccessor {
	chefCollection := ChefInstanceAccessor{instance: instance}
	chefCollection.InitChefCollection()

	return &chefCollection
}

func (chefInstance *ChefInstanceAccessor) InitChefCollection() {
	collection := "chefs"
	log.Println("accessing '"+ collection + "' collection")
	chefInstance.collection = chefInstance.instance.database.Collection(collection)
}

func (chefInstance *ChefInstanceAccessor) ConnectToChefsCollection() (context.CancelFunc, *mongo.Client){
	instance := chefInstance.instance
	instance.Connect()
	chefInstance.InitChefCollection()
	return instance.cancel, instance.client
}

func (chefInstance *ChefInstanceAccessor) CreateChef(chef Chef) (*mongo.InsertOneResult, error) {
	instance := chefInstance.instance
	cancel, client := chefInstance.ConnectToChefsCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	chef.ID = primitive.NewObjectID()

	return chefInstance.collection.InsertOne(instance.ctx, chef)
}

func (chefInstance *ChefInstanceAccessor) FindChef(id string) (Chef, error) {

	var chef Chef

	instance := chefInstance.instance
	cancel, client := chefInstance.ConnectToChefsCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	// not sure why this is not finding records
	err := chefInstance.collection.FindOne(instance.ctx, bson.M{"_id": id}).Decode(&chef)

	return chef, err
}

func (chefInstance *ChefInstanceAccessor) FindAllChefs() ([]Chef, error){

	var chefs []Chef

	instance := chefInstance.instance
	cancel, client := chefInstance.ConnectToChefsCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	cursor, err := chefInstance.collection.Find(instance.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(instance.ctx)
	for cursor.Next(instance.ctx){
		var chef Chef
		if err = cursor.Decode(&chef); err != nil {
			return chefs, err
		}
			chefs = append(chefs, chef)
	}

	return chefs, err
}