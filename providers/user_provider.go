package providers

import (
	"belajar-go/controllers"
	"belajar-go/repositories"
	"belajar-go/services"

	"gorm.io/gorm"
)

type UserProvider struct {
	UserController *controllers.UserController
}

func NewUserProvider(db *gorm.DB) *UserProvider {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	return &UserProvider{
		UserController: controller,
	}
}
