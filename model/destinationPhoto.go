package model

type DestinationPhoto struct {
	Destination        Destination `gorm:"foreignKey:DestinationId"`
	DestinationPhotoID uint        `json:"destination_photo_id" gorm:"primaryKey;autoIncrement"`
	DestinationId      uint        `json:"destination_id" gorm:"not null"`
	PhotoURL           string      `json:"photo_url" gorm:"not null"`
	CreatedAt          string      `json:"created_at" gorm:"autoCreateTime"`
}
