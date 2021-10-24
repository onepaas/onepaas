package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/onepaas/onepaas/internal/pkg/ulid"
)

type Identity struct {
	Id         string    `json:"-" gorm:"type:string;size:26;primaryKey"`
	UserId     string    `json:"-" gorm:"type:string;size:26"`
	Subject    string	 `json:"-" gorm:"type:text;not null"`
	Provider   string	 `json:"-" gorm:"type:text;not null"`
	Meta       struct{}  `json:"-" gorm:"type:jsonb"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	ModifiedAt time.Time `json:"-" gorm:"autoUpdateTime"`
	User		User	 `gorm:"foreignKey:Id"`
}

func (i *Identity) BeforeCreate(tx *gorm.DB) (err error) {
	i.Id = ulid.Generate()

	return
}
