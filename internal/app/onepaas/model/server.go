package model

import (
	"github.com/jackc/pgtype"
	"github.com/onepaas/onepaas/internal/pkg/ulid"
	"github.com/onepaas/onepaas/pkg/api/v1"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	Id         string      `json:"id" gorm:"->;<-:create;type:string;size:26;primaryKey"`
	Type       string      `json:"type" gorm:"type:string;size:20;not null"`
	Properties pgtype.JSON `json:"properties" gorm:"type:json;default:{};not null"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt time.Time   `json:"modified_at" gorm:"autoUpdateTime"`
}

func NewServer(server v1.Server) (*Server, error) {
	model := &Server{
		Type: server.Spec.Type,
	}

	err := model.Properties.Set(server.Spec.Properties)

	return model, err
}

func (m *Server) BeforeCreate(*gorm.DB) (err error) {
	m.Id = ulid.Generate()

	return
}

func (m *Server) MarshalServerAPI() (v1.Server, error) {
	apiModel := v1.Server{
		Metadata: v1.Metadata{
			UID:        m.Id,
			CreatedAt:  m.CreatedAt,
			ModifiedAt: m.ModifiedAt,
		},
		Spec: v1.ServerSpec{
			Type: m.Type,
		},
	}

	err := m.Properties.AssignTo(&apiModel.Spec.Properties)

	return apiModel, err
}
