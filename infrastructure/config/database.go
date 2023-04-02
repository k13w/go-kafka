package config

import (
	"database/sql"
	"github.com/k13w/go-kafka/core/domain/helper"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 7320
	user     = "kiew"
	password = "kiew"
	dbName   = "souls"
)

func DatabaseConnection() *sql.DB {
	//sqlInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", "postgres://kiew:kiew@localhost:7320/souls?sslmode=disable")
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!!")

	return db
}
