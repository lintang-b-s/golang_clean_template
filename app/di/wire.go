//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"lintangbs.org/lintang/template/app/start"
	"lintangbs.org/lintang/template/config"
	"lintangbs.org/lintang/template/pkg/postgres"
)

func InitApp(cfg *config.Config, handler *gin.Engine) *postgres.Postgres {
	wire.Build(
		start.InitHTTPandGRPC,
	)

	return nil
}
