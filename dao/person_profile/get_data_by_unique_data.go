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
) {
    query := fmt.Sprintf(`
        SELECT id, uuid_key
        FROM %s
        WHERE first_name = $1 AND last_name = $2 AND phone = $3
    `, dao.GetDBTable(ctx, "person_profile"))

    stmt, err := d.db.Prepare(query)
    if err != nil {
        return repository.PersonProfileModel{}, err
    }

    err = stmt.QueryRow(
        dtoIn.FirstName.String, dtoIn.LastName.String, dtoIn.Phone.String,
    ).Scan(&result.Id, &result.UuidKey)
    
    return result, err
}
