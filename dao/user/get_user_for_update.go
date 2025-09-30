package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get Data from database for optimistic lock
func (d *userPostgresqlSQLDAO) GetUserForUpdate(
    ctx *context.ContextModel,
    tx *sql.Tx,
    uuid_key string,
) (
    repository.UserModel,
    error,
) {
    query := fmt.Sprintf(`
        SELECT id, updated_at
        FROM %s
        WHERE uuid_key = $1
    `, dao.GetDBTable(ctx, "user"))

    param := []interface{}{uuid_key}
    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2 "
        param = append(param, ctx.Limitation.UserID)
    }

    query += " FOR UPDATE "

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.UserModel{}, err
    }

    var model repository.UserModel
    err = stmt.QueryRow(param...).Scan(&model.ID, &model.UpdatedAt)

    return model, err
}