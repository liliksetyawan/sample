package wardes_profile_image
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type WardesProfileImageDAO interface {
    
    InsertWardesProfileImage(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        repository.WardesProfileImageModel, 
        error, 
    )
    
    UpdateWardesProfileImage(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        error, 
    )
    
    DeleteWardesProfileImage(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        error, 
    )
    
    GetWardesProfileImage(
        ctx *context.ContextModel,
        uuid_key string,
    ) (
        error, 
    )
    
    GetListWardesProfileImage(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    CountListWardesProfileImage(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        int, 
        error, 
    )
    
    GetDataByUniqueData(
        ctx *context.ContextModel,
        dtoIn repository.WardesProfileImageModel,
    )(
        result repository.WardesProfileImageModel,
        err error,
    )

    GetWardesProfileImageForUpdate(
        ctx *context.ContextModel,
        tx *sql.Tx,
        uuid_key string,
    ) (
        repository.WardesProfileImageModel,
        error,
    )

}

