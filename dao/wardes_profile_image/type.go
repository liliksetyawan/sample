package wardes_profile_image
                                                                                                                                                                                                                                                                        

import (
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type WardesProfileImageDAO interface {
    
    CountWardesProfileImage(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        int, 
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

