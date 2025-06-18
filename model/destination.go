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
	Image1           string          `json:"image1" gorm:"not null"`
	Image2           string          `json:"image2" gorm:"not null"`
	Image3           string          `json:"image3" gorm:"not null"`
	Image4           string          `json:"image4" gorm:"not null"`
	Price            int             `json:"price" gorm:"not null"`
	AssessmentResult decimal.Decimal `json:"assessment_result" gorm:"type:decimal(10,3)"`
	Address          string          `json:"address" gorm:"not null"`
	Location         string          `json:"location" gorm:"not null"`
	CreatedAt        time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
