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
	Cover            string          `json:"cover" gorm:"null"`
	Price            int             `json:"price" gorm:"not null"`
	AssessmentResult decimal.Decimal `json:"assessment_result" gorm:"type:decimal(10,3)"`
	Toilet           int             `json:"toilet" gorm:"not null;default:0"`
	Parking          int             `json:"parking" gorm:"not null;default:0"`
	Restarea         int             `json:"restarea" gorm:"not null;default:0"`
	Restaurant       int             `json:"restaurant" gorm:"not null;default:0"`
	Address          string          `json:"address" gorm:"not null"`
	Location         string          `json:"location" gorm:"not null"`
	CreatedAt        time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
