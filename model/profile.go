package model

import "time"

type Profile struct {
	User      User      `gorm:"foreignKey:UserId"`
	ProfileID uint      `json:"profile_id" gorm:"primaryKey;autoIncrement"`
	UserId    uint      `json:"user_id" gorm:"not null;unique"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Photo     string    `json:"photo" gorm:"not null"`
	Bod       time.Time `json:"bod" gorm:"not null;type:date"`
	Address   string    `json:"address" gorm:"not null"`
}
