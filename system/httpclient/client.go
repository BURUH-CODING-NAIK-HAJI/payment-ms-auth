package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rizface/golang-api-template/app/errorgroup"
)

type HttpClientProperties struct {
	Url, Method string
	Body        map[string]interface{}
	Headers     map[string]string
}

func New(properties *HttpClientProperties) (*http.Response, error) {
	if properties.Body == nil {
		properties.Body = map[string]interface{}{}
	}
	payloadBytes, err := json.Marshal(properties.Body)
	if err != nil {
		panic(errorgroup.InternalServerError)
	}

	httprequest, err := http.NewRequest(properties.Method, properties.Url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(errorgroup.InternalServerError)
	}
	for key, v := range properties.Headers {
		httprequest.Header.Add(key, v)
	}

	client := new(http.Client)
	response, err := client.Do(httprequest)

	return response, err
}
