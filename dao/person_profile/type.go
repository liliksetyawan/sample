package person_profile
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

type PersonProfileDAO interface {
    
    InsertPersonProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.PersonProfileModel,
    ) (
        repository.PersonProfileModel, 
        error, 
    )
    
    UpdatePersonProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.PersonProfileModel,
    ) (
        error, 
    )
    
    DeletePersonProfile(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.PersonProfileModel,
    ) (
        error, 
    )
    
    GetPersonProfile(
        ctx *context.ContextModel,
        uuid_key string,
    ) (
        repository.PersonProfileModel, 
        error, 
    )
    
    GetListPersonProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []repository.PersonProfileModel, 
        error, 
    )
    
    CountListPersonProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        int, 
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

