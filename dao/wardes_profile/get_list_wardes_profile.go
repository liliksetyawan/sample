package wardes_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Retrieves a list of wardes profiles
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT na.id, na.code, wp.id, wp.code, wp.username, wp.phone, wp.email, u.id, u.username
        FROM %s wp
        JOIN %s na ON wp.nexchief_account_id = na.id
        JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "nexchief_account"), dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.WardesProfileModel
            err := rows.Scan(
                &temp.ID, &temp.NexchiefAccountID, &temp.Username, &temp.Phone, &temp.Email,
            )
            return temp, err
        },
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
