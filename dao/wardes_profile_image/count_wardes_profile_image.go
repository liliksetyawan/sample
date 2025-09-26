package wardes_profile_image

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Counts the number of records in the wardes_profile_image table
func (d *wardes_profile_imagePostgresqlSQLDAO) CountWardesProfileImage(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    tableName := fmt.Sprintf(`
        %s wpi
        JOIN %s wp ON wpi.nexchief_account_id = wp.id
    `, dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "wardes_profile"))

    getCountParam := d.NewGetListDataParam(
        "",
        dtoIn,
        searchParams,
        nil,
    ).TableName(
        tableName,
    )

    return d.GetCountDataWithDefaultMustCheck(getCountParam)
}
