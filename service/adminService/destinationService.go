package adminService

import adminDto "backend/dto/adminDto"

func (s *adminService) GetAllDestination() ([]adminDto.DestinationResponseDTO, error) {
	return s.adminRepository.GetAllDestination()
}
func (s *adminService) GetDestinationById(id int) (adminDto.DestinationResponseDTO, error) {
	return s.adminRepository.GetDestinationById(id)
}
func (s *adminService) CreateDestination(payload adminDto.DestinationCreateDTO) error {

	
	
	return s.adminRepository.CreateDestination(payload)
}
func (s *adminService) UpdateDestination(id int, payload adminDto.DestinationUpdateDTO) error {
	return s.adminRepository.UpdateDestination(id, payload)
}
func (s *adminService) DeleteDestination(id int) error {
	return s.adminRepository.DeleteDestination(id)
}

