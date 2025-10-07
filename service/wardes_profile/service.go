package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/services"
	wardes_profileDAO "nexsoft.co.id/example/dao/wardes_profile"
	wardes_profile_imageDAO "nexsoft.co.id/example/dao/wardes_profile_image"
	"nexsoft.co.id/example/dto/in"
)

func NewWardesProfileService(
	wardes_profile_imageDAO wardes_profile_imageDAO.WardesProfileImageDAO,
	wardes_profileDAO wardes_profileDAO.WardesProfileDAO,
) WardesProfileService {

	operator := make(services.DefaultOperators)

	service := &wardes_profileService{
		wardes_profile_imageDAO: wardes_profile_imageDAO,
		wardes_profileDAO:       wardes_profileDAO,
		searchOperator:          operator,
	}

	return service
}

type wardes_profileService struct {
	wardes_profile_imageDAO wardes_profile_imageDAO.WardesProfileImageDAO
	wardes_profileDAO       wardes_profileDAO.WardesProfileDAO
	searchOperator          services.DefaultOperators
}

func (s *wardes_profileService) GetDTO() interface{} {
	return &in.WardesProfileDTOIn{}
}

func (s *wardes_profileService) GetMultipartDTO() interface{} {
	return nil
}

func (s *wardes_profileService) GetListScope() []string {
	return []string{"read", "write"}
}

func (s *wardes_profileService) parseDTO(dtoIn interface{}) *in.WardesProfileDTOIn {
	return dtoIn.(*in.WardesProfileDTOIn)
}

func (s *wardes_profileService) GetListValidLimit() []int {
	return []int{10, 20, 30, 50, 100}
}

func (s *wardes_profileService) GetListValidOrderBy() []string {
	return []string{"id"}
}

func (s *wardes_profileService) GetListValidSearch() []string {
	return s.searchOperator.GetListValidSearch()
}

func (s *wardes_profileService) GetDefaultOperator() services.DefaultOperators {
	return s.searchOperator
}
