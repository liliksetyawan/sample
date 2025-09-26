package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will update an existing user in the database.
func (d *userPostgresqlSQLDAO) UpdateUser(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.UserModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            updated_at = $1, updated_by = $2, updated_client = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "user"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.UpdatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.ID.Int64,
    )
    return err
}
