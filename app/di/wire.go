//go:build wireinject
// +build wireinject

package di

import (
	"net"

	"lintangbs.org/lintang/template/config"
	"lintangbs.org/lintang/template/internal/grpc"
	"lintangbs.org/lintang/template/internal/repository/postgres"
	"lintangbs.org/lintang/template/internal/rest"
	"lintangbs.org/lintang/template/internal/webapi"
	"lintangbs.org/lintang/template/monitor"
	"lintangbs.org/lintang/template/pb"
	"lintangbs.org/lintang/template/pkg/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// var ProviderSet = wire.NewSet(gorm.NewGorm)
var monitorSet = wire.NewSet(
	postgres.NewContainerRepo,
	monitor.NewService,
	wire.Bind(new(monitor.ContainerRepository), new(*postgres.ContainerRepository)),
	wire.Bind(new(rest.MonitorService), new(*monitor.Service)),
)

func InitRouterApi(*config.Config, *gin.Engine) *gin.RouterGroup {
	wire.Build(
		gorm.NewGorm,
		monitorSet,
		rest.NewRouter,
	)
	return nil
}

var monitorGrpcSet = wire.NewSet(
	webapi.NewPrometheusAPI,
	monitor.NewMonitorServer,
	wire.Bind(new(monitor.PrometheusApi), new(*webapi.PrometheusAPIImpl)),
	wire.Bind(new(pb.MonitorServiceServer), new(*monitor.MonitorServerImpl)),
)

func InitGrpcMonitorApi(promeAddress string, listener net.Listener) error {
	wire.Build(
		monitorGrpcSet,
		grpc.RunGRPCServer,
	)
	return nil
}
