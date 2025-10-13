package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/services"
    "github.com/nexsoft-git/nexcommon/bundles"
    "nexsoft.co.id/example/dto/in"
    WardesProfileImageDAO "nexsoft.co.id/example/dao/wardes_profile_image"
    WardesProfileDAO "nexsoft.co.id/example/dao/wardes_profile"
    UserDAO "nexsoft.co.id/example/dao/user"
)

func NewWardesProfileService(
    wardes_profile_imageDAO wardes_profile_imageDAO.WardesProfileImageDAO,
    bundles bundles.Bundles,
    wardes_profileDAO wardes_profileDAO.WardesProfileDAO,
    userDAO userDAO.UserDAO,
) WardesProfileService {

    operator := make(services.DefaultOperators)

    operator["wp.username"] = model.DefaultOperator{
        DataType: "char",
        Operator: []string{ "lk", "eq", "in" },
    }

    service := &wardes_profileService{ 
        wardes_profile_imageDAO: wardes_profile_imageDAO,
        bundles: bundles,
        wardes_profileDAO: wardes_profileDAO,
        userDAO: userDAO,
        searchOperator: operator,
    }

    return service
}

type wardes_profileService struct { 
    wardes_profile_imageDAO wardes_profile_imageDAO.WardesProfileImageDAO
    bundles bundles.Bundles
    wardes_profileDAO wardes_profileDAO.WardesProfileDAO
    userDAO userDAO.UserDAO
    searchOperator services.DefaultOperators
}

func (s *wardes_profileService) GetDTO() interface{} {
    return &in.WardesProfileDTOIn{}
}

func (s *wardes_profileService) GetMultipartDTO() interface{} {
    return nil
}

func (s *wardes_profileService) GetListScope() []string {
    return []string{ "read", "write" }
}

func (s *wardes_profileService) parseDTO(dtoIn interface{}) *in.WardesProfileDTOIn {
    return dtoIn.(*in.WardesProfileDTOIn)
}

func (s *wardes_profileService) GetListValidLimit() []int {
    return []int{ 10, 20, 30, 50, 100 }
}

func (s *wardes_profileService) GetListValidOrderBy() []string {
    return []string{ "id" }
}

func (s *wardes_profileService) GetListValidSearch() []string {
    return s.searchOperator.GetListValidSearch()
}

func (s *wardes_profileService) GetDefaultOperator() services.DefaultOperators {
    return s.searchOperator
}