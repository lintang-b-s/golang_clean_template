package monitor

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"lintangbs.org/lintang/template/domain"
)

type ContainerRepository interface {
	GetById(ctx context.Context) (domain.Monitor, error)
}

type Service struct {
	containerRepo ContainerRepository
}

func NewService(c ContainerRepository) *Service {
	return &Service{
		containerRepo: c,
	}
}

func (m *Service) TesDoang(ctx context.Context) (domain.Monitor, error) {
	monitor, err := m.containerRepo.GetById(ctx)
	if err != nil {
		zap.L().Error("kntl", zap.Error(errors.New("ktl")))
	}
	zap.L().Debug("Hello sadoakdaas", zap.String("user", "lintang"),
		zap.Int("age", 20))
	zap.L().Error("Hello error", zap.String("user", "lintang"),
		zap.Int("age", 20))

	zap.L().Info("Hello error", zap.String("user", "lintang"),
		zap.Int("age", 20))

	zap.L().Warn("Hello error", zap.String("user", "lintang"),
		zap.Int("age", 20))
	return monitor, nil
}
