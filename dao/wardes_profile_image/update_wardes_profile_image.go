package wardes_profile_image

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update existing wardes_profile_image record
func (d *wardes_profile_imagePostgresqlSQLDAO) UpdateWardesProfileImage(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileImageModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            updated_at = $1, updated_by = $2, updated_client = $3
        WHERE id = $4
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

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
