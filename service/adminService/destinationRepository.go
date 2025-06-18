package adminService

import adminDto "backend/dto/adminDto"

func (s *adminService) GetAllDestination() ([]adminDto.DestinationResponseDTO, error) {
	return s.adminRepository.GetAllDestination()
}
