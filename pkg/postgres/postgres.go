package postgres

import (
	"context"
	"database/sql"
	"net/url"

	"lintangbs.org/lintang/template/config"

	"go.uber.org/zap"
)

type Postgres struct {
	Pool *sql.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("postgres", "pass"),
		Path:   "dogker",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		zap.L().Fatal("sql.Open", zap.Error(err))
	}

	if err := db.PingContext(context.Background()); err != nil {
		zap.L().Fatal("db.PingContext", zap.Error(err))
	}

	return &Postgres{db}
}

func ClosePostgres(pg *sql.DB) {
	_ = pg.Close()
	zap.L().Debug("postgres closed!")

}
