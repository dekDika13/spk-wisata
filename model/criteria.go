package model

import (
	"time"

	"gorm.io/gorm"
)

type Criteria struct {
	CriteriaID  uint           `json:"criteria_id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
