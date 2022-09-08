package authcontroller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rizface/golang-api-template/app/controller/authcontroller"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/entity/securityentity"
	"github.com/rizface/golang-api-template/app/service/authservice"
	"github.com/rizface/golang-api-template/database/myredis"
	"github.com/rizface/golang-api-template/system/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) FindOne(username string) (*responseentity.User, error) {
	args := m.Mock.Called(username)
	argUsername := args.Get(0).(string)
	password, _ := bcrypt.GenerateFromPassword([]byte("dummy"), bcrypt.DefaultCost)
	if argUsername == "valid" {
		userId := uniuri.New()
		return &responseentity.User{
			Id:       userId,
			Username: "valid",
			Password: string(password),
			Profile: responseentity.Profile{
				Id:     uniuri.New(),
				UserId: userId,
				Name:   "Valid",
			},
		}, nil
	} else {
		return nil, errors.New("internal server error")
	}
}

func (m *Mock) Create(payload *requestentity.Register) (*responseentity.User, error) {
	payload = m.Called(payload).Get(0).(*requestentity.Register)
	if payload.Name == "Fariz" {
		userId := uniuri.New()
		return &responseentity.User{
			Id:       userId,
			Username: payload.Username,
			Password: nil,
			Profile: responseentity.Profile{
				Id:     uniuri.New(),
				UserId: userId,
				Name:   payload.Name,
			},
		}, nil
	}

	return nil, errors.New("Error")
}

func TestAuthControllerSuccess(t *testing.T) {
	repository := new(Mock)
	redis := myredis.New()
	service := authservice.New(repository, redis)
	controller := authcontroller.New(service)
	router := router.CreateRouter()
	authcontroller.Setup(router, controller)

	repository.On("FindOne", "valid").Return("valid")
	payload := requestentity.Login{
		Username: "valid",
		Password: "dummy",
	}

	payloadBytes, _ := json.Marshal(payload)
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payloadBytes))
	recorder := httptest.NewRecorder()
	controller.Login(recorder, request)

	responseBody, _ := io.ReadAll(recorder.Result().Body)
	generatedTokenSchema := new(securityentity.GeneratedResponseJwt)
	json.Unmarshal(responseBody, generatedTokenSchema)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, payload.Username, generatedTokenSchema.User.Username)
	assert.Equal(t, "string", reflect.TypeOf(generatedTokenSchema.Token.Bearer).String())
	assert.Equal(t, "string", reflect.TypeOf(generatedTokenSchema.Token.Refresh).String())
}

func TestAuthControllerFailedPayloadNotAllowed(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			_, ok := err.(validation.Errors)
			assert.True(t, ok)
		}
	}()

	repository := new(Mock)
	redis := myredis.New()
	service := authservice.New(repository, redis)
	controller := authcontroller.New(service)
	router := router.CreateRouter()
	authcontroller.Setup(router, controller)

	repository.On("FindOne", "valid").Return("valid")
	payload := requestentity.Login{}

	payloadBytes, _ := json.Marshal(payload)
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payloadBytes))
	recorder := httptest.NewRecorder()
	controller.Login(recorder, request)
}

func TestRegisterSuccess(t *testing.T) {
	repository := new(Mock)
	redis := myredis.New()
	service := authservice.New(repository, redis)
	controller := authcontroller.New(service)
	payload := &requestentity.Register{
		Name:     "Fariz",
		Username: "riz",
		Password: "password",
	}

	repository.On("Create", payload).Return(payload)
	bytesPayload, _ := json.Marshal(payload)
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bytesPayload))
	recorder := httptest.NewRecorder()
	controller.Register(recorder, request)

	userDataResponse := new(responseentity.User)
	byteResponse, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(byteResponse, userDataResponse)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, payload.Username, userDataResponse.Username)
}
