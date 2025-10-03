package parameter
                                                                                                            

import (
    "github.com/nexsoft-git/nexcommon/context"
    "database/sql"
    "nexsoft.co.id/example/repository"
)

type ParameterDAO interface {
    
    InsertParameter(
        ctx *context.ContextModel,
        tx *sql.Tx,
        dtoIn repository.ParameterModel,
    ) (
        repository.ParameterModel, 
        error, 
    )
    
    GetDataByUniqueData(
        ctx *context.ContextModel,
        dtoIn repository.ParameterModel,
    )(
        result repository.ParameterModel,
        err error,
    )

    GetParameterForUpdate(
        ctx *context.ContextModel,
        tx *sql.Tx,
        uuid_key string,
    ) (
        repository.ParameterModel,
        error,
    )

}

