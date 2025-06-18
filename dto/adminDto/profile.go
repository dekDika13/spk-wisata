package admindto

import "mime/multipart"

type ProfileDTO struct {
	UserId   uint   `json:"user_id" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Photo    string `json:"photo" validate:"required"`
	Bod      string `json:"bod" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
type InsertProfileDTO struct {
	UserId   uint                  `json:"user_id" validate:"required"`
	FullName string                `json:"full_name" validate:"required"`
	Photo    *multipart.FileHeader `form:"photo"`
	Bod      string                `json:"bod" validate:"required"`
	Address  string                `json:"address" validate:"required"`
}
