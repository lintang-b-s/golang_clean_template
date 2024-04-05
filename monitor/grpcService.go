package monitor

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"lintangbs.org/lintang/template/pb"
)

type PrometheusApi interface {
	Hello(ctx context.Context) (string, error)
}

type MonitorServerImpl struct {
	pb.UnimplementedMonitorServiceServer
	prome PrometheusApi
}

func NewMonitorServer(prome PrometheusApi) *MonitorServerImpl {
	return &MonitorServerImpl{prome: prome}
}

func (server *MonitorServerImpl) Hello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {
	zap.L().Error("asldksadds", zap.Error(errors.New("asdsadas")))
	return &pb.HelloResponse{
		Res: "kontol",
	}, nil
}
