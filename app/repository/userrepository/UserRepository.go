package userrepository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
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

	user := new(responseentity.User)
	bytesResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytesResponse, user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (userrepository *UserRepository) Create(payload *requestentity.Register) (*responseentity.User, error) {
	return nil, nil
}
