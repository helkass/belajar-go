package users

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,max=110,min=3"`
	Email string `json:"email" binding:"omitempty,email,max=110"`
}
