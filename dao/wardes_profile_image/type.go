package wardes_profile_image
                                                                                                                                                                                

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type WardesProfileImageDAO interface {
    
    InsertWardesProfileImageWithTemporary(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.WardesProfileImageModel,
    ) (
        repository.WardesProfileImageModel, 
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

