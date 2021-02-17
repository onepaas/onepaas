package types

type CreateUserRequest struct {
	Email           string `json:"email" binding:"required,email,uniqueness=users;email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Name            string `json:"name"`
}
