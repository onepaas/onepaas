package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// RegistryRepository contains the interface for a model.Registry repository
type RegistryRepository interface {
	Create(ctx context.Context, registry *model.Registry) error
	FindByID(_ context.Context, id string) (model.Registry, error)
	FindAll(ctx context.Context) ([]model.Registry, error)
}

type registryRepository struct {
	*gorm.DB
}

// NewRegistryRepository creates a RegistryRepository
func NewRegistryRepository(db *gorm.DB) RegistryRepository {
	return &registryRepository{DB: db}
}

func (r *registryRepository) FindByID(_ context.Context, id string) (model.Registry, error) {
	var record model.Registry
	result := r.DB.First(&record, "id = ?", id)

	return record, result.Error
}

func (r *registryRepository) Create(_ context.Context, registry *model.Registry) error {
	result := r.DB.Create(&registry)

	return result.Error
}

func (r *registryRepository) FindAll(_ context.Context) ([]model.Registry, error) {
	list := make([]model.Registry, 0)
	result := r.DB.Find(&list)

	return list, result.Error
}
