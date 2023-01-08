package delivery

import (
	"net/http"

	"github.com/GDSC-UIT/egreenbin-api/common"
	"github.com/GDSC-UIT/egreenbin-api/component"
	"github.com/GDSC-UIT/egreenbin-api/models"
	"github.com/GDSC-UIT/egreenbin-api/modules/person/usecases"
	"github.com/gin-gonic/gin"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PersonHandler  represent the httphandler for article
type PersonHandler struct {
	PersonUsecase usecases.PersonUsecase
}

// PersonHandler will initialize the articles/ resources endpoint
func NewPersonHandler(gin *gin.RouterGroup, appCtx component.AppContext, us usecases.PersonUsecase) {
	handler := &PersonHandler{
		PersonUsecase: us,
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

	listAr, err := a.PersonUsecase.GetAll(ctx)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, common.SimpleSuccessResponse(listAr))
}

// GetByID will get persons by given id
func (a *PersonHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	person, err := a.PersonUsecase.GetByID(ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, person)
}

// Create person will create a new person based on given request body
func (a *PersonHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		return
	}

	err := a.PersonUsecase.Create(ctx, &person)
	if err != nil {
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
		// DO SOMETHING WITH THE ERROR
	}
	err := a.PersonUsecase.Update(ctx, id, map[string]interface{}{
		"name":  requestBody.Name,
		"genre": requestBody.Genre,
	})
	if err != nil {
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// Delete will delete a person by given id
func (a *PersonHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	err := a.PersonUsecase.Delete(ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}
