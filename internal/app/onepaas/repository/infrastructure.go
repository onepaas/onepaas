package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// InfrastructureRepository contains the interface for a model.Infrastructure repository
type InfrastructureRepository interface {
	Create(ctx context.Context, infra *model.Infrastructure) error
}

type infrastructureRepository struct {
	*gorm.DB
}

// NewInfraRepository creates a new ServerRepository
func NewInfraRepository(db *gorm.DB) InfrastructureRepository {
	return &infrastructureRepository{DB: db}
}

func (r *infrastructureRepository) Create(_ context.Context, infra *model.Infrastructure) error {
	result := r.DB.Create(&infra)

	return result.Error
}
