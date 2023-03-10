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

// TeacherHandler  represent the httphandler for article
type TeacherHandler struct {
	DB *mongo.Database
}

// TeacherHandler will initialize the articles/ resources endpoint
func NewTeacherHandler(gin *gin.RouterGroup, appCtx component.AppContext, db *mongo.Database) {
	handler := &TeacherHandler{
		DB: db,
	}
	teachers := gin.Group("/teachers")
	{
		teachers.GET("", handler.GetTeachers)
		teachers.POST("", handler.Create)
		teachers.GET(":id", handler.GetByID)
		teachers.PUT(":id", handler.Update)
		teachers.DELETE(":id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
func (a *TeacherHandler) GetTeachers(c *gin.Context) {
	ctx := c.Request.Context()
	var teachers []models.Teacher
	cursor, err := a.DB.Collection("teachers").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &teachers); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(teachers))
}

// GetByID will get teachers by given id
func (a *TeacherHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the user with the matching ID in the "teachers" collection
	var teacher models.Teacher
	err = a.DB.Collection("teachers").FindOne(ctx, bson.M{"_id": objectID}).Decode(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, common.SimpleSuccessResponse(teacher))
}

// Create teacher will create a new teacher based on given request body
func (a *TeacherHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		return
	}

	teacher.ID = primitive.NewObjectID()
	_, err := a.DB.Collection("teachers").InsertOne(ctx, teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, common.SimpleSuccessResponse(teacher))
}

// Update will update a teacher by given id
func (a *TeacherHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	var requestBody models.Teacher

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	updates := map[string]interface{}{
		"name":           requestBody.Name,
		"Code":           requestBody.Code,
		"ImageAvatarUrl": requestBody.ImageAvatarUrl,
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update the user with the matching ID in the "teachers" collection
	_, err = a.DB.Collection("teachers").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// Delete will delete a teacher by given id
func (a *TeacherHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Delete the user with the matching ID in the "teachers" collection
	_, err = a.DB.Collection("teachers").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
