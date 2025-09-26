package wardes_profile_image

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Retrieves the profile image details for a specific user ID.
func (d *wardes_profile_imagePostgresqlSQLDAO) GetWardesProfileImageByUserID(
     ctx *context.ContextModel,
     id int64,
) (
     repository.WardesProfileImageModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT created_at, created_by, created_client, updated_at, updated_by, updated_client
        FROM %s
        WHERE nexchief_account_id = $1
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    var model repository.WardesProfileImageModel
    err := d.db.QueryRow(query, id).Scan(
        &model.CreatedAt, &model.CreatedBy, &model.CreatedClient, 
        &model.UpdatedAt, &model.UpdatedBy, &model.UpdatedClient,
    )
    return model, err
}
