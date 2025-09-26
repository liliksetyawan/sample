package user

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will count the number of users in the database based on given parameters.
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
    )

    return d.GetCountDataWithDefaultMustCheck(getCountParam)
}
