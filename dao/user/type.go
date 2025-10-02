package user
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type UserDAO interface {
    
    InsertUser(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.UserModel,
    ) (
        repository.UserModel, 
        error, 
    )
    
    GetUser(
        ctx *context.ContextModel,
        uuid_key string,
    ) (
        repository.UserModel, 
        error, 
    )
    
    UpdateUser(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.UserModel,
    ) (
        error, 
    )
    
    DeleteUser(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.UserModel,
    ) (
        error, 
    )
    
    GetListUser(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    CountListUser(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        int, 
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

