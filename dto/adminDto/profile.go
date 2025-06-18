package admindto

import "mime/multipart"

type ProfileDTO struct {
	UserId   uint   `json:"user_id" `
	FullName string `json:"full_name" `
	Photo    string `json:"photo" `
	Bod      string `json:"bod" `
	Address  string `json:"address" `
	Role     uint   `json:"role" `
	Username string `json:"username" `
	Email    string `json:"email" `
	Phone    string `json:"phone" `
}
type InsertProfileDTO struct {
	UserId   uint                  `json:"user_id" validate:"required"`
	FullName string                `json:"full_name" `
	Photo    *multipart.FileHeader `form:"photo"`
	Bod      string                `json:"bod" `
	Address  string                `json:"address" `
	Username string                `json:"username"`
	Password string                `json:"password" validate:"required"`
}
