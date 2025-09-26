package wardes_profile_image

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update the existing wardes profile image record.
func (d *wardes_profile_imagePostgresqlSQLDAO) UpdateWardesProfileImage(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileImageModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s
        SET path_image = $1, updated_by = $2, updated_client = $3, updated_at = $4
        WHERE id = $5
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.PathImage.String, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID, dtoIn.UpdatedAt.Time, dtoIn.ID.Int64,
    )
    return err
}
