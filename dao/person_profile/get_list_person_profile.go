package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get a list of person profile records with search and pagination
func (d *person_profilePostgresqlSQLDAO) GetListPersonProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, created_by, created_client, created_at, 
               updated_by, updated_client, updated_at
        FROM %s
    `, dao.GetDBTable(ctx, "person_profile"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.PersonProfileModel
            err := rows.Scan(
                &temp.UUIDKey, &temp.CreatedBy, &temp.CreatedClient, &temp.CreatedAt,
                &temp.UpdatedBy, &temp.UpdatedClient, &temp.UpdatedAt,
            )
            return temp, err
        },
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}

