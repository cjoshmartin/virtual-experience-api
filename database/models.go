package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID               primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	ChefID			 primitive.ObjectID    `json:"chefid" bson:"chefid" binding:"required"`
	Name             string               `json:"name" bson:"name" binding:"required"`
	Description      string               `json:"description, omitempty" bson:"description, omitempty"`
	Attendees        []primitive.ObjectID `json:"attendees" bson:"attendees"`
	DateOfExperience primitive.DateTime   `json:"dateofexperience, omitempty" bson:"dateofexperience, omitempty"`
	Price            float32              `json:"price" bson:"price" binding:"required"`
}
type User struct {
	ID          primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name        string               `json:"name, omitempty" bson:"name, omitempty" binding:"required"`
	Email       string               `json:"email, omitempty" bson:"email, omitempty" binding:"required"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}
type Chef struct {
	ID          primitive.ObjectID   `bson:"_id, omitempty"`
	Name        string               `json:"name" bson:"name" binding:"required"`
	Email	string `json:"email" bson:"email" binding:"required"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}

type Order struct {
	ID           primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	ExperienceId primitive.ObjectID `json:"experienceid, omitempty" bson:"experienceid, omitempty" binding:"required"`
	ChefId       primitive.ObjectID `json:"chefid, omitempty" bson:"chefid, omitempty" binding:"required"`
	PurchaseTime primitive.Timestamp `json:"purchasetime" bson:"purchasetime"`
	SubTotal     float32            `json:"subtotal" bson:"subtotal" binding:"required"`
	Tip          float32            `json:"tip, omitempty" bson:"tip, omitempty"`
	Taxes        float32            `json:"taxes" bson:"taxes" binding:"required"`
		Total float32 `json:"total, omitempty" bson:"total, omitempty"`
}
