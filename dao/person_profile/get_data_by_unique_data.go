package person_profile

import (
    "database/sql"
    "fmt"
    "strings"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *person_profilePostgresqlSQLDAO) GetDataByUniqueData(
    ctx *context.ContextModel,
    dtoIn repository.PersonProfileModel,
) (
    result repository.PersonProfileModel,
    err error,
) {
    query := fmt.Sprintf(`
        SELECT id, uuid_key, first_name, last_name
        FROM %s
        WHERE nik = $1 OR email = $2
    `, dao.GetDBTable(ctx, "person_profile"))

    stmt, err := d.db.Prepare(query)
    if err != nil {
        return repository.PersonProfileModel{}, err
    }

    err = stmt.QueryRow(dtoIn.Nik.String, dtoIn.Email.String).Scan(
        &result.ID, &result.UuidKey, &result.FirstName, &result.LastName,
    )
    return result, err
}
