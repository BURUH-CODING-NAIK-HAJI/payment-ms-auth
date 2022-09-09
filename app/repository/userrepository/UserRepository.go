package userrepository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/system/httpclient"
)

type UserRepositoryInterface interface {
	FindOne(username string) (*responseentity.User, error)
	Create(payload *requestentity.Register) (*responseentity.User, error)
}

type UserRepository struct {
}

func New() UserRepositoryInterface {
	return &UserRepository{}
}

func generateErrorGroup(bytesResponse []byte) error {
	errGroup := errorgroup.Error{}
	json.Unmarshal(bytesResponse, &errGroup)
	return errGroup
}

func (userrepository *UserRepository) FindOne(username string) (*responseentity.User, error) {
	client := httpclient.HttpClientProperties{
		Url:     fmt.Sprintf(os.Getenv("MS_USER_BASE")+"/%s/resources", username),
		Method:  http.MethodGet,
		Body:    nil,
		Headers: nil,
	}
	response, err := httpclient.New(&client)
	if err != nil {
		return nil, err
	}

	bytesResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		user := new(responseentity.User)
		err = json.Unmarshal(bytesResponse, user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
	return nil, generateErrorGroup(bytesResponse)
}

func (userrepository *UserRepository) Create(payload *requestentity.Register) (*responseentity.User, error) {
	client := httpclient.HttpClientProperties{
		Url:    os.Getenv("MS_USER_BASE"),
		Method: http.MethodPost,
		Body: map[string]interface{}{
			"name":     payload.Name,
			"username": payload.Username,
			"password": payload.Password,
		},
		Headers: nil,
	}

	response, err := httpclient.New(&client)
	if err != nil {
		return nil, err
	}

	bytesResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		user := new(responseentity.User)
		err = json.Unmarshal(bytesResponse, user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, generateErrorGroup(bytesResponse)
}
