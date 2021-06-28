package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Experience struct {
	ID               primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	ChefID			 primitive.ObjectID    `json:"chefid" bson:"chefid" binding:"required"`
	Name             string               `json:"name" bson:"name" binding:"required"`
	Description      string               `json:"description, omitempty" bson:"description, omitempty"`
	Attendees        []primitive.ObjectID `json:"attendees" bson:"attendees"`
	DateAndTime time.Time            `json:"dateandtime" bson:"dateandtime" binding:"required"`
	Price            float32              `json:"price" bson:"price" binding:"required"`
}
type User struct {
	ID          primitive.ObjectID   `json:"_id, omitempty" bson:"_id, omitempty"`
	Name        string               `json:"name, omitempty" bson:"name, omitempty" binding:"required"`
	Email       string               `json:"email, omitempty" bson:"email, omitempty" binding:"required"`
}
type Chef struct {
	ID          primitive.ObjectID   `bson:"_id, omitempty"`
	Name        string               `json:"name" bson:"name" binding:"required"`
	Email	string `json:"email" bson:"email" binding:"required"`
}

type Order struct {
	ID           primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	ExperienceId primitive.ObjectID `json:"experienceid, omitempty" bson:"experienceid, omitempty" binding:"required"`
	ChefId       primitive.ObjectID `json:"chefid, omitempty" bson:"chefid, omitempty" binding:"required"`
	UserID		 primitive.ObjectID`json:"userid, omitempty" bson:"userid, omitempty" binding:"required"`
	DateAndTime time.Time 			`json:"dateandtime" bson:"dateandtime"`
	SubTotal     float32            `json:"subtotal" bson:"subtotal" binding:"required"`
	Tip          float32            `json:"tip, omitempty" bson:"tip, omitempty"`
	Taxes        float32            `json:"taxes" bson:"taxes" binding:"required"`
		Total float32 `json:"total, omitempty" bson:"total, omitempty"`
}
