package wardes_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will retrieve the details of a wardes_profile by its ID.
func (d *wardes_profilePostgresqlSQLDAO) GetWardesProfile(
     ctx *context.ContextModel,
     uuid_key string,
) (
     repository.WardesProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT created_at, created_by, created_client, 
               updated_at, updated_by, updated_client
        FROM %s
        WHERE uuid_key = $1
    `, dao.GetDBTable(ctx, "wardes_profile"))

    var model repository.WardesProfileModel
    err := d.db.QueryRow(query, uuid_key).Scan(
        &model.CreatedAt, &model.CreatedBy, &model.CreatedClient,
        &model.UpdatedAt, &model.UpdatedBy, &model.UpdatedClient,
    )
    return model, err
}
