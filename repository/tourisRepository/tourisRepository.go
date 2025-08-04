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
	CreateReviewTouris(payloads tourisDto.CreateReviewDTO ) error
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
		Select("*").
		Where("reviews.user_id = ?", userId).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (u *tourisRepository) CreateReviewTouris(payloads tourisDto.CreateReviewDTO) error {
	review := &model.Review{
		DestinationId: payloads.DestinationId,
		UserId:        payloads.UserId,
		ReviewDetail:  payloads.ReviewDetail,
		RatingC1:      payloads.RatingC1,
		RatingC2:      payloads.RatingC2,
		RatingC4:      payloads.RatingC4,
		RatingC6:      payloads.RatingC6,
		RatingC7:      payloads.RatingC7,
	}
	if err := u.db.Create(review).Error; err != nil {			
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

