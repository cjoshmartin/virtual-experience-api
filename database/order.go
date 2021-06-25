package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderInstanceAccessor struct {
	instance   *Instance
	collection *mongo.Collection
}

func OrderInit(instance *Instance) *OrderInstanceAccessor {
	orderCollection := OrderInstanceAccessor{instance: instance}
	orderCollection.InitCollection()

	return &orderCollection
}

func (orderInstance *OrderInstanceAccessor) InitCollection() {
	collection := "orders"
	log.Println("accessing '"+ collection + "' collection")
	orderInstance.collection = orderInstance.instance.database.Collection(collection)
}

func (orderInstance *OrderInstanceAccessor) ConnectToOrdersCollection() (context.CancelFunc, *mongo.Client){
	instance := orderInstance.instance
	instance.Connect()
	orderInstance.InitCollection()
	return instance.cancel, instance.client
}

func (orderInstance *OrderInstanceAccessor) Create(order Order) (*mongo.InsertOneResult, error) {
	instance := orderInstance.instance
	cancel, client := orderInstance.ConnectToOrdersCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	return orderInstance.collection.InsertOne(instance.ctx, order)
}

func (orderInstance *OrderInstanceAccessor) FindOrder(hexString string) (Order, error) {
	id := GetID(hexString)

	var order Order

	instance := orderInstance.instance
	cancel, client := orderInstance.ConnectToOrdersCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	err := orderInstance.collection.FindOne(instance.ctx, bson.M{"_id": id}).Decode(&order)

	return order, err
}

func (orderInstance *OrderInstanceAccessor) UpdateRecord(hexString string, data bson.D) (*mongo.UpdateResult, error) {
	id := GetID(hexString)

	instance := orderInstance.instance
	cancel, client := orderInstance.ConnectToOrdersCollection()
	defer cancel()
	defer client.Disconnect(instance.ctx)

	result, err := orderInstance.collection.UpdateOne(orderInstance.instance.ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", data},
		},
	)

	log.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return result, err
}
