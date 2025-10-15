package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/bundles"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexcommon/services"
	userDAO "nexsoft.co.id/example/dao/user"
	wardesProfileDAO "nexsoft.co.id/example/dao/wardes_profile"
	wardesProfileImageDAO "nexsoft.co.id/example/dao/wardes_profile_image"
	"nexsoft.co.id/example/dto/in"
)

func NewWardesProfileService(
	userDAO userDAO.UserDAO,
	bundles bundles.Bundles,
	wardesProfileImageDAO wardesProfileImageDAO.WardesProfileImageDAO,
	wardesProfileDAO wardesProfileDAO.WardesProfileDAO,
) WardesProfileService {

	operator := make(services.DefaultOperators)

	operator["wp.username"] = model.DefaultOperator{
		DataType: "char",
		Operator: []string{"lk", "eq", "in"},
	}

	service := &wardes_profileService{
		userDAO:               userDAO,
		bundles:               bundles,
		wardesProfileImageDAO: wardesProfileImageDAO,
		wardesProfileDAO:      wardesProfileDAO,
		searchOperator:        operator,
	}

	return service
}

type wardes_profileService struct {
	userDAO               userDAO.UserDAO
	bundles               bundles.Bundles
	wardesProfileImageDAO wardesProfileImageDAO.WardesProfileImageDAO
	wardesProfileDAO      wardesProfileDAO.WardesProfileDAO
	searchOperator        services.DefaultOperators
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
