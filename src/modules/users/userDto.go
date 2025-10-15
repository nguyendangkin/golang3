package users

type RegisterUserRequest struct {
	FullName string `json:"fullName" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
