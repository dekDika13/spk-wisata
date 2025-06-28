package model

import (
	"time"

	"gorm.io/gorm"
)

type DetailCriteria struct {
	User             User           `gorm:"foreignKey:UserId"`
	Destination      Destination    `gorm:"foreignKey:DestinationId"`
	Criteria         Criteria       `gorm:"foreignKey:CriteriaId"`
	DetailCriteriaID uint           `json:"detail_criteria_id" gorm:"primaryKey;autoIncrement"`
	UserId           uint           `json:"user_id" gorm:"not null"`
	DestinationId    uint           `json:"destination_id" gorm:"not null"`
	CriteriaId       uint           `json:"criteria_id" gorm:"not null"`
	Value            int            `json:"value" gorm:"not null;default:0"`
	CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
