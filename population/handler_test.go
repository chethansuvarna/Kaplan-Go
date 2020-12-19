package population

import (
	mocks "Kaplan-Go/population/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPopulationHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Population Handler")
}

var _ = Describe("Population Handler", func() {
	var (
		mockController        *gomock.Controller
		mockPopulationService *mocks.MockService
		populationHandler     Handler
		testContext           *gin.Context
		responseRecorder      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockPopulationService = mocks.NewMockService(mockController)
		populationHandler = NewPopulationHandler(mockPopulationService)
		responseRecorder = httptest.NewRecorder()
		testContext, _ = gin.CreateTestContext(responseRecorder)
	})

	Context("when Hello World is called", func() {

		It("should return valid response", func() {
			testContext.Request = httptest.NewRequest("GET", "/helloworld", nil)
			expectedResponse := "hello world"

			mockPopulationService.EXPECT().HelloWorld(testContext).Return(expectedResponse).Times(1)

			populationHandler.HelloWorld(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			Expect(strings.TrimSuffix(responseRecorder.Body.String(), "\n")).To(Equal("\"hello world\""))
		})
	})

	Context("when GetCountries is called", func() {

		It("should return valid response", func() {
			testContext.Request = httptest.NewRequest("GET", "/countries", nil)
			responseBodyBytes, _ := json.Marshal(expectedResponse)

			mockPopulationService.EXPECT().GetCountries(testContext).Return(expectedResponse, nil).Times(1)

			populationHandler.GetCountries(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			Expect(strings.TrimSuffix(responseRecorder.Body.String(), "\n")).To(Equal(string(responseBodyBytes)))
		})

		It("should return valid response", func() {
			testContext.Request = httptest.NewRequest("GET", "/countries", nil)

			mockPopulationService.EXPECT().GetCountries(testContext).Return(expectedResponse, errors.New("error")).Times(1)

			populationHandler.GetCountries(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("when GetPopulationForCountry() is called", func() {

		It("should return countries and population in ascending order", func() {
			request, _ := json.Marshal(getCountriesGoRequestAscending)
			testContext.Request = httptest.NewRequest("GET", "/countriesgo", strings.NewReader(string(request)))

			responseBodyBytes, _ := json.Marshal(getCountriesGoExpectedResponseAscending)

			mockPopulationService.EXPECT().GetPopulationForCountry(testContext, getCountriesGoRequestAscending).Return(getCountriesGoExpectedResponseAscending, nil).Times(1)

			populationHandler.GetPopulationForCountry(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			Expect(strings.TrimSuffix(responseRecorder.Body.String(), "\n")).To(Equal(string(responseBodyBytes)))
		})

		It("should return countries and population in ascending order", func() {
			request, _ := json.Marshal(getCountriesGoRequestAscending)
			testContext.Request = httptest.NewRequest("GET", "/countriesgo", strings.NewReader(string(request)))

			mockPopulationService.EXPECT().GetPopulationForCountry(testContext, getCountriesGoRequestAscending).Return(getCountriesGoExpectedResponseAscending, errors.New("error")).Times(1)

			populationHandler.GetPopulationForCountry(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
		})

		It("should return countries and population in ascending order", func() {
			request := []byte("{")
			testContext.Request = httptest.NewRequest("GET", "/countriesgo", strings.NewReader(string(request)))

			mockPopulationService.EXPECT().GetPopulationForCountry(testContext, getCountriesGoRequestAscending).Return(getCountriesGoExpectedResponseAscending, errors.New("error")).Times(1)

			populationHandler.GetPopulationForCountry(testContext)
			Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
		})

	})
})
