package router

import (
	"Kaplan-Go/httpGeneric"
	"Kaplan-Go/population"

	"Kaplan-Go/config"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(configuration config.Config) *gin.Engine {
	router := gin.Default()
	route := router.Group("/")
	client := httpGeneric.NewHttpGenericClient(5, configuration)
	populationService := population.NewPopulationService(configuration, client)
	populationHandler := population.NewPopulationHandler(populationService)
	route.GET("/helloworld", populationHandler.HelloWorld)
	route.GET("/countries", populationHandler.GetCountries)
	route.GET("/countriesgo", populationHandler.GetCountriesGO)

	return router
}
