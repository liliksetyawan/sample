package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get details of the wardes_profile entity.
func (d *wardes_profilePostgresqlSQLDAO) GetDetailWardesProfile(
     ctx *context.ContextModel,
     uuid_key string,
) (
     repository.WardesProfileImageDetailModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT wp.id, wp.username, wp.personal_name, wp.email,
               wpi.id, wpi.path_image,
               wp.created_at, wp.created_by, wp.created_client,
               wp.updated_at, wp.updated_by, wp.updated_client
        FROM %s wp
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        WHERE wp.uuid_key = $1
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "wardes_profile_image"))

    var model repository.WardesProfileImageDetailModel
    err := d.db.QueryRow(query, uuid_key).Scan(
        &model.Id, &model.Username, &model.PersonalName, &model.Email,
        &model.Id, &model.PathImage,
        &model.CreatedAt, &model.CreatedBy, &model.CreatedClient,
        &model.UpdatedAt, &model.UpdatedBy, &model.UpdatedClient,
    )
    return model, err
}
