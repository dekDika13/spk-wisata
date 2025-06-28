package admindto

import "github.com/shopspring/decimal"

type DestinationCreateDTO struct {
	Name             string          `json:"name" validate:"required"`
	Description      string          `json:"description" validate:"required"`
	Image1           string          `json:"image1" validate:"required"`
	Image2           string          `json:"image2" validate:"required"`
	Image3           string          `json:"image3" validate:"required"`
	Image4           string          `json:"image4" validate:"required"`
	Price            int             `json:"price" validate:"required"`
	Address          string          `json:"address" validate:"required"`
	Location         string          `json:"location" validate:"required"`
}

type DestinationResponseDTO struct {
	DestinationId    uint            `json:"destination_id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Image1           string          `json:"image1"`
	Image2           string          `json:"image2"`
	Image3           string          `json:"image3"`
	Image4           string          `json:"image4"`
	Price            int             `json:"price"`
	AverageRating    decimal.Decimal `json:"average_rating"`
	AssessmentResult decimal.Decimal `json:"assessment_result"`
	Address          string          `json:"address"`
	Location         string          `json:"location"`
}
