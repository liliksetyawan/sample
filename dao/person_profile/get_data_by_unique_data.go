package person_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *person_profilePostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.PersonProfileModel,
)( 
    result repository.PersonProfileModel,
    err error,
){
    query := fmt.Sprintf(`
        SELECT id, uuid_key
        FROM %s
        WHERE email = $1
    `, dao.GetDBTable(ctx, "person_profile"))

    err = d.db.QueryRow(query, dtoIn.Email.String).Scan(&result.Id, &result.UuidKey)
    return
}
