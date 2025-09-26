package wardes_profile
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

import (
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type WardesProfileDAO interface {
    
    CountWardesProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        int, 
        error, 
    )
    
    InsertWardesProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileModel,
    ) (
        repository.WardesProfileModel, 
        error, 
    )
    
    GetListWardesProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    GetWardesProfile(
        ctx *context.ContextModel,
        uuid_key string,
    ) (
        repository.WardesProfileModel, 
        error, 
    )
    
    UpdateWardesProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileModel,
    ) (
        error, 
    )
    
    DeleteWardesProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileModel,
    ) (
        error, 
    )
    
    GetWardesProfileByCode(
        ctx *context.ContextModel,
        uuid_key string,
    ) (
        repository.WardesProfileModel, 
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

