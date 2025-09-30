package person_profile
                                                                                                                                                                                                                                                                        

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type PersonProfileDAO interface {
    
    GetDetailPersonProfile(
        ctx *context.ContextModel,
        id int64,
    ) (
        repository.PersonProfileModel, 
        error, 
    )
    
    InsertPersonProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.PersonProfileModel,
    ) (
        repository.PersonProfileModel, 
        error, 
    )
    
    GetListPersonProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
        error, 
    )
    
    GetDataByUniqueData(
        ctx *context.ContextModel,
        dtoIn repository.PersonProfileModel,
    )(
        result repository.PersonProfileModel,
        err error,
    )

    GetPersonProfileForUpdate(
        ctx *context.ContextModel,
        tx *sql.Tx,
        uuid_key string,
    ) (
        repository.PersonProfileModel,
        error,
    )

}

