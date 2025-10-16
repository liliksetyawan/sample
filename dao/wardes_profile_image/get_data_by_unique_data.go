
package wardes_profile_image

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *wardes_profile_imagePostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.WardesProfileImageModel,
)(
    result repository.WardesProfileImageModel,
    err error,
){
    //No Unique Constraint on the database, return empty result
    return repository.WardesProfileImageModel{}, nil
}

