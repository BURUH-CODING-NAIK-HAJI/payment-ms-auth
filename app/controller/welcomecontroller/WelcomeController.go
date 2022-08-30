package welcomecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/service/welcomeservice"
)

type WelcomeControllerInterface interface {
	Welcome(w http.ResponseWriter, r *http.Request)
}

type WelcomeController struct {
	welcomeservice welcomeservice.WelcomeServiceInterface
}

func New(welcomeservice welcomeservice.WelcomeServiceInterface) WelcomeControllerInterface {
	return &WelcomeController{
		welcomeservice: welcomeservice,
	}
}

func Setup(router *chi.Mux, controller WelcomeControllerInterface) {
	router.Get("/", controller.Welcome)
}

func (welcome *WelcomeController) Welcome(w http.ResponseWriter, r *http.Request) {
	response := welcome.welcomeservice.Welcome()
	json.NewEncoder(w).Encode(response)
}
