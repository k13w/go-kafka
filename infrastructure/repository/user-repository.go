package repository

import (
	"context"
	"database/sql"
	helper2 "github.com/k13w/go-kafka/core/domain/helper"
	"github.com/k13w/go-kafka/core/domain/model"
	"github.com/k13w/go-kafka/core/domain/port"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) port.UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (b UserRepositoryImpl) Save(ctx context.Context, user model.User) {
	tx, err := b.Db.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	createUserSql := `INSERT INTO "user" (nome, id) VALUES ($1, DEFAULT)`

	_, err = tx.ExecContext(ctx, createUserSql, user.Name)
	helper2.PanicIfError(err)
}
