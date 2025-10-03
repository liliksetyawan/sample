package wardes_profile_image

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Count total wardes_profile_image records matching search criteria
func (d *wardes_profile_imagePostgresqlSQLDAO) CountListWardesProfileImage(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    table := fmt.Sprintf(`%s`, dao.GetDBTable(ctx, "wardes_profile_image"))

    param := d.NewGetListDataParam(
        "",
        dtoIn,
        searchParams,
        nil,
    ).TableName(table).CreatedBy(ctx.Limitation.UserID)

    return d.GetCountDataWithDefaultMustCheck(param)
}