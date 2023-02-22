package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" `
	Name           string             `bson:"name" json:"name"`
	Genre          int                `bson:"genre" json:"genre"`
	NumOfCorrect   int                `bson:"NumOfCorrect" json:"NumOfCorrect"`
	NumOfWrong     int                `bson:"NumOfWrong" json:"NumOfWrong"`
	ImageAvatarUrl string             `bson:"ImageAvatarUrl" json:"ImageAvatarUrl"`
	ParentEmail    string             `bson:"ParentEmail" json:"ParentEmail"`
	Note           string             `bson:"Note" json:"Note"`
	//int? numOfCorrect; // number of correct
	// int? numOfWrong; // number of wrong
	// String? imageAvatarUrl; // link url of image avatar student
	//String? parentEmail; // email of parent's student
	//String? note; // note when send email to parent's student
	// defalut value is null or ""
}
