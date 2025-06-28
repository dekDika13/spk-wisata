package tourisdto

import (
	"time"
)

type ReviewResponseDTO struct {
	ReviewId      uint      `json:"review_id"`
	UserId        uint      `json:"user_id"`
	DestinationId uint      `json:"destination_id"`
	Name          string    `json:"name"`
	Image1        string    `json:"image1"`
	CreatedAt     time.Time `json:"created_at"`
	Rating        int       `json:"rating"`
	ReviewDetail  string    `json:"review_detail"`
}
