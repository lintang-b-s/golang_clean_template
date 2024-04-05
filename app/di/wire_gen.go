// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"lintangbs.org/lintang/template/config"
	"lintangbs.org/lintang/template/internal/grpc"
	"lintangbs.org/lintang/template/internal/repository/postgres"
	"lintangbs.org/lintang/template/internal/rest"
	"lintangbs.org/lintang/template/internal/webapi"
	"lintangbs.org/lintang/template/monitor"
	"lintangbs.org/lintang/template/pb"
	"lintangbs.org/lintang/template/pkg/gorm"
	"net"
)

// Injectors from wire.go:

func InitRouterApi(configConfig *config.Config, engine *gin.Engine) *gin.RouterGroup {
	gormGorm := gorm.NewGorm(configConfig)
	containerRepository := postgres.NewContainerRepo(gormGorm)
	service := monitor.NewService(containerRepository)
	routerGroup := rest.NewRouter(engine, service)
	return routerGroup
}

func InitGrpcMonitorApi(promeAddress string, listener net.Listener) error {
	prometheusAPIImpl := webapi.NewPrometheusAPI(promeAddress)
	monitorServerImpl := monitor.NewMonitorServer(prometheusAPIImpl)
	error2 := grpc.RunGRPCServer(monitorServerImpl, listener)
	return error2
}

// wire.go:

// var ProviderSet = wire.NewSet(gorm.NewGorm)
var monitorSet = wire.NewSet(postgres.NewContainerRepo, monitor.NewService, wire.Bind(new(monitor.ContainerRepository), new(*postgres.ContainerRepository)), wire.Bind(new(rest.MonitorService), new(*monitor.Service)))

var monitorGrpcSet = wire.NewSet(webapi.NewPrometheusAPI, monitor.NewMonitorServer, wire.Bind(new(monitor.PrometheusApi), new(*webapi.PrometheusAPIImpl)), wire.Bind(new(pb.MonitorServiceServer), new(*monitor.MonitorServerImpl)))
