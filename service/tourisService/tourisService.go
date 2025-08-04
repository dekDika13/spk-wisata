package touristservice

import (
	adminDto "backend/dto/adminDto"
	tourisDto "backend/dto/tourisDto"
	tourisRepository "backend/repository/tourisRepository"
)

type TourisService interface {

	// TODO Profile Admin
	GetProfileTouris(userId uint) (adminDto.ProfileDTO, error)
	GetAllReviewTouris(userId uint) ([]tourisDto.ReviewResponseDTO, error)
	CreateReviewTouris( payloads tourisDto.CreateReviewDTO) error
	UpdateReviewTouris(reviewId uint) error
	DeleteReviewTouris(reviewId uint) error
}
type tourisService struct {
	tourisRepository tourisRepository.TourisRepository
}

func NewTourisService(tourisRepo tourisRepository.TourisRepository) *tourisService {
	return &tourisService{
		tourisRepository: tourisRepo,
	}
}

func (s *tourisService) GetProfileTouris(userId uint) (adminDto.ProfileDTO, error) {
	profile, err := s.tourisRepository.GetProfileTouris(userId)
	if err != nil {
		return adminDto.ProfileDTO{}, err
	}
	return profile, nil
}

// TODO Get All Review Touris
func (s *tourisService) GetAllReviewTouris(userId uint) ([]tourisDto.ReviewResponseDTO, error) {
	reviews, err := s.tourisRepository.GetAllReviewTouris(userId)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *tourisService) CreateReviewTouris( payloads tourisDto.CreateReviewDTO) error {
	if err := s.tourisRepository.CreateReviewTouris(payloads); err != nil {
		return err
	}
	return nil
}

func (s *tourisService) UpdateReviewTouris(reviewId uint) error {
	if err := s.tourisRepository.UpdateReviewTouris(reviewId); err != nil {
		return err
	}
	return nil
}
func (s *tourisService) DeleteReviewTouris(reviewId uint) error {
	if err := s.tourisRepository.DeleteReviewTouris(reviewId); err != nil {
		return err
	}
	return nil
}
