package repositories

import (
	"belajar-go/models"

	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll(query UserQuery) ([]models.User, error)
	WithDB(db *gorm.DB) UserRepository
	Delete(id string) error
	Update(id string, user *models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type UserQuery struct {
	Limit  int
	Offset int
	Name   string
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) WithDB(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll(q UserQuery) ([]models.User, error) {
	var users []models.User

	db := r.db

	if q.Name != "" {
		db = db.Where("name LIKE ?", "%"+q.Name+"%")
	}

	if q.Limit > 0 {
		db = db.Limit(q.Limit)
	}

	if q.Offset > 0 {
		db = db.Offset(q.Offset)
	}

	err := db.Find(&users).Error

	return users, err
}

func (r *userRepository) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&models.User{})

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return result.Error
}

func (r *userRepository) Update(id string, user *models.User) (models.User, error) {
	var updatedUser models.User

	result := r.db.Model(&updatedUser).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(user)

	if result.RowsAffected == 0 {
		return models.User{}, errors.New("user not found")
	}

	return updatedUser, nil
}
