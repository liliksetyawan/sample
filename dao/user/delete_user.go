package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Soft delete a user record
func (d *userPostgresqlSQLDAO) DeleteUser(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.UserModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            deleted = true, updated_by = $1, updated_at = $2, updated_client = $3
        WHERE uuid_key = $4
    `, dao.GetDBTable(ctx, "user"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        ctx.Limitation.UserID, dtoIn.UpdatedAt.Time, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UUIDKey.String,
    )
    return err
}
