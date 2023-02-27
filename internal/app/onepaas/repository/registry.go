package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// RegistryRepository contains the interface for a model.Registry repository
type RegistryRepository interface {
	Create(ctx context.Context, registry *model.Registry) error
}

type registryRepository struct {
	*gorm.DB
}

// NewRegistryRepository creates a RegistryRepository
func NewRegistryRepository(db *gorm.DB) RegistryRepository {
	return &registryRepository{DB: db}
}

func (r *registryRepository) Create(_ context.Context, registry *model.Registry) error {
	result := r.DB.Create(&registry)

	return result.Error
}
