package welcomeservice

import (
	"github.com/rizface/golang-api-template/app/repository/welcomerepository"
)

type WelcomeServiceInterface interface {
	Welcome() string
}

type WelcomeService struct {
	welcomerepository welcomerepository.WelcomeRepositoryInterface
}

func New(welcomerepository welcomerepository.WelcomeRepositoryInterface) WelcomeServiceInterface {
	return &WelcomeService{
		welcomerepository: welcomerepository,
	}
}

func (welcomeservice *WelcomeService) Welcome() string {
	response := welcomeservice.welcomerepository.Welcome()
	return response
}
