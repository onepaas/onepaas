package repository

import (
	"github.com/onepaas/onepaas/internal/app/onepaas/model"
	"github.com/onepaas/onepaas/internal/app/onepaas/types"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(cReq types.CreateUserRequest) (int64, error) {
	user := model.User{
		Email: cReq.Email,
		Name:  cReq.Name,
	}

	result := r.DB.Create(&user)

	return result.RowsAffected, result.Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)

	result := r.Where("email = ?", email).Take(&user)

	return user, result.Error
}
