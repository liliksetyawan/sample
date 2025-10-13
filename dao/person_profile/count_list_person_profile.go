
package person_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/model"
)

// Descriptions: Count total person profiles with search capabilities
func (d *person_profilePostgresqlSQLDAO) CountListPersonProfile(
    ctx *context.ContextModel,
    dtoIn in.GetListRequest,
    searchParams []model.SearchParam,
) (
    int,
    error,
) {
    table := fmt.Sprintf(`%s`, dao.GetDBTable(ctx, "person_profile"))

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

