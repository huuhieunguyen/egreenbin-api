package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID      primitive.ObjectID `bson:"_id" json:"id" `
	Student Student            `bson:"student,inline" json:"Student"`
	Content string             `bson:"content" json:"Content"`
	// DateSort   primitive.DateTime `bson:"dateSort" json:"DateSort"`
	Type        string             `bson:"type" json:"type"`
	DateCreated primitive.DateTime `bson:"dateCreated" json:"DateCreated"`
	DateUpdated primitive.DateTime `bson:"dateUpdated" json:"DateUpdated"`

	// String idStudent; // id of student of comment
	// String content; // content of comment
	// String dateSort; // xem class dateSort, co the tra ve 3 kieu: date, week, month
	// String type;	// date, week, month
	// DateTime? dateCreate; // date create of comment
}
