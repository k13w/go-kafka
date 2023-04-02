package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/k13w/go-kafka/application/controller"
)

func NewRouter(userController *controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", userController.Create)

	return router
}
