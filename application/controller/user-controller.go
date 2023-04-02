package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/k13w/go-kafka/core/domain/helper"
	"github.com/k13w/go-kafka/core/domain/usecase/user"
	"github.com/k13w/go-kafka/core/domain/utils/request"
	"github.com/k13w/go-kafka/core/domain/utils/response"
	"net/http"
)

type UserController struct {
	UseCase user.IUserUsecase
}

func NewUserController(usecase user.IUserUsecase) *UserController {
	return &UserController{
		UseCase: usecase,
	}
}

func (controller *UserController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadRequestBody(requests, &userCreateRequest)

	controller.UseCase.Create(requests.Context(), userCreateRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}
