package wardes_profile

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type WardesProfileDAO interface {
    
    GetListWardesProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    GetDataByUniqueData(
        ctx *context.ContextModel,
        dtoIn repository.WardesProfileModel,
    )(
        result repository.WardesProfileModel,
        err error,
    )

    GetWardesProfileForUpdate(
        ctx *context.ContextModel,
        tx *sql.Tx,
        uuid_key string,
    ) (
        repository.WardesProfileModel,
        error,
    )

}
