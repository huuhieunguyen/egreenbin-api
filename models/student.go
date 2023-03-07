package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" `
	Code           string             `bson:"code" json:"code"`
	Name           string             `bson:"name" json:"name"`
	NumOfCorrect   int                `bson:"NumOfCorrect" json:"NumOfCorrect"`
	NumOfWrong     int                `bson:"NumOfWrong" json:"NumOfWrong"`
	ImageAvatarUrl string             `bson:"ImageAvatarUrl" json:"ImageAvatarUrl"`
	ParentEmail    string             `bson:"ParentEmail" json:"ParentEmail"`
	Note           string             `bson:"Note" json:"Note"`
}
