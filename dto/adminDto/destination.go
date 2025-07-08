package admindto

import (
	"mime/multipart"

	"github.com/shopspring/decimal"
)

type DestinationCreateDTO struct {
	Name        string                  `form:"name" validate:"required"`
	Description string                  `form:"description" validate:"required"`
	Cover       *multipart.FileHeader   `form:"cover"`
	Images      []*multipart.FileHeader `form:"images"`
	Price       int                     `form:"price" validate:"required"`
	Address     string                  `form:"address" validate:"required"`
	Location    string                  `form:"location" validate:"required"`
}
type DestinationInsertDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Image1   [3]*multipart.FileHeader `form:"image1"`
	Price    int                      `json:"price"`
	Address  string                   `json:"address"`
	Location string                   `json:"location"`
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

type DestinationUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
	Image3      string `json:"image3"`
	Image4      string `json:"image4"`
	Price       int    `json:"price"`
	Address     string `json:"address"`
	Location    string `json:"location"`
}
