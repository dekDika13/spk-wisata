package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DetailCriteria struct {
	User             User            `gorm:"foreignKey:UserId"`
	Destination      Destination     `gorm:"foreignKey:DestinationId"`
	Criteria         Criteria        `gorm:"foreignKey:CriteriaId"`
	DetailCriteriaID uint            `json:"detail_criteria_id" gorm:"primaryKey;autoIncrement"`
	UserId           uint            `json:"user_id" gorm:"not null"`
	DestinationId    uint            `json:"destination_id" gorm:"not null"`
	CriteriaId       uint            `json:"criteria_id" gorm:"not null"`
	Rating           decimal.Decimal `json:"rating" gorm:"not null;type:decimal(10,2)"`
	CreatedAt        time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
