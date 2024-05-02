//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"lintangbs.org/lintang/template/app/start"
	"lintangbs.org/lintang/template/config"
)

func InitApp(cfg *config.Config, handler *gin.Engine) *start.InitWireApp {
	wire.Build(
		start.InitHTTPandGRPC,
	)

	return nil
}
