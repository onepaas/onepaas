package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/onepaas/onepaas/internal/pkg/ulid"
)

type User struct {
	Id         string     `json:"id" gorm:"type:string;size:26;primaryKey"`
	Name       string     `json:"name" gorm:"type:string;size:255"`
	Email      string     `json:"email" gorm:"type:string;size:255"`
	Meta       struct{}   `json:"-" gorm:"type:jsonb"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt time.Time  `json:"modified_at" gorm:"autoUpdateTime"`
	Identities []Identity `json:"-" gorm:"foreignKey:UserId"`
}

func (i *User) BeforeCreate(tx *gorm.DB) (err error) {
	i.Id = ulid.Generate()

	return
}
