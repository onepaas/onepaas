package model

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/onepaas/onepaas/internal/pkg/ulid"
)

type User struct {
	Id         string    `pg:"type:varchar(26)"`
	Username   string    `pg:"type:varchar(255)"`
	Email      string    `pg:"type:varchar(255)"`
	Name       string    `pg:"type:varchar(255)"`
	Meta       struct{}  `pg:"type:jsonb"`
	CreatedAt  time.Time `pg:"type:timestamptz"`
	ModifiedAt time.Time `pg:"type:timestamptz"`
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
