package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update an existing person profile record
func (d *person_profilePostgresqlSQLDAO) UpdatePersonProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.PersonProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            updated_at = $1, updated_by = $2, updated_client = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "person_profile"))

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

