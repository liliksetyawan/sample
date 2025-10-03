package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Delete wardes profile record
func (d *wardes_profilePostgresqlSQLDAO) DeleteWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            deleted = true,
            updated_by = $1,
            updated_client = $2,
            updated_at = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        ctx.Limitation.UserID,
        ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time,
        dtoIn.ID.Int64,
    )
    return err
}
