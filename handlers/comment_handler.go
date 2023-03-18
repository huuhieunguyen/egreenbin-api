package handlers

import (
	"context"
	"encoding/json"
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

// CommentHandler  represent the httphandler for article
type CommentHandler struct {
	DB *mongo.Database
}

// CommentHandler will initialize the articles/ resources endpoint
func NewCommentHandler(gin *gin.RouterGroup, appCtx component.AppContext, db *mongo.Database) {
	handler := &CommentHandler{
		DB: db,
	}
	comments := gin.Group("/comments")
	{
		comments.GET("", handler.GetComments)
		comments.POST("", handler.Create)
		comments.GET(":id", handler.GetByID)
		// comments.GET("", handler.getCommentsByStudentID)
		comments.PUT(":id", handler.Update)
		comments.DELETE(":id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
// func (a *CommentHandler) GetComments(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	var comments []models.Comment
// 	cursor, err := a.DB.Collection("comments").Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = cursor.All(ctx, &comments); err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

//		c.JSON(http.StatusOK, common.SimpleSuccessResponse(comments))
//	}
func (a *CommentHandler) GetComments(c *gin.Context) {
	ctx := c.Request.Context()
	var comments []models.Comment
	cursor, err := a.DB.Collection("comments").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(comments))
}

// GetByID will get comment by given id
func (a *CommentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the comment with the matching ID in the "comments" collection
	var comment models.Comment
	err = a.DB.Collection("comments").FindOne(ctx, bson.M{"_id": objectID}).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, common.SimpleSuccessResponse(comment))
}

// Query Comment By StudentID will get comment by given student id
// func (a *CommentHandler) GetByStuID(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	// stu_id := c.Param("_studentID")
// 	stu_id := c.Query("_studentID")

// 	objectID, err := primitive.ObjectIDFromHex(stu_id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, nil)
// 		return
// 	}
// 	// Find the user with the matching ID in the "comments" collection
// 	var studentID models.Comment
// 	err = a.DB.Collection("comments").FindOne(ctx, bson.M{"_id": objectID}).Decode(&studentID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, nil)
// 		return
// 	}
// 	c.JSON(http.StatusOK, common.SimpleSuccessResponse(studentID))
// }

// func (a *CommentHandler) getCommentsByStudentID(c *gin.Context) {
// 	studentID := c.Query("studentID")

// 	// Query the comments collection in MongoDB to find all comments by given studentID
// 	comments, err := a.DB.Collection("comments").Find(context.TODO(), bson.M{"_studentID": studentID})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Could not retrieve comments",
// 		})
// 		return
// 	}

// 	var commentList []models.Comment
// 	defer comments.Close(context.TODO())

// 	// Iterate through comments and add to list
// 	for comments.Next(context.TODO()) {
// 		var comment models.Comment
// 		err := comments.Decode(&comment)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Could not decode comments",
// 			})
// 			return
// 		}
// 		commentList = append(commentList, comment)
// 	}
// 	c.JSON(http.StatusOK, commentList)
// }

// Create comment will create a new comment based on given request body
// func (a *CommentHandler) Create(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		var comment models.Comment
// 		if err := c.ShouldBindJSON(&comment); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
// 				return
// 		}

// 		comment.ID = primitive.NewObjectID()
// 		now := time.Now()
// 		comment.DateCreated = now
// 		comment.DateUpdated = now

// 		_, err := a.DB.Collection("comments").InsertOne(ctx, comment)
// 		if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
// 				return
// 		}

// 		c.JSON(http.StatusCreated, gin.H{"data": comment})
// }

// Create comment will create a new comment based on given request body
type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (a *CommentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var comment models.Comment
	// var student models.Student
	if err := json.NewDecoder(c.Request.Body).Decode(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ID = primitive.NewObjectID()
	// student.ID = primitive.NewObjectID()

	now := primitive.NewDateTimeFromTime(time.Now())
	comment.DateCreated = now
	// comment.DateUpdated = now

	if _, err := a.DB.Collection("comments").InsertOne(ctx, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		Status:  "success",
		Data:    comment,
		Message: "Comment has been created.",
	}
	c.JSON(http.StatusCreated, res)
}

// Update will update a comment by given id
func (a *CommentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	var requestBody models.Comment

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	updates := map[string]interface{}{
		"content": requestBody.Content,
		// "DateSort":   requestBody.DateSort,
		"type": requestBody.Type,
		// "dateCreated": requestBody.DateCreated,
		// "dateUpdated": primitive.NewDateTimeFromTime(time.Now()),
		"dataUpdated": requestBody.DateUpdated,
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update the document with the matching ID in the "comments" collection
	_, err = a.DB.Collection("comments").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// Delete will delete a comment by given id
func (a *CommentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Delete the user with the matching ID in the "comments" collection
	_, err = a.DB.Collection("comments").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
