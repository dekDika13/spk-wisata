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
	Toilet      int                     `form:"toilet" validate:"required"`
	Parking     int                     `form:"parking" validate:"required"`
	Restarea    int                     `form:"restarea" validate:"required"`
	Restaurant  int                     `form:"restaurant" validate:"required"`
	Price       int                     `form:"price" validate:"required"`
	Address     string                  `form:"address" validate:"required"`
	Location    string                  `form:"location" validate:"required"`
}
type DestinationImageDTO struct {
	CoverUrl    string   `json:"cover_url"`
	ImageUrl []string `json:"image_url"`
}

type DestinationResponseDTO struct {
	DestinationId    uint            `json:"destination_id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Cover            string          `json:"cover"`
	Galery           []string        `json:"gallery"`
	Toilet           int             `json:"toilet"`
	Parking          int             `json:"parking"`
	Restarea         int             `json:"restarea"`
	Restaurant       int             `json:"restaurant"`
	Price            int             `json:"price"`
	Rating           decimal.Decimal `json:"rating"`
	AssessmentResult decimal.Decimal `json:"assessment_result"`
	Address          string          `json:"address"`
	Location         string          `json:"location"`
}
