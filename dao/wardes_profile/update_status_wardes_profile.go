package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will update an existing wardes_profile record.
func (d *wardes_profilePostgresqlSQLDAO) UpdateStatusWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            status = $1,
            updated_by = $2,
            updated_client = $3,
            updated_at = $4
        WHERE id = $5
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.Status.String,
        ctx.Limitation.UserID,
        ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time,
        dtoIn.ID.Int64,
    )
    return err
}
