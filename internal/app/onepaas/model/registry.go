package model

import (
	"github.com/onepaas/onepaas/internal/pkg/ulid"
	"github.com/onepaas/onepaas/pkg/api/v1"
	"gorm.io/gorm"
	"time"
)

type Registry struct {
	Id         string    `json:"id" gorm:"->;<-:create;type:string;size:26;primaryKey"`
	URL        string    `json:"url" gorm:"type:text;not null"`
	Username   string    `json:"username" gorm:"type:string;size:255;not null"`
	Secret     string    `json:"secret" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt time.Time `json:"modified_at" gorm:"autoUpdateTime"`
}

func NewRegistry(registry v1.Registry) *Registry {
	return &Registry{
		URL:      registry.Spec.URL,
		Username: registry.Spec.Username,
		Secret:   registry.Spec.Secret,
	}
}

func (m *Registry) BeforeCreate(*gorm.DB) (err error) {
	m.Id = ulid.Generate()

	return
}

func (m *Registry) MarshalRegistryAPI() v1.Registry {
	return v1.Registry{
		Metadata: v1.Metadata{
			UID:        m.Id,
			CreatedAt:  m.CreatedAt,
			ModifiedAt: m.ModifiedAt,
		},
		Spec: v1.RegistrySpec{
			URL:      m.URL,
			Username: m.Username,
			Secret:   m.Secret,
		},
	}
}

func (m *Registry) UnmarshalRegistryAPI(registry v1.Registry) {
	m.URL = registry.Spec.URL
	m.Username = registry.Spec.Username
	m.Secret = registry.Spec.Secret
}
