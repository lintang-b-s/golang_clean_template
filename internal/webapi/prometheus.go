package webapi

import (
	"context"
)

type BukanPromeApiClientAsli struct {
	message string
}

func NewClient(conf string) *BukanPromeApiClientAsli {

	return &BukanPromeApiClientAsli{
		message: conf,
	}
}

type PrometheusAPIImpl struct {
	client *BukanPromeApiClientAsli
}

func NewPrometheusAPI(adress string) *PrometheusAPIImpl {

	promeClient := NewClient("asdsad")

	return &PrometheusAPIImpl{client: promeClient}
}

func (b *PrometheusAPIImpl) Hello(ctx context.Context) (string, error) {
	return "asdsadsadsadd", nil
}
