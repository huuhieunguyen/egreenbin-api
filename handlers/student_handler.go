package handlers

import (
	"context"
	"net/http"

	"github.com/GDSC-UIT/egreenbin-api/common"
	"github.com/GDSC-UIT/egreenbin-api/component"
	"github.com/GDSC-UIT/egreenbin-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StudentHandler  represent the httphandler for article
type StudentHandler struct {
	DB *mongo.Database
}

// StudentHandler will initialize the articles/ resources endpoint
func NewStudentHandler(gin *gin.RouterGroup, appCtx component.AppContext, db *mongo.Database) {
	handler := &StudentHandler{
		DB: db,
	}
	students := gin.Group("/students")
	{
		students.GET("", handler.GetStudents)
		students.POST("", handler.Create)
		students.GET(":id", handler.GetByID)
		students.PUT(":id", handler.Update)
		students.DELETE(":id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
func (a *StudentHandler) GetStudents(c *gin.Context) {
	ctx := c.Request.Context()
	var students []models.Student
	cursor, err := a.DB.Collection("students").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &students); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(students))
}

// GetByID will get students by given id
func (a *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the user with the matching ID in the "students" collection
	var student models.Student
	err = a.DB.Collection("students").FindOne(ctx, bson.M{"_id": objectID}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, common.SimpleSuccessResponse(student))
}

// Create student will create a new student based on given request body
func (a *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		return
	}

	student.ID = primitive.NewObjectID()
	_, err := a.DB.Collection("students").InsertOne(ctx, student)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, common.SimpleSuccessResponse(student))
}

// Update will update a student by given id
func (a *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	var requestBody models.Student

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	updates := map[string]interface{}{
		"name":           requestBody.Name,
		"Code":           requestBody.Code,
		"NumOfCorrect":   requestBody.NumOfCorrect,
		"NumOfWrong":     requestBody.NumOfWrong,
		"ImageAvatarUrl": requestBody.ImageAvatarUrl,
		"ParentEmail":    requestBody.ParentEmail,
		"Note":           requestBody.Note,
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update the user with the matching ID in the "students" collection
	_, err = a.DB.Collection("students").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// Delete will delete a student by given id
func (a *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Delete the user with the matching ID in the "students" collection
	_, err = a.DB.Collection("students").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
