package user

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will retrieve a user from the database based on a given ID.
func (d *userPostgresqlSQLDAO) GetUser(
     ctx *context.ContextModel,
     uuid_key string,
) (
     repository.UserModel, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT uuid_key, created_by, created_client, created_at, updated_by, updated_client, updated_at
        FROM %s
        WHERE uuid_key = $1
    `, dao.GetDBTable(ctx, "user"))

    var model repository.UserModel
    err := d.db.QueryRow(query, uuid_key).Scan(
        &model.UUIDKey, &model.CreatedBy, &model.CreatedClient, &model.CreatedAt,
        &model.UpdatedBy, &model.UpdatedClient, &model.UpdatedAt,
    )

    return model, err
}
