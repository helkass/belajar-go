package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Email string    `json:"email" gorm:"unique"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
