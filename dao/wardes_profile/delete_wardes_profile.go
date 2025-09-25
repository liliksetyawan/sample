package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function deletes a wardes profile by its ID.
func (d *wardes_profilePostgresqlSQLDAO) DeleteWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            deleted = true, updated_at = $1, updated_by = $2, updated_client = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "wardes_profile"))

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
