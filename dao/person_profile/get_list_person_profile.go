
package person_profile

import (
    "database/sql"
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get list of person profiles with pagination and search capabilities
func (d *person_profilePostgresqlSQLDAO) GetListPersonProfile(
    ctx *context.ContextModel,
    dtoIn in.GetListRequest,
    searchParams []model.SearchParam,
) (
    []interface{},
    error,
) {
    query := fmt.Sprintf(`
        SELECT person_profile_id, title_id, title_name, first_name, last_name, sex
        FROM %s
    `, dao.GetDBTable(ctx, "person_profile"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.PersonProfileModel
            err := rows.Scan(
                &temp.PersonProfileID,
                &temp.TitleID,
                &temp.TitleName,
                &temp.FirstName,
                &temp.LastName,
                &temp.Sex,
            )
            return temp, err
        },
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}

