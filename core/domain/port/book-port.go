package port

import (
	"context"
	"github.com/k13w/go-kafka/core/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User)
}
