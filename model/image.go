package model

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	Gallery   Gallery        `gorm:"foreignKey:GalleryId"`
	ImageID   uint           `json:"image_id" gorm:"primaryKey;autoIncrement"`
	GalleryId uint           `json:"gallery_id" gorm:"null"`
	ImageURL  string         `json:"image_url" gorm:"null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
