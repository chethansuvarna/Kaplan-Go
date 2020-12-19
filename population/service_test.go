package population

import (
	"Kaplan-Go/config"
	httpMocks "Kaplan-Go/httpGeneric/mock"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPopulationService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Population Service")
}

var (
	baseURL              = "some"
	getpopulationURL     = "/population"
	helloWorld           = "/helloworld"
	getCountriesEndPoint = "countries"
)

var _ = Describe("population Service", func() {
	var (
		mockController        *gomock.Controller
		mockHTTPClient        *httpMocks.MockClient
		testConfigUtil        config.Util
		testPopulationService Service
		testConfig            config.Config
		testContext           *gin.Context
		responseRecorder      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockHTTPClient = httpMocks.NewMockClient(mockController)
		configMap := map[string]interface{}{
			"port":                  "8080",
			"baseURL":               "some/",
			"getCountriesEndPoint":  "countries",
			"getPopulationEndpoint": "population/",
		}
		testConfigUtil = config.NewConfigUtil(configMap)
		testConfig = config.NewConfig(testConfigUtil)
		testPopulationService = NewPopulationService(testConfig, mockHTTPClient)
		responseRecorder = httptest.NewRecorder()
		testContext, _ = gin.CreateTestContext(responseRecorder)
	})

	Context("When hello world is called", func() {
		It("should return hello world", func() {
			response := []byte(helloWroldResponse)

			resp := testPopulationService.HelloWorld(testContext)
			Expect(resp).To(Equal(string(response)))
		})
	})

	Context("When GetCountries() is called", func() {
		It("should return list of countries", func() {
			response, _ := json.Marshal(expectedResponse)
			mockHTTPClient.EXPECT().Get(testContext, getCountriesEndPoint).Return(response, nil).Times(1)
			_, err := testPopulationService.GetCountries(testContext)
			Expect(err).To(BeNil())
		})

		It("should return error when unmarshalling fails", func() {
			mockHTTPClient.EXPECT().Get(testContext, getCountriesEndPoint).Return(nil, errors.New("error")).Times(1)
			_, err := testPopulationService.GetCountries(testContext)
			Expect(err).NotTo(BeNil())
		})
	})
})
