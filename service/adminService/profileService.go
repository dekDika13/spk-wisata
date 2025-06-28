package adminService

import adminDto "backend/dto/adminDto"

func (s *adminService) GetProfileAdmin(userId uint) (adminDto.ProfileDTO, error) {
	profile, err := s.adminRepository.GetProfileAdmin(userId)
	if err != nil {
		return adminDto.ProfileDTO{}, err
	}
	return profile, nil
}
