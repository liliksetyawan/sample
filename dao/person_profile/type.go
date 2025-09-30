package person_profile
                                                                                                                                                                                                                        

import (
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type PersonProfileDAO interface {
    
    GetListPersonProfile(
        ctx *context.ContextModel,
        dtoIn in.GetListRequest,
        searchParams []model.SearchParam,
    ) (
        []interface{}, 
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

