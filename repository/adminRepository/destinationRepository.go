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
	return view, nil
}
func (u *adminRepository) GetDestinationById(id int) (adminDto.DestinationResponseDTO, error) {
	view := adminDto.DestinationResponseDTO{}

	if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).First(&view).Error; err != nil {
		return view, err
	}
	return view, nil
}
