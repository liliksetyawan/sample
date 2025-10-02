package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Delete a person profile record
func (d *person_profilePostgresqlSQLDAO) DeletePersonProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.PersonProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            deleted = true, updated_by = $1, 
            updated_client = $2, updated_at = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "person_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        ctx.Limitation.ServiceUserID, 
        ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time,
        dtoIn.ID.Int64,
    )
    return err
}

