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
    // @Affected Field: created_at, created_by, created_client, updated_at, updated_by, updated_client
    // @Affected Table: dao.GetDBTable(ctx, "wardes_profile_image")  

    query := fmt.Sprintf(`
        SELECT id, uuid_key, nexchief_account_id, wardes_profile_id, type, path_image,
            created_by, created_client, created_at, updated_by, updated_at, updated_client, deleted
        FROM %s
        WHERE nexchief_account_id = $1
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    var model repository.WardesProfileImageModel
    err := d.db.QueryRow(query, id).Scan(
        &model.ID, &model.UUIDKey, &model.NexchiefAccountID, &model.WardesProfileID, &model.Type, &model.PathImage,
        &model.CreatedBy, &model.CreatedClient, &model.CreatedAt, &model.UpdatedBy, &model.UpdatedAt, &model.UpdatedClient, &model.Deleted,
    )
    return model, err
}
