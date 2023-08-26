package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// InfrastructureRepository contains the interface for a model.Infrastructure repository
type InfrastructureRepository interface {
	Create(ctx context.Context, infra *model.Infrastructure) error
	FindByID(ctx context.Context, id string) (model.Infrastructure, error)
	FindAll(ctx context.Context) ([]model.Infrastructure, error)
}

type infrastructureRepository struct {
	*gorm.DB
}

// NewInfraRepository creates a new ServerRepository
func NewInfraRepository(db *gorm.DB) InfrastructureRepository {
	return &infrastructureRepository{DB: db}
}

func (r *infrastructureRepository) FindByID(_ context.Context, id string) (model.Infrastructure, error) {
	var record model.Infrastructure
	result := r.DB.First(&record, "id = ?", id)

	return record, result.Error
}

func (r *infrastructureRepository) FindAll(_ context.Context) ([]model.Infrastructure, error) {
	list := make([]model.Infrastructure, 0)
	result := r.DB.Find(&list)

	return list, result.Error
}

func (r *infrastructureRepository) Create(_ context.Context, infra *model.Infrastructure) error {
	result := r.DB.Create(&infra)

	return result.Error
}
