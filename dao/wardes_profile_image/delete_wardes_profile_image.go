package wardes_profile_image

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Delete wardes_profile_image record
func (d *wardes_profile_imagePostgresqlSQLDAO) DeleteWardesProfileImage(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileImageModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            deleted = true, updated_by = $1, 
            updated_client = $2, updated_at = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

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
