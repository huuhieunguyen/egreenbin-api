package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/GDSC-UIT/egreenbin-api/common"
	"github.com/GDSC-UIT/egreenbin-api/component"
	"github.com/GDSC-UIT/egreenbin-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GarbageHandler  represent the httphandler for article
type GarbageHandler struct {
	DB *mongo.Database
}

// GarbageHandler will initialize the articles/ resources endpoint
func NewGarbageHandler(gin *gin.RouterGroup, appCtx component.AppContext, db *mongo.Database) {
	handler := &GarbageHandler{
		DB: db,
	}
	garbage := gin.Group("/garbage")
	{
		garbage.POST("", handler.Create)
		garbage.GET("", handler.GetGarbageThrow)
		garbage.GET(":id", handler.GetByID)
		garbage.DELETE(":id", handler.DeleteByID)
	}
}

// func (a *GarbageHandler) UpdateStudentRightWrong(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	studentID, err := primitive.ObjectIDFromHex(c.Param("student_id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if student exists
// 	var student models.Student
// 	err = a.DB.Collection("students").FindOne(context.Background(), bson.M{"_id": studentID}).Decode(&student)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
// 		return
// 	}

// 	// Parse the request body as a Garbage struct
// 	var note models.Garbage
// 	if err := c.BindJSON(&note); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Update the student's NumOfCorrect or NumOfWrong field based on the IsRight flag in the Garbage
// 	var update bson.M
// 	if note.IsRight {
// 		update = bson.M{"$inc": bson.M{"num_of_correct": 1}}
// 	} else {
// 		update = bson.M{"$inc": bson.M{"num_of_wrong": 1}}
// 	}
// 	_, err = a.DB.Collection("students").UpdateOne(ctx, bson.M{"_id": studentID}, update)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Student added to course"})
// }

// Create garbage will create a new garbage based on given request body

type ResponseGarbageThrow struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (a *GarbageHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse the request body as a Garbage struct
	var garbage models.Garbage
	if err := c.BindJSON(&garbage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	garbage.ID = primitive.NewObjectID()
	garbage.DateThrow = primitive.NewDateTimeFromTime(time.Now())

	student_id := garbage.StudentID
	// studentID, err := primitive.ObjectIDFromHex(c.Param("student_id"))
	studentID, err := primitive.ObjectIDFromHex(student_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorHaaHaa": err.Error()})
		return
	}

	// Check if student exists
	var student models.Student
	// studentID := garbage.StudentID
	err = a.DB.Collection("students").FindOne(context.Background(), bson.M{"_id": studentID}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	if _, err := a.DB.Collection("garbage").InsertOne(ctx, garbage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		Status:  "success",
		Data:    garbage,
		Message: "Garbage throwing has been created.",
	}
	c.JSON(http.StatusCreated, res)

	////////////////////////////////////////////////////////////////////////////////////

	// Update the student's NumOfCorrect or NumOfWrong field based on the IsRight flag in the Garbage
	var update bson.M
	if garbage.IsRight {
		update = bson.M{"$inc": bson.M{"numOfCorrect": 1}}
	} else {
		update = bson.M{"$inc": bson.M{"numOfWrong": 1}}
	}
	_, err = a.DB.Collection("students").UpdateOne(ctx, bson.M{"_id": studentID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

// FetchArticle will fetch the article based on given params
func (a *GarbageHandler) GetGarbageThrow(c *gin.Context) {
	ctx := c.Request.Context()
	var garbage []models.Garbage
	cursor, err := a.DB.Collection("garbage").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &garbage); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(garbage))
}

// GetByID will get garbage by given id
func (a *GarbageHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the user with the matching ID in the "garbage" collection
	var garbage models.Garbage
	err = a.DB.Collection("garbage").FindOne(ctx, bson.M{"_id": objectID}).Decode(&garbage)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, common.SimpleSuccessResponse(garbage))
}

// Delete will delete a student by given id
func (a *GarbageHandler) DeleteByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Delete the collection with the matching ID in the "garbage" collection
	_, err = a.DB.Collection("garbage").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
