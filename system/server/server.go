package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/controller/authcontroller"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"github.com/rizface/golang-api-template/app/service/authservice"
	"github.com/rizface/golang-api-template/database/myredis"
)

func SetupController(router *chi.Mux) {

	userRepository := userrepository.New()
	redis := myredis.New()
	authService := authservice.New(userRepository, redis)
	authController := authcontroller.New(authService)
	authcontroller.Setup(router, authController)
}

func CreateHttpServer(router http.Handler) *http.Server {
	var appPort string

	if len(os.Getenv("APP_PORT")) == 0 {
		appPort = "9000"
	} else {
		appPort = os.Getenv("APP_PORT")
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", appPort),
		Handler: router,
	}

	return httpServer
}
