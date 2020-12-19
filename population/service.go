package population

import (
	"Kaplan-Go/config"
	"Kaplan-Go/httpGeneric"
	"Kaplan-Go/models"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	HelloWorld(*gin.Context) string
	GetCountries(*gin.Context) (map[string]interface{}, error)
	GetPopulationForCountry(*gin.Context, models.PopulationRequest) (models.PopulationResponse, error)
	MakeRequest(url, country string, ch chan<- interface{}, context *gin.Context)
}

type populationService struct {
	config     config.Config
	httpClient httpGeneric.Client
}

func NewPopulationService(config config.Config, client httpGeneric.Client) Service {
	return populationService{
		config:     config,
		httpClient: client,
	}
}

// HelloWorld Sample API Which Returns Hello World
func (service populationService) HelloWorld(context *gin.Context) string {
	return "hello world"
}

// GetCountries to get the list of countries
func (service populationService) GetCountries(context *gin.Context) (result map[string]interface{}, err error) {
	responsebytes, _ := service.httpClient.Get(context, service.config.GetCountriesEndPoint())
	if err != nil {
		log.Println("error while making get call to third party", err)
		return
	}
	err = json.Unmarshal(responsebytes, &result)
	if err != nil {
		log.Println("Error while unmarshalling the data", err)
		return
	}
	return
}

// GetPopulationForCountry to get population for the given country
func (service populationService) GetPopulationForCountry(context *gin.Context, requestBody models.PopulationRequest) (resp models.PopulationResponse, err error) {
	ch := make(chan interface{})
	countries := requestBody.Countries
	clientPopulation := make(map[int]interface{})
	errorMap := make([]map[string]interface{}, 0)
	for _, country := range countries {
		go service.MakeRequest(service.config.GetPopulationEndpoint()+country+"/2020-12-18/", country, ch, context)
	}
	for range countries {
		response := <-ch
		switch response.(type) {
		case map[string]interface{}:
			res := response.(map[string]interface{})
			errorMap = append(errorMap, map[string]interface{}{"country": res["country"], "ErrorMessage": res["ErrorMessage"]})
		case map[int]interface{}:
			for k, v := range response.(map[int]interface{}) {
				clientPopulation[k] = v
			}
		}
	}

	keys := make([]int, 0, len(clientPopulation))
	for k := range clientPopulation {
		keys = append(keys, k)
	}
	if strings.TrimSpace(strings.ToLower(requestBody.Sort)) == "ascending" {
		sort.Ints(keys)
		for _, k := range keys {
			resp.Population = append(resp.Population, map[string]interface{}{"country": clientPopulation[k], "population": k})
		}
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		for _, k := range keys {
			resp.Population = append(resp.Population, map[string]interface{}{"country": clientPopulation[k], "population": k})
		}
	}
	resp.Errors = append(resp.Errors, errorMap...)
	return
}

// MakeRequest make get call for the given URL
func (service populationService) MakeRequest(url, country string, ch chan<- interface{}, context *gin.Context) {
	var (
		thirdPartyResponse models.PopulationThirdPartyResponse
		err                error
	)
	responseBody, errMsg := make(map[int]interface{}), make(map[string]interface{})

	fmt.Println("inside make request", url)
	responsebytes, err := service.httpClient.Get(context, url)
	if err != nil {
		log.Printf("error while calling third party URL:%s Err:-%v", url, err)
		errMsg["country"] = country
		errMsg["ErrorMessage"] = err.Error()
		ch <- errMsg
		return
	}

	err = json.Unmarshal(responsebytes, &thirdPartyResponse)
	if err != nil {
		log.Printf("error while unmarshalling response for URL:%s Err:-%v", url, err)
		errMsg["country"] = country
		errMsg["ErrorMessage"] = err.Error()
		ch <- errMsg
		return
	}
	responseBody[thirdPartyResponse.TotalPopulation.Population] = country
	ch <- responseBody
}
