package user
                                                                                                                                                                                                                        

import (
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type UserDAO interface {
    
    GetListUser(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    InsertUser(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.UserModel,
    ) (
        repository.UserModel, 
        error, 
    )
    
    GetDataByUniqueData(
        ctx *context.ContextModel,
        dtoIn repository.UserModel,
    )(
        result repository.UserModel,
        err error,
    )

    GetUserForUpdate(
        ctx *context.ContextModel,
        tx *sql.Tx,
        uuid_key string,
    ) (
        repository.UserModel,
        error,
    )

}

