package user

import (
	"context"
	"github.com/k13w/go-kafka/core/domain/model"
	"github.com/k13w/go-kafka/core/domain/port"
	"github.com/k13w/go-kafka/core/domain/utils/request"
)

type IUserUsecase interface {
	Create(ctx context.Context, request request.UserCreateRequest)
}

type UserUsecase struct {
	UserRepository port.UserRepository
}

func NewUserUseCase(userRepository port.UserRepository) IUserUsecase {
	return &UserUsecase{UserRepository: userRepository}
}

func (u *UserUsecase) Create(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Name: request.Name,
	}

	u.UserRepository.Save(ctx, user)
}
