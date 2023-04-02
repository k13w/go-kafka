package main

import (
	"github.com/k13w/go-kafka/application/controller"
	"github.com/k13w/go-kafka/application/router"
	"github.com/k13w/go-kafka/core/domain/helper"
	usecase2 "github.com/k13w/go-kafka/core/domain/usecase/user"
	"github.com/k13w/go-kafka/infrastructure/config"
	repository2 "github.com/k13w/go-kafka/infrastructure/repository"
	"net/http"
)

func main() {
	db := config.DatabaseConnection()

	repository := repository2.NewBookRepository(db)

	usecase := usecase2.NewUserUseCase(repository)

	userController := controller.NewUserController(usecase)

	routes := router.NewRouter(userController)

	server := http.Server{Addr: "localhost:7000", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
