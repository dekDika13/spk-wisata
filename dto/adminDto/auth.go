package admindto

import "mime/multipart"

type RegisterDTO struct {
	Role     uint   `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Photo    string `json:"photo" validate:"required"`
	Idphoto  string `json:"idphoto" validate:"required"`
	Bod      string `json:"bod"  validate:"required"`
	Address  string `json:"address" validate:"required"`
}
type RegisterInsertDTO struct {
	Role     uint                  `json:"role" validate:"required"`
	Username string                `json:"username" validate:"required"`
	Password string                `json:"password" validate:"required"`
	Email    string                `json:"email" validate:"required,email"`
	Phone    string                `json:"phone" validate:"required"`
	FullName string                `json:"full_name" validate:"required"`
	Photo    *multipart.FileHeader `form:"photo"`
	Bod      string                `json:"bod"  validate:"required"`
	Address  string                `json:"address" validate:"required"`
}

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginJWT struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
