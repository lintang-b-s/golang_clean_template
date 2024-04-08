package start

import (
	"net"

	"lintangbs.org/lintang/template/config"
	"lintangbs.org/lintang/template/internal/grpc"
	"lintangbs.org/lintang/template/internal/repository/pgRepo"
	"lintangbs.org/lintang/template/internal/rest"
	"lintangbs.org/lintang/template/internal/webapi"
	"lintangbs.org/lintang/template/monitor"
	"lintangbs.org/lintang/template/pkg/postgres"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitHTTPandGRPC(cfg *config.Config, handler *gin.Engine) *postgres.Postgres {
	// Router
	pg := postgres.NewPostgres(cfg)
	containerRepository := pgRepo.NewContainerRepo(pg)
	service := monitor.NewService(containerRepository)
	rest.NewRouter(handler, service)

	address := "0.0.0.0:5099"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		postgres.ClosePostgres(pg.Pool)
		zap.L().Fatal("cannot start server: ", zap.Error(err))
	}

	// GRPC
	prometheusAPI := webapi.NewPrometheusAPI("asdsadsaas")
	monitorServerImpl := monitor.NewMonitorServer(prometheusAPI)
	err = grpc.RunGRPCServer(monitorServerImpl, listener)
	if err != nil {
		postgres.ClosePostgres(pg.Pool)
		zap.L().Fatal("cannot start GRPC  Server", zap.Error(err))
	}
	return pg
}
