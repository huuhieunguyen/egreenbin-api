package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Garbage struct {
	ID        primitive.ObjectID `bson:"_id" json:"id" `
	StudentID primitive.ObjectID `bson:"studentID" json:"studentID"`
	Name      string             `bson:"name" json:"name"`
	DateThrow primitive.DateTime `bson:"dateThrow" json:"dateThrow"`
	IsRight   bool               `bson:"isRight" json:"isRight"`
}

// String id; // id of trash
// String StudentID: // gửi lên server của Thịnh để lấy về ID
// String StudentName: // gửi lên server của Thịnh để lấy về Name
// DateTime dateThrow: // ngày vứt rác
// Bool isRight: // vứt rác đúng hay sai (true: right, false: Wrong)
