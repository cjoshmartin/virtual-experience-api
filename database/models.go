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
	Name        string               `json:"name, omitempty" bson:"name, omitempty"`
	Email       string               `json:"email, omitempty" bson:"email, omitempty"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}
type Chef struct {
	ID          primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name        string               `json:"name, omitempty" bson:"name, omitempty"`
	Experiences []primitive.ObjectID `json:"experiences, omitempty" bson:"experiences, omitempty"`
}

type Order struct {
	ID           primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	ExperienceId primitive.ObjectID `json:"experienceid, omitempty" bson:"experienceid, omitempty"`
	ChefId       primitive.ObjectID `json:"chefid, omitempty" bson:"chefid, omitempty"`
	PurchaseTime primitive.DateTime `json:"purchasetime, omitempty" bson:"purchasetime, omitempty"`
	SubTotal     float32            `json:"subtotal, omitempty" bson:"subtotal, omitempty"`
	Tip          float32            `json:"tip, omitempty" bson:"tip, omitempty"`
	Taxes        float32            `json:"taxes, omitempty" bson:"taxes, omitempty"`
	// Total, calculate it on the fly
}
