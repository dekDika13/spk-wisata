package adminRepository

import (
	adminDto "backend/dto/adminDto"
	"backend/model"
)

func (u *adminRepository) GetAllDestination() ([]adminDto.DestinationResponseDTO, error) {
	view := []adminDto.DestinationResponseDTO{}

	if err := u.db.Model(&model.Destination{}).Select("destinations.*").Find(&view).Error; err != nil {
		return nil, err
	}
	return view,nil
}
