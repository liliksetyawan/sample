package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update existing wardes profile data in database
func (d *wardes_profilePostgresqlSQLDAO) UpdateWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            updated_at = $1, updated_by = $2, updated_client = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.UpdatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.ID.Int64,
    )
    return err
}
