package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"github.com/onepaas/onepaas/internal/app/onepaas/types"
)

type userRepository struct {
	*pg.DB
}

func NewUserRepository(db *pg.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(cReq types.CreateUserRequest) (pg.Result, error) {
	user := model.User{
		Username: cReq.Username,
		Email: cReq.Email,
		Name: cReq.Name,
	}

	return r.Model(&user).Insert()
}
