package middleware_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/middleware"
	"github.com/stretchr/testify/assert"
)

func TestErrorMiddleware(t *testing.T) {
	godotenv.Load("../../.env")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := os.OpenFile("/notexistsfile.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
	})
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	middleware.ErrorHandler(handler).ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestErrorMiddlewareErrorValidator(t *testing.T) {
	godotenv.Load("../../.env")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginPayload := requestentity.Login{}
		panic(loginPayload.Validate())
	})
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	middleware.ErrorHandler(handler).ServeHTTP(recorder, request)
	responseBody, _ := io.ReadAll(recorder.Result().Body)
	response := &responseentity.Error{}
	json.Unmarshal(responseBody, response)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
