package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Count total wardes profile records based on search criteria
func (d *wardes_profilePostgresqlSQLDAO) CountListWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    table := fmt.Sprintf(`%s`, dao.GetDBTable(ctx, "wardes_profile"))

    param := d.NewGetListDataParam(
        "",
        dtoIn,
        searchParams,
        nil,
    ).TableName(
        table,
    ).Prefix(dao.NewParamPrefix().
        Add("id", "wp"),
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetCountDataWithDefaultMustCheck(param)
}
