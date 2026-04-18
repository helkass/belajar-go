package services

import (
	"belajar-go/dto/users"
	"belajar-go/models"
	"belajar-go/repositories"
)

type UserService interface {
	CreateUser(req users.CreateUserRequest) (models.User, error)
	GetUsers() ([]models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, req users.UpdateUserRequest) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req users.CreateUserRequest) (models.User, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.repo.Create(&user)
	return user, err
}

func (s *userService) GetUsers() ([]models.User, error) {
	users, err := s.repo.FindAll(repositories.UserQuery{})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) DeleteUser(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(id string, req users.UpdateUserRequest) (models.User, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	updatedUser, err := s.repo.Update(id, &user)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}
