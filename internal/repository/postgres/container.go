package postgres

import (
	"context"

	"lintangbs.org/lintang/template/domain"
	"lintangbs.org/lintang/template/pkg/gorm"
)

type ContainerRepository struct {
	DB *gorm.Gorm
}

func NewContainerRepo(db *gorm.Gorm) *ContainerRepository {
	return &ContainerRepository{db}
}

func (r *ContainerRepository) GetById(ctx context.Context) (domain.Monitor, error ) {
	return domain.Monitor{Message: "asdsasd"}, nil
}
