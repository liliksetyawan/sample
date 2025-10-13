package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/services"
    "github.com/nexsoft-git/nexcommon/bundles"
    "nexsoft.co.id/example/dto/in"
    userDAO "nexsoft.co.id/example/dao/user"
    wardesProfileDAO "nexsoft.co.id/example/dao/wardes_profile"
    wardesProfileImageDAO "nexsoft.co.id/example/dao/wardes_profile_image"
)

func NewWardesProfileService(
    bundles bundles.Bundles,
    userdao userDAO.UserDAO,
    wardesProfiledao wardesProfileDAO.WardesProfileDAO,
    wardesProfileImagedao wardesProfileImageDAO.WardesProfileImageDAO,
) WardesProfileService {

    operator := make(services.DefaultOperators)

    operator["wp.username"] = model.DefaultOperator{
        DataType: "char",
        Operator: []string{ "lk", "eq", "in" },
    }

    service := &wardes_profileService{ 
        bundles: bundles,
        userDAO: userDAO,
        wardes_profileDAO: wardes_profileDAO,
        wardes_profile_imageDAO: wardes_profile_imageDAO,
        searchOperator: operator,
    }

    return service
}

type wardes_profileService struct { 
    bundles bundles.Bundles
    userDAO userDAO.UserDAO
    wardes_profileDAO wardes_profileDAO.WardesProfileDAO
    wardes_profile_imageDAO wardes_profile_imageDAO.WardesProfileImageDAO
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