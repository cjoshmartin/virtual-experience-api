package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID               primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name             string               `json:"name, omitempty" bson:"name, omitempty"`
	Description      string               `json:"description, omitempty" bson:"description, omitempty"[`
	Attendees        []primitive.ObjectID `json:"attendees, omitempty" bson:"attendees, omitempty"`
	DateOfExperience primitive.DateTime   `json:"dateofexperience, omitempty" bson:"dateofexperience, omitempty"`
	Price            float32              `json:"price, omitempty bson:"price, omitempty"`
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
	ExperienceId primitive.ObjectID `json:"experienceid, omitempty" bson:"experienceid, omitempty"`
	ChefId       primitive.ObjectID `json:"chefid, omitempty" bson:"chefid, omitempty"`
	PurchaseTime primitive.Timestamp `json:"purchasetime, omitempty" bson:"purchasetime, omitempty"`
	SubTotal     float32            `json:"subtotal" bson:"subtotal" binding:"required"`
	Tip          float32            `json:"tip, omitempty" bson:"tip, omitempty"`
	Taxes        float32            `json:"taxes" bson:"taxes" binding:"required"`
		Total float32 `json:"total, omitempty" bson:"total, omitempty"`
}
