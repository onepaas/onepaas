package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// ApplicationRepository contains the interface for a model.application repository
type ApplicationRepository interface {
	FindAll(ctx context.Context) ([]model.Application, error)
	FindByID(ctx context.Context, id string) (model.Application, error)
	Create(ctx context.Context, application *model.Application) error
	Update(ctx context.Context, application *model.Application, values model.Application) error
}

type applicationRepository struct {
	*gorm.DB
}

// NewApplicationRepository creates an ApplicationRepository
func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{DB: db}
}

func (a *applicationRepository) FindAll(context.Context) ([]model.Application, error) {
	applications := make([]model.Application, 0)
	result := a.DB.Find(&applications)

	return applications, result.Error
}

func (a *applicationRepository) FindByID(_ context.Context, id string) (model.Application, error) {
	var application model.Application
	result := a.DB.First(&application, "id = ?", id)

	return application, result.Error
}

func (a *applicationRepository) Create(_ context.Context, application *model.Application) error {
	result := a.DB.Create(&application)

	return result.Error
}

func (a *applicationRepository) Update(_ context.Context, application *model.Application, values model.Application) error {
	result := a.DB.Model(&application).Updates(values)

	return result.Error
}
