package wardes_profile_image

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will count the records in the wardes_profile_image table.
func (d *wardes_profile_imagePostgresqlSQLDAO) CountWardesProfileImage(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    // @Affected Field: created_at, created_by, created_client, updated_at, updated_by, updated_client
    // @Affected Table: dao.GetDBTable(ctx, "wardes_profile_image")  
    // @Valid Search:
    // - code : The code associated with the profile image
    // - name : The name associated with the profile image

    table := dao.GetDBTable(ctx, "wardes_profile_image")

    query := fmt.Sprintf(`
        SELECT COUNT(*)
        FROM %s
    `, table)

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams,
        nil,
    ).AdditionalWhere(" code = ?? OR name = ?? ").
    AdditionalWhereParam(dtoIn.Code, dtoIn.Name)

    count, err := d.GetCountDataWithDefaultMustCheck(param)
    return count, err
}
