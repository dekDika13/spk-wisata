package model

import "time"

type Gallery struct {
	Destination   Destination `gorm:"foreignKey:DestinationId"`
	GalleryID     uint        `json:"gallery_id" gorm:"primaryKey;autoIncrement"`
	DestinationId uint        `json:"destination_id" gorm:"not null"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
}
