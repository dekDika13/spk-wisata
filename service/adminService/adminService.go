package adminService

import (
	adminDto "backend/dto/adminDto"
	"backend/middleware"
	adminRepository "backend/repository/adminRepository"
	"backend/utils"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	// TODO Register Admin
	RegisterAdmin(payloads adminDto.RegisterInsertDTO) error
	LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error)

	// TODO Destination Admin
	GetAllDestination() ([]adminDto.DestinationResponseDTO, error)
	GetDestinationById(id int) (adminDto.DestinationResponseDTO, error)
	CreateDestination(payload adminDto.DestinationCreateDTO) error
	UpdateDestination(id int, payload adminDto.DestinationUpdateDTO) error
	DeleteDestination(id int) error

	// TODO Profile Admin
	GetProfileAdmin(userId uint) (adminDto.ProfileDTO, error)
}
type adminService struct {
	adminRepository adminRepository.AdminRepository
}

func NewAdminService(adminRepo adminRepository.AdminRepository) *adminService {
	return &adminService{
		adminRepository: adminRepo,
	}
}

// TODO Register Admin
func (s *adminService) RegisterAdmin(payloads adminDto.RegisterInsertDTO) error {
	photoURL := ""
	idphoto := ""
	if payloads.Photo != nil {
		filename := uuid.New().String()
		file, err := payloads.Photo.Open()
		if err != nil {
			return err // failed to open file
		}
		defer file.Close()
		url,id, err := utils.UploadToCloudinary(file, filename)
		if err != nil {
			return err // bisa log juga
		}
		photoURL = url
		idphoto = id
	}
	pas := adminDto.RegisterDTO{
		Role:     payloads.Role,
		Username: payloads.Username,
		Password: payloads.Password,
		Email:    payloads.Email,
		Phone:    payloads.Phone,
		FullName: payloads.FullName,
		Photo:    photoURL,
		Bod:      payloads.Bod,
		Address:  payloads.Address,
		Idphoto: idphoto,
	}

	pw, err := utils.HashBcrypt(pas.Password)

	if err != nil {
		return err
	}

	pas.Password = pw

	return s.adminRepository.RegisterAdmin(pas)
}

// TODO LOGIN ADMIN
func (s *adminService) LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error) {
	var temp adminDto.LoginJWT

	res, err := s.adminRepository.LoginAdmin(payloads)

	if errh := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(payloads.Password)); errh != nil {
		return temp, errors.New("username or password incorrect")
	}

	token, errt := middleware.CreateToken(res.UserID, res.Role, res.Username)

	temp = adminDto.LoginJWT{
		Username: res.Username,
		Token:    token,
	}

	if err != nil {
		return temp, err
	}

	if errt != nil {
		return temp, errt
	}

	return temp, nil
}
