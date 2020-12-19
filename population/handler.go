package population

import (
	"Kaplan-Go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Handler interface {
	HelloWorld(*gin.Context)
	GetCountries(*gin.Context)
	GetPopulationForCountry(*gin.Context)
}

type populationHandler struct {
	populationService Service
}

func NewPopulationHandler(populationService Service) Handler {
	return populationHandler{
		populationService: populationService,
	}
}

func (ctrl populationHandler) HelloWorld(context *gin.Context) {
	responseBody := ctrl.populationService.HelloWorld(context)
	context.JSON(http.StatusOK, responseBody)
}

func (ctrl populationHandler) GetCountries(context *gin.Context) {
	response, err := ctrl.populationService.GetCountries(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, response)

}

func (ctrl populationHandler) GetPopulationForCountry(context *gin.Context) {
	var requestBody models.PopulationRequest
	bindingErr := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if bindingErr != nil {
		log.Println("Error While Binding Request Body")
		context.AbortWithStatusJSON(http.StatusBadRequest, bindingErr)
		return
	}
	response, err := ctrl.populationService.GetPopulationForCountry(context, requestBody)
	if err != nil {
		log.Println("Error While fetching population for the given countries")
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, response)

}
