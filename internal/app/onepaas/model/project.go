package model

import (
	"github.com/onepaas/onepaas/internal/pkg/slug"
	"github.com/onepaas/onepaas/pkg/api/v1"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"

	"github.com/onepaas/onepaas/internal/pkg/ulid"
)

type Project struct {
	Id          string         `json:"id" gorm:"->;<-:create;type:string;size:26;primaryKey"`
	Name        string         `json:"name" gorm:"type:string;size:255;not null"`
	Slug        string         `json:"slug" gorm:"type:string;size:255;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Meta        datatypes.JSON `json:"-" gorm:"type:jsonb"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt  time.Time      `json:"modified_at" gorm:"autoUpdateTime"`
}

func NewProject(project v1.Project) *Project {
	return &Project{
		Name:        project.Spec.Name,
		Slug:        project.Spec.Slug,
		Description: project.Spec.Description,
	}
}

func (p *Project) BeforeCreate(_ *gorm.DB) (err error) {
	p.Id = ulid.Generate()
	p.Slug = slug.Make(p.Name)

	return
}

func (p *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Name") {
		return ErrorChangedColumn("Name not allowed to change")
	}

	return nil
}

func (p *Project) MarshalProjectAPI() v1.Project {
	return v1.Project{
		Metadata: v1.Metadata{
			UID:        p.Id,
			CreatedAt:  p.CreatedAt,
			ModifiedAt: p.ModifiedAt,
		},
		Spec: v1.ProjectSpec{
			Name:        p.Name,
			Slug:        p.Slug,
			Description: p.Description,
		},
	}
}

func (p *Project) UnmarshalProjectAPI(api v1.Project) {
	p.Name = api.Spec.Name
	p.Slug = api.Spec.Slug
	p.Description = api.Spec.Description
}
