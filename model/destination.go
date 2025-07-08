package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Destination struct {
	DestinationID    uint            `json:"destination_id" gorm:"primaryKey;autoIncrement"`
	Name             string          `json:"name" gorm:"not null"`
	Description      string          `json:"description" gorm:"not null"`
	Cover            string          `json:"cover" gorm:"not null"`
	Price            int             `json:"price" gorm:"not null"`
	AverageRating    decimal.Decimal `json:"average_rating" gorm:"type:decimal(10,3)"`
	AssessmentResult decimal.Decimal `json:"assessment_result" gorm:"type:decimal(10,3)"`
	Address          string          `json:"address" gorm:"not null"`
	Location         string          `json:"location" gorm:"not null"`
	CreatedAt        time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
