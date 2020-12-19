package httpGeneric

import (
	"Kaplan-Go/config"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gin-gonic/gin"
)

func TestHTTPClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Http client")
}

var _ = Describe("population Service", func() {
	var (
		testHTTPClient   Client
		testConfig       config.Config
		testContext      *gin.Context
		responseRecorder *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		configMap := map[string]interface{}{
			"port": "8080",
		}
		testConfigUtil := config.NewConfigUtil(configMap)
		testConfig = config.NewConfig(testConfigUtil)
		testContext, _ = gin.CreateTestContext(responseRecorder)
		testContext.Request, _ = http.NewRequest("GET", "/", nil)
		testHTTPClient = NewHttpGenericClient(100, testConfig)
	})

	Context("When GET() is called", func() {
		It("should return valid", func() {
			responseBody := []byte("hello world")
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.Write(responseBody)
			}))
			actualResponseBody, _ := testHTTPClient.Get(testContext, server.URL)
			Expect(actualResponseBody).To(Equal(responseBody))
		})

		It("should return error when error occurs while get call", func() {
			// responseBody := []byte("hello world")
			// server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// 	rw.Write(responseBody)
			// }))
			actualResponseBody, err := testHTTPClient.Get(testContext, "")
			Expect(actualResponseBody).To(BeNil())
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("error while making get Request"))

		})
	})
})
