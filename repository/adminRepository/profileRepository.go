package adminRepository

import (
	adminDto "backend/dto/adminDto"
	"backend/model"
)

func (u *adminRepository) GetProfileAdmin(userId uint) (adminDto.ProfileDTO, error) {
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
