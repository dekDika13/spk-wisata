package adminRepository

import (
	adminDto "backend/dto/adminDto"
	tourisDto "backend/dto/tourisDto"
	"backend/model"

	"gorm.io/gorm"
)

type TourisRepository interface {

	// TODO Profile Admin
	GetProfileTouris(userId uint) (adminDto.ProfileDTO, error)
	GetAllReviewTouris(userId uint) ([]tourisDto.ReviewResponseDTO, error)
	CreateReviewTouris(reviewId uint ) error
	UpdateReviewTouris(reviewId uint) error
	DeleteReviewTouris(reviewId uint) error
}

type tourisRepository struct {
	db *gorm.DB
}

func NewTourisRepository(db *gorm.DB) *tourisRepository {
	return &tourisRepository{db}
}

func (u *tourisRepository) GetProfileTouris(userId uint) (adminDto.ProfileDTO, error) {
	var profile adminDto.ProfileDTO

	if err := u.db.Model(&model.Profile{}).
		Select("profiles.user_id, users.username, profiles.full_name, profiles.photo, profiles.bod, profiles.address,users.role,users.email,users.phone").
		Joins("JOIN users ON users.user_id = profiles.user_id").
		Where("profiles.user_id = ?", userId).
		First(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}

func (u *tourisRepository) GetAllReviewTouris(userId uint) ([]tourisDto.ReviewResponseDTO, error) {
	var reviews []tourisDto.ReviewResponseDTO

	if err := u.db.Model(&model.Review{}).
		Select("reviews.review_id, reviews.user_id, reviews.destination_id, reviews.rating, reviews.review_detail, reviews.created_at, destinations.image1, destinations.name").
		Joins("JOIN users ON users.user_id = reviews.user_id").
		Joins("JOIN destinations ON destinations.destination_id = reviews.destination_id").
		Where("reviews.user_id = ?", userId).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (u *tourisRepository) CreateReviewTouris(reviewId uint) error {
	if err := u.db.Model(&model.Review{}).
		Where("review_id = ?", reviewId).
		Update("is_active", true).Error; err != nil {
		return err
	}

	return nil
}
func (u *tourisRepository) UpdateReviewTouris(reviewId uint) error {
	if err := u.db.Model(&model.Review{}).
		Where("review_id = ?", reviewId).
		Update("is_active", false).Error; err != nil {
		return err
	}

	return nil
}
func (u *tourisRepository) DeleteReviewTouris(reviewId uint) error {
	if err := u.db.Model(&model.Review{}).
		Where("review_id = ?", reviewId).
		Update("is_active", false).Error; err != nil {
		return err
	}

	return nil
}

