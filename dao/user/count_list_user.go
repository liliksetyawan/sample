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

// Descriptions: Count the number of users matching search criteria
func (d *userPostgresqlSQLDAO) CountListUser(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    table := fmt.Sprintf(`%s`, dao.GetDBTable(ctx, "user"))

    getCountParam := d.NewGetListDataParam(
        "",
        dtoIn,
        searchParams,
        nil,
    ).TableName(
        table,
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetCountDataWithDefaultMustCheck(getCountParam)
}
