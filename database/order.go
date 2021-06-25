package database

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderInstanceAccessor struct {
	instance   *DatabaseInstance
	collection *mongo.Collection
}

func OrderInit(instance *DatabaseInstance) *OrderInstanceAccessor {
	orderCollection := OrderInstanceAccessor{instance: instance}
	orderCollection.InitCollection()

	return &orderCollection
}

func (orderInstance *OrderInstanceAccessor) InitCollection() {
	orderInstance.collection = orderInstance.instance.database.Collection("orders")
}

func (orderInstance *OrderInstanceAccessor) Create(order Order) (*mongo.InsertOneResult, error) {
	return orderInstance.collection.InsertOne(orderInstance.instance.ctx, order)
}

func (orderInstance *OrderInstanceAccessor) FindOrder(hexString string) (Order, error) {
	id := GetID(hexString)

	var order Order

	err := orderInstance.collection.FindOne(orderInstance.instance.ctx, bson.M{"_id": id}).Decode(&order)

	return order, err
}

func (orderInstance *OrderInstanceAccessor) UpdateRecord(hexString string, data bson.D) (*mongo.UpdateResult, error) {
	id := GetID(hexString)

	result, err := orderInstance.collection.UpdateOne(orderInstance.instance.ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", data},
		},
	)

	log.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return result, err
}
