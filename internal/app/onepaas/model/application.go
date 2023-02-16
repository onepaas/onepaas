package model

import (
	"github.com/onepaas/onepaas/internal/pkg/ulid"
	"github.com/onepaas/onepaas/pkg/api/v1"
	"gorm.io/gorm"
	"time"
)

type Application struct {
	Id            string    `json:"id" gorm:"->;<-:create;type:string;size:26;primaryKey"`
	Name          string    `json:"name" gorm:"type:string;size:255;not null"`
	RepositoryURL string    `json:"repository_url" gorm:"type:text;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt    time.Time `json:"modified_at" gorm:"autoUpdateTime"`
}

func NewApplication(application v1.Application) *Application {
	return &Application{
		Name:          application.Spec.Name,
		RepositoryURL: application.Spec.RepositoryURL,
	}
}

func (a *Application) BeforeCreate(*gorm.DB) (err error) {
	a.Id = ulid.Generate()

	return
}

func (a *Application) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Name") {
		return ErrorChangedColumn("Name can not be changed")
	}

	return nil
}

func (a *Application) MarshalApplicationAPI() v1.Application {
	return v1.Application{
		Metadata: v1.Metadata{
			UID:        a.Id,
			CreatedAt:  a.CreatedAt,
			ModifiedAt: a.ModifiedAt,
		},
		Spec: v1.ApplicationSpec{
			Name:          a.Name,
			RepositoryURL: a.RepositoryURL,
		},
	}
}

func (a *Application) UnmarshalApplicationAPI(api v1.Application) {
	a.Name = api.Spec.Name
	a.RepositoryURL = api.Spec.RepositoryURL
}
