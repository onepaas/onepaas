package model

import (
	"github.com/jackc/pgtype"
	"github.com/onepaas/onepaas/internal/pkg/ulid"
	"github.com/onepaas/onepaas/pkg/api/v1"
	"gorm.io/gorm"
	"time"
)

type Infrastructure struct {
	Id         string      `json:"id" gorm:"->;<-:create;type:string;size:26;primaryKey"`
	Type       string      `json:"type" gorm:"type:string;size:20;not null"`
	Properties pgtype.JSON `json:"properties" gorm:"type:json;default:{};not null"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt time.Time   `json:"modified_at" gorm:"autoUpdateTime"`
}

func NewInfrastructure(infra v1.Infrastructure) (*Infrastructure, error) {
	model := &Infrastructure{
		Type: infra.Spec.Type,
	}

	err := model.Properties.Set(infra.Spec.Properties)

	return model, err
}

func (m *Infrastructure) BeforeCreate(*gorm.DB) (err error) {
	m.Id = ulid.Generate()

	return
}

func (m *Infrastructure) MarshalInfrastructureAPI() (v1.Infrastructure, error) {
	apiModel := v1.Infrastructure{
		Metadata: v1.Metadata{
			UID:        m.Id,
			CreatedAt:  m.CreatedAt,
			ModifiedAt: m.ModifiedAt,
		},
		Spec: v1.InfrastructureSpec{
			Type: m.Type,
		},
	}

	err := m.Properties.AssignTo(&apiModel.Spec.Properties)

	return apiModel, err
}
