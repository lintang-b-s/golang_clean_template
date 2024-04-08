package pgRepo

import (
	"context"

	"lintangbs.org/lintang/template/domain"
	"lintangbs.org/lintang/template/pkg/postgres"
)

type ContainerRepository struct {
	db *postgres.Postgres
}

func NewContainerRepo(db *postgres.Postgres) *ContainerRepository {
	return &ContainerRepository{db}
}

func (r *ContainerRepository) GetById(ctx context.Context) (domain.Monitor, error) {
	return domain.Monitor{Message: "asdsasd"}, nil
}
