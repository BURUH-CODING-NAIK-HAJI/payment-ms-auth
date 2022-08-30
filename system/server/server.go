package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/controller/welcomecontroller"
	"github.com/rizface/golang-api-template/app/repository/welcomerepository"
	"github.com/rizface/golang-api-template/app/service/welcomeservice"
)

func SetupController(router *chi.Mux) {
	// SETUP WELCOME REPOSITORY
	welcomerepository := welcomerepository.New()
	// SETUP WELCOME SERVICE
	welcomeService := welcomeservice.New(welcomerepository)
	// SETUP WELCOME CONTROLLER
	welcomeController := welcomecontroller.New(welcomeService)
	welcomecontroller.Setup(router, welcomeController)
}

func CreateHttpServer(router http.Handler) *http.Server {
	var appPort string

	if len(appPort) == 0 {
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
