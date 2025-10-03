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

// Descriptions: Get list of wardes_profile_image records with search and pagination
func (d *wardes_profile_imagePostgresqlSQLDAO) GetListWardesProfileImage(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, created_by, created_client, created_at,
            updated_by, updated_client, updated_at
        FROM %s
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.WardesProfileImageModel
            err := rows.Scan(
                &temp.UUIDKey, &temp.CreatedBy, &temp.CreatedClient, &temp.CreatedAt,
                &temp.UpdatedBy, &temp.UpdatedClient, &temp.UpdatedAt,
            )
            return temp, err
        },
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
