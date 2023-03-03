package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

// ServerRepository contains the interface for a model.Server repository
type ServerRepository interface {
	Create(ctx context.Context, server *model.Server) error
}

type serverRepository struct {
	*gorm.DB
}

// NewServerRepository creates a new ServerRepository
func NewServerRepository(db *gorm.DB) ServerRepository {
	return &serverRepository{DB: db}
}

func (r *serverRepository) Create(_ context.Context, server *model.Server) error {
	result := r.DB.Create(&server)

	return result.Error
}
