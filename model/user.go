package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Profile   *Profile       `gorm:"foreignKey:UserId"`
	UserID    uint           `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Role      uint           `json:"role" gorm:"not null"`
	Username  string         `json:"username" gorm:"not null;unique"`
	Password  string         `json:"password" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Phone     string         `json:"phone" gorm:"not null;unique"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
