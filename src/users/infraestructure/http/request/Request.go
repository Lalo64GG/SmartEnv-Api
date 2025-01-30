package request

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}