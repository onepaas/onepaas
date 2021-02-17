package model

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/onepaas/onepaas/internal/pkg/ulid"
)

type User struct {
	Id         string    `json:"id" pg:"type:varchar(26)"`
	Email      string    `json:"email" pg:"type:varchar(255)"`
	Password   string	 `json:"-" pg:"type:text"`
	Name       string    `json:"name" pg:"type:varchar(255)"`
	Meta       struct{}  `json:"-" pg:"type:jsonb"`
	CreatedAt  time.Time `json:"created_at" pg:"type:timestamptz"`
	ModifiedAt time.Time `json:"modified_at" pg:"type:timestamptz"`
}

var _ pg.BeforeInsertHook = (*User)(nil)

func (u *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now().UTC()

	u.Id = ulid.Generate()
	u.CreatedAt = now
	u.ModifiedAt = now

	return ctx, nil
}

var _ pg.BeforeUpdateHook = (*User)(nil)

func (u *User) BeforeUpdate(ctx context.Context) (context.Context, error) {
	u.ModifiedAt = time.Now().UTC()

	return ctx, nil
}
