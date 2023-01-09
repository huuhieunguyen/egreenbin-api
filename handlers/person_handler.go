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

// PersonHandler  represent the httphandler for article
type PersonHandler struct {
	DB *mongo.Database
}

// PersonHandler will initialize the articles/ resources endpoint
func NewPersonHandler(gin *gin.RouterGroup, appCtx component.AppContext, db *mongo.Database) {
	handler := &PersonHandler{
		DB: db,
	}
	persons := gin.Group("/persons")
	{
		persons.GET("", handler.GetPersons)
		persons.POST("", handler.Create)
		persons.GET(":id", handler.GetByID)
		persons.PUT(":id", handler.Update)
		persons.DELETE(":id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
func (a *PersonHandler) GetPersons(c *gin.Context) {
	ctx := c.Request.Context()
	var persons []models.Person
	cursor, err := a.DB.Collection("persons").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &persons); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(persons))
}

// GetByID will get persons by given id
func (a *PersonHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the user with the matching ID in the "persons" collection
	var person models.Person
	err = a.DB.Collection("persons").FindOne(ctx, bson.M{"_id": objectID}).Decode(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, common.SimpleSuccessResponse(person))
}

// Create person will create a new person based on given request body
func (a *PersonHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		return
	}

	person.ID = primitive.NewObjectID()
	_, err := a.DB.Collection("persons").InsertOne(ctx, person)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, common.SimpleSuccessResponse(person))
}

// Update will update a person by given id
func (a *PersonHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	var requestBody models.Person

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	updates := map[string]interface{}{
		"name":  requestBody.Name,
		"genre": requestBody.Genre,
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update the user with the matching ID in the "persons" collection
	_, err = a.DB.Collection("persons").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// Delete will delete a person by given id
func (a *PersonHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Delete the user with the matching ID in the "persons" collection
	_, err = a.DB.Collection("persons").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
