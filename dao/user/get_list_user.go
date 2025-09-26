package user

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will retrieve a list of users from the database based on given parameters.
func (d *userPostgresqlSQLDAO) GetListUser(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, created_by, created_client, created_at, updated_by, updated_client, updated_at
        FROM %s
    `, dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.UserModel
            err := rows.Scan(
                &temp.UUIDKey, &temp.CreatedBy, &temp.CreatedClient, &temp.CreatedAt,
                &temp.UpdatedBy, &temp.UpdatedClient, &temp.UpdatedAt,
            )
            return temp, err
        },
    ).FieldConverter(dao.NewParaFieldConverter().
        Add("id", "id").
        Add("username", "username").
        Add("email", "email").
        Add("created_at", "created_at"),
    )

    return d.GetListDataWithDefaultMustCheck(param)
}
