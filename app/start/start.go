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
	"lintangbs.org/lintang/template/pkg/rabbitmq"

	grpcClient "google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InitWireApp struct {
	PG         *postgres.Postgres
	RMQ        *rabbitmq.RabbitMQ
	GRPCServer *grpcClient.Server
}

func InitHTTPandGRPC(cfg *config.Config, handler *gin.Engine) *InitWireApp {
	// Router
	pg := postgres.NewPostgres(cfg)
	containerRepository := pgRepo.NewContainerRepo(pg)
	service := monitor.NewService(containerRepository)
	rest.NewRouter(handler, service)
	rmq := rabbitmq.NewRabbitMQ(cfg)

	address := cfg.GRPC.URLGrpc
	listener, err := net.Listen("tcp", address)
	if err != nil {
		zap.L().Fatal("cannot start server: ", zap.Error(err))
	}

	// GRPC

	prometheusAPI := webapi.NewPrometheusAPI("asdsadsaas")
	monitorServerImpl := monitor.NewMonitorServer(prometheusAPI)

	grpcServerChan := make(chan *grpcClient.Server)

	go func() {
		err := grpc.RunGRPCServer(monitorServerImpl, listener, grpcServerChan)
		if err != nil {
			zap.L().Fatal("cannot start GRPC  Server", zap.Error(err))
		}
	}()

	var grpcServer = <-grpcServerChan

	return &InitWireApp{
		PG:         pg,
		RMQ:        rmq,
		GRPCServer: grpcServer,
	}
}
