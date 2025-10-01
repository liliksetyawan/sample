package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/services"
    "github.com/nexsoft-git/nexcommon/bundles"
    "nexsoft.co.id/example/dto/in"
    wardes_profileDAO "nexsoft.co.id/example/dao/wardes_profile"
)

func NewWardesProfileService(
    wardes_profileDAO wardes_profileDAO.WardesProfileDAO,
) WardesProfileService {

    operator := make(services.DefaultOperators)

    operator["code"] = model.DefaultOperator{
        DataType: "char",
        Operator: []string{ "lk", "eq" },
    }

    operator["name"] = model.DefaultOperator{
        DataType: "char",
        Operator: []string{ "lk", "eq" },
    }

    operator["nexchief_account_id"] = model.DefaultOperator{
        DataType: "number",
        Operator: []string{ "lk", "eq" },
    }

    service := &wardes_profileService{ 
        wardes_profileDAO: wardes_profileDAO,
        searchOperator: operator,
    }

    return service
}

type wardes_profileService struct { 
    wardes_profileDAO wardes_profileDAO.WardesProfileDAO
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