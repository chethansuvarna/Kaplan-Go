package httpGeneric

import (
	"Kaplan-Go/config"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Client interface {
	Get(*gin.Context, string) (responseBody []byte, err error)
}

type HttpGenericClient struct {
	client http.Client
	config config.Config
}

func NewHttpGenericClient(timeout int, config config.Config) Client {
	return HttpGenericClient{client: http.Client{Timeout: time.Second * time.Duration(timeout)}, config: config}
}

// Get generic method to make get calls
func (httpClient HttpGenericClient) Get(context *gin.Context, url string) (responseBody []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	response, err := httpClient.client.Do(request)
	if err != nil {
		err = errors.New("error while making get Request")
		return
	}
	defer response.Body.Close()

	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New("error while reading Reponse")
		return
	}
	return
}
