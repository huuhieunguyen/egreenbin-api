package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Garbage struct {
	ID primitive.ObjectID `bson:"_id" json:"id" `
	// StudentID primitive.ObjectID `bson:"studentID" json:"studentID"`
	StudentID string             `bson:"studentID" json:"studentID"`
	Name      string             `bson:"name" json:"name"`
	DateThrow primitive.DateTime `bson:"dateThrow" json:"dateThrow"`
	IsRight   bool               `bson:"isRight" json:"isRight"`
}
