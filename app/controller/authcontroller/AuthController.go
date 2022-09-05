package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/app/service/authservice"
)

type AuthControllerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type AuthController struct {
	authService authservice.AuthServiceInterface
}

func New(authService authservice.AuthServiceInterface) AuthControllerInterface {
	return &AuthController{
		authService: authService,
	}
}

func Setup(router *chi.Mux, controller AuthControllerInterface) *chi.Mux {
	// router.Route("/v1/auth", func(r chi.Router) {
	// 	r.Post("/login", controller.Login)
	// 	r.Post("/register", controller.Register)
	// })
	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
	return router
}

func (auth *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	payload := new(requestentity.Login)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		panic(errorgroup.InternalServerError)
	}

	err = payload.Validate()
	if err != nil {
		panic(err)
	}

	result := auth.authService.Login(payload)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (auth *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	payload := new(requestentity.Register)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		panic(errorgroup.InternalServerError)
	}

	err = payload.Validate()
	if err != nil {
		panic(err)
	}

	result := auth.authService.Register(payload)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
