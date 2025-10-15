package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dto/in"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexcommon/services"
)

type WardesProfileService interface {
    
    InsertWardesProfile(
        ctx *context.ContextModel,
        param model.URLParam,
        dto interface{},
    ) (
        header map[string]string,
        output interface{},
        err error,
    )
    
    GetWardesProfileList(
        ctx *context.ContextModel,
        searchParam []model.SearchParam,
        dto in.GetListRequest,
    ) (
        header map[string]string,
        output interface{},
        err error,
    )
    
    GetWardesProfileCount(
        ctx *context.ContextModel,
        param model.URLParam,
        dto interface{},
    ) (
        header map[string]string,
        output interface{},
        err error,
    )
    
    UpdateWardesProfile(
        ctx *context.ContextModel,
        param model.URLParam,
        dto interface{},
    ) (
        header map[string]string,
        output interface{},
        err error,
    )
    
    GetDTO() interface{}
    GetListScope() []string
    GetListValidLimit() []int
    GetListValidOrderBy() []string
    GetListValidSearch() []string
    GetDefaultOperator() services.DefaultOperators
    GetMultipartDTO() interface{}
}