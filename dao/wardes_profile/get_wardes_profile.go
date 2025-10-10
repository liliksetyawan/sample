package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get single wardes profile record by ID
func (d *wardes_profilePostgresqlSQLDAO) GetWardesProfile(
     ctx *context.ContextModel,
     uuid_key string,
) (
     repository.WardesProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, created_by, created_client, created_at,
               updated_by, updated_client, updated_at
        FROM %s
        WHERE uuid_key = $1 AND deleted = false
    `, dao.GetDBTable(ctx, "wardes_profile"))

    args := []interface{}{uuid_key}
    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2"
        args = append(args, ctx.Limitation.UserID)
    }

    var model repository.WardesProfileModel
    err := d.db.QueryRow(query, args...).Scan(
        &model.UUIDKey, &model.CreatedBy, &model.CreatedClient, &model.CreatedAt,
        &model.UpdatedBy, &model.UpdatedClient, &model.UpdatedAt,
    )

    return model, err
}
