package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id" `
	StudentID primitive.ObjectID `bson:"_studentID" json:"studentID"`
	Content   string             `bson:"content" json:"content"`
	// DateSort   primitive.DateTime `bson:"dateSort" json:"DateSort"`
	Type        string             `bson:"type" json:"type"`
	DateCreated primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	// DateUpdated primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	DateUpdated string `bson:"dateUpdated" json:"dateUpdated"`
}
