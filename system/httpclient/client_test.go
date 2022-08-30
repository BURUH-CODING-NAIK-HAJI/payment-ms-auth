package httpclient_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/rizface/golang-api-template/system/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestHttpGet(t *testing.T) {
	var responseBody []map[string]interface{}

	properties := httpclient.HttpClientProperties{
		Url:     "https://jsonplaceholder.typicode.com/users",
		Method:  http.MethodGet,
		Body:    nil,
		Headers: nil,
	}
	response, err := httpclient.New(&properties)
	if err != nil {
		fmt.Println(err.Error())
	}
	responseBytes, _ := io.ReadAll(response.Body)
	json.Unmarshal(responseBytes, &responseBody)
	assert.True(t, len(responseBody) > 0)
}

func TestHttpPost(t *testing.T) {
	properties := httpclient.HttpClientProperties{
		Url:    "https://eok5svyykgib69k.m.pipedream.net",
		Method: http.MethodPost,
		Body: map[string]interface{}{
			"name": "Jhon Doe",
		},
		Headers: map[string]string{
			"X-POWERED-BY": "My Self",
		},
	}
	response, err := httpclient.New(&properties)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.True(t, response.StatusCode == 200)
}
