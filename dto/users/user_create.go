package users

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,max=110,min=3"`
	Email string `json:"email" binding:"required,email,max=110"`
}
