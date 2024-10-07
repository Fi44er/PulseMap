package model

import (
	"time"

	"github.com/Fi44er/PulseMap/backend/auth_service/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string     `json:"id" gorm:"type:string;unique;not null;index;primary_key"`
	Username     string     `json:"username" gorm:"type:string;not null;unique"`
	Email        string     `json:"email" gorm:"type:string;unique;not null"`
	PasswordHash string     `json:"password_hash" gorm:"type:string;not null"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	IsActive     bool       `json:"is_active" gorm:"type:bool;default:true"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.PasswordHash = utils.HashPass([]byte(user.PasswordHash))
	return nil
}
