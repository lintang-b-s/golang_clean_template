package webapi

import (
	"context"

	"go.uber.org/zap"
)

type BukanPromeApiClientAsli struct {
	message string
}

func  NewClient(conf string) *BukanPromeApiClientAsli {
	
	return &BukanPromeApiClientAsli{
		message: conf,
	}
}

type PrometheusAPIImpl struct {
	
	client *BukanPromeApiClientAsli
}

func NewPrometheusAPI(adress string) *PrometheusAPIImpl {
	
	promeClient := NewClient("asdsad")
	zap.L().Info("kontollasddasdsdsdsdsdsdsadsd")

	return &PrometheusAPIImpl{client: promeClient}
}


func (b *PrometheusAPIImpl) Hello(ctx context.Context) (string, error) {
	return "asdsadsadsadd", nil
}



