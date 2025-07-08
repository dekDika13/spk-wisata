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

func (u *adminRepository) CreateDestination(payload  adminDto.DestinationCreateDTO) error {
	destination := model.Destination{
		Name:        payload.Name,
		Description: payload.Description,
		
		Price:      payload.Price,
		Address:    payload.Address,
		Location:   payload.Location,
	}

	if err := u.db.Create(&destination).Error; err != nil {
		return err
	}
	return nil
}
func (u *adminRepository) UpdateDestination(id int, payload adminDto.DestinationUpdateDTO) error {
	destination := model.Destination{
		Name:        payload.Name,
		Description: payload.Description,
		Image1:     payload.Image1,
		Image2:     payload.Image2,
		Image3:     payload.Image3,
		Image4:     payload.Image4,
		Price:      payload.Price,
		Address:    payload.Address,
		Location:   payload.Location,
	}

	if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).Updates(destination).Error; err != nil {
		return err
	}
	return nil
}
func (u *adminRepository) DeleteDestination(id int) error {
	if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).Delete(&model.Destination{}).Error; err != nil {
		return err
	}
	return nil
}
