package repository

import (
	"context"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"gorm.io/gorm"
)

//ProjectRepository contains the interface for a model.Project repository
type ProjectRepository interface {
	FindAll(ctx context.Context) ([]model.Project, error)
	FindByID(ctx context.Context, id string) (model.Project, error)
	Create(ctx context.Context, project *model.Project) error
	Update(ctx context.Context, project *model.Project, values model.Project) error
}

type projectRepository struct {
	*gorm.DB
}

//NewProjectRepository creates a ProjectRepository
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{DB: db}
}

func (p *projectRepository) FindAll(ctx context.Context) ([]model.Project, error) {
	projects := make([]model.Project, 0)
	result := p.DB.Find(&projects)

	return projects, result.Error
}

func (p *projectRepository) FindByID(ctx context.Context, id string) (model.Project, error) {
	var project model.Project
	result := p.DB.First(&project, "id = ?", id)

	return project, result.Error
}

func (p *projectRepository) Create(ctx context.Context, project *model.Project) error {
	result := p.DB.Create(&project)

	return result.Error
}

func (p *projectRepository) Update(ctx context.Context, project *model.Project, values model.Project) error {
	result := p.DB.Model(&project).Updates(values)

	return result.Error
}
