package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Review struct {
	Destination   Destination     `gorm:"fereignKey:DestinationId"`
	User          User            `gorm:"foreignKey:UserId"`
	ReviewID      uint            `json:"review_id" gorm:"primaryKey;autoIncrement"`
	DestinationId uint            `json:"destination_id" gorm:"not null"`
	UserId        uint            `json:"user_id" gorm:"not null;"`
	ReviewDetail  string          `json:"review_detail" gorm:"not null"`
	Rating        decimal.Decimal `json:"rating" gorm:"not null;type:decimal(10,2)"`
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
