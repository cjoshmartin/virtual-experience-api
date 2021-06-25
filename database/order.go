package database

import (
	// "log"
	// "fmt"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderInstanceAccessor struct{
	instance  DatabaseInstance

}


func (orderInstance OrderInstanceAccessor) GetCollection() *mongo.Collection {
	return orderInstance.instance.database.Collection("orders")
}

// func (orderInstance OrderInstanceAccessor) Create(order Order) (*mongo.InsertOneResult, error) {
// 	orders := DatabaseAccessor.GetCollection()

// 	return orders.InsertOne(ctx, order)
// }

// func FindOrder(id string) (Order, error) {
// 	orders := GetOrdersCollection()

// 	var order Order

// 	err := orders.FindOne(ctx, Order{ID: id}).Decode(&order)

// 	return order, err
// }

// func UpdateRecord(id string, data bson.D){
// 	orders := GetOrdersCollection()

// 	_id, _ := primitive.ObjectIDFromHex(id)

// 	result, err := orders.UpdateOne(ctx,
// 		bson.M{"_id": id},
// 		bson.D{
// 			{"$set", data},
// 		},
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
// }