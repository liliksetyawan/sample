package wardes_profile_image
                                                                                                                                                                                                                                                                                                                                            

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type WardesProfileImageDAO interface {
    
    UpdateWardesProfileImage(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        error, 
    )
    
    InsertWardesProfileImage(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        repository.WardesProfileImageModel, 
        error, 
    )
    
    GetWardesProfileImageByUserID(
        ctx *context.ContextModel,
        id int64,
    ) (
        repository.WardesProfileImageModel, 
        error, 
    )
    
    CountWardesProfileImage(
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

