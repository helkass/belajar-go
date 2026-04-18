package providers

import (
	"belajar-go/controllers"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controllers.UserController
}

func NewContainer(db *gorm.DB) *Container {
	// panggil user provider
	userProvider := NewUserProvider(db)

	return &Container{
		UserController: userProvider.UserController,
	}
}
