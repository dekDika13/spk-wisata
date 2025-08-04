package tourisdto

type CreateReviewDTO struct {
	DestinationId uint   `json:"destination_id " validate:"required"`
	UserId        uint   `json:"user_id" validate:"required"`
	ReviewDetail  string `json:"review_detail" validate:"required"`
	RatingC1      int    `json:"rating_c1" validate:"required"`
	RatingC2      int    `json:"rating_c2" validate:"required"`
	RatingC4      int    `json:"rating_c4" validate:"required"`
	RatingC6      int    `json:"rating_c6" validate:"required"`
	RatingC7      int    `json:"rating_c7" validate:"required"`
}

type ReviewResponseDTO struct {
	ReviewId      uint   `json:"review_id"`
	DestinationId uint   `json:"destination_id"`
	UserId        uint   `json:"user_id"`
	ReviewDetail  string `json:"review_detail"`
	RatingC1      int    `json:"rating_c1"`
	RatingC2      int    `json:"rating_c2"`
	RatingC4      int    `json:"rating_c4"`
	RatingC6      int    `json:"rating_c6"`
	RatingC7      int    `json:"rating_c7"`
}
