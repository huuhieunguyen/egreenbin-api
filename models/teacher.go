package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" `
	Code           string             `bson:"code" json:"code"`
	Name           string             `bson:"name" json:"name"`
	ImageAvatarUrl string             `bson:"imageAvatarUrl" json:"imageAvatarUrl"`
}
