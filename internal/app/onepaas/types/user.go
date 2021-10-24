package types

type CreateUserRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email" binding:"required,email,uniqueness=users;email"`
}
