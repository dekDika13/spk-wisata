package adminRepository

import (
	adminDto "backend/dto/adminDto"
	"backend/model"
	"backend/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	//TODO Register Admin
	RegisterAdmin(payload adminDto.RegisterDTO) error
	LoginAdmin(payloads adminDto.LoginDTO) (model.User, error)

	// TODO Profile Admin
	GetProfileAdmin(userId uint) (adminDto.ProfileDTO, error)

	// TODO Destination Admin
	GetAllDestination() ([]adminDto.DestinationResponseDTO, error)
	GetDestinationById(id int) (adminDto.DestinationResponseDTO, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterDTO) error {

	bod, _ := time.Parse("2006-01-02", payloads.Bod)

	err := u.db.Transaction(func(tx *gorm.DB) error {
		// Simpan user
		user := &model.User{
			Username: payloads.Username,
			Password: payloads.Password, // hash dulu ya
			Email:    payloads.Email,
			Phone:    payloads.Phone,
			Role:     payloads.Role,
		}
		if err := tx.Create(&user).Error; err != nil {
			_ = utils.DeleteFromCloudinary(payloads.Idphoto)
			return err
		}

		// Simpan profile
		profile := &model.Profile{
			UserId:   user.UserID,
			FullName: payloads.FullName,
			Photo:    payloads.Photo,
			Bod:      bod,
			Address:  payloads.Address,
		}
		if err := tx.Create(&profile).Error; err != nil {

			_ = utils.DeleteFromCloudinary(payloads.Idphoto) // hapus foto
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads adminDto.LoginDTO) (model.User, error) {
	var admin model.User

	query := u.db.Where("username = ?", payloads.Username).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("username is incorrect")
	}

	return admin, nil
}
