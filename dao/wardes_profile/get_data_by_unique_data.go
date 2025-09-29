package wardes_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *wardes_profilePostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.WardesProfileModel,
)( 
    result repository.WardesProfileModel,
    err error,
){
    query := fmt.Sprintf(`
        SELECT id, uuid_key
        FROM %s
        WHERE username = $1 OR email = $2
    `, dao.GetDBTable(ctx, "wardes_profile"))

    err = d.db.QueryRow(query, dtoIn.Username.String, dtoIn.Email.String).Scan(
        &result.ID, &result.UUIDKey,
    )
    return
}
