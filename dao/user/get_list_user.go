package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Retrieves a list of users
func (d *userPostgresqlSQLDAO) GetListUser(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT u.id, u.client_id, wp.id, wp.username, wp.email,
            u.created_at, u.created_by, u.created_client,
            u.updated_at, u.updated_by, u.updated_client
        FROM %s u
        JOIN %s wp ON u.person_profile_id = wp.id
    `, dao.GetDBTable(ctx, "user"), dao.GetDBTable(ctx, "wardes_profile"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.WardesProfileWithUserModel
            err := rows.Scan(
                &temp.User.ID, &temp.User.ClientID, &temp.WardesProfile,
                &temp.User.Username, &temp.User.Email,
                &temp.User.CreatedAt.Time, &temp.User.CreatedBy, &temp.User.CreatedClient,
                &temp.User.UpdatedAt.Time, &temp.User.UpdatedBy, &temp.User.UpdatedClient,
            )
            return temp, err
        },
    ).FieldConverter(dao.NewParaFieldConverter().
        Add("username", "wp.username").
        Add("email", "wp.email"),
    )

    return d.GetListDataWithDefaultMustCheck(param)
}
