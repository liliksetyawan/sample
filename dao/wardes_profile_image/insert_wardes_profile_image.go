package wardes_profile_image

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Insert a new wardes profile image record
func (d *wardes_profile_imagePostgresqlSQLDAO) InsertWardesProfileImage(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileImageModel,
) (
     repository.WardesProfileImageModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            nexchief_account_id, wardes_profile_id, type, path_image,
            created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, $4,
            $5, $6, $7,
            $8, $9, $10
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.WardesProfileImageModel{}, err
    }

    var model repository.WardesProfileImageModel
    err = stmt.QueryRow(
        dtoIn.NexchiefAccountID.Int64, dtoIn.WardesProfileID.Int64, dtoIn.Type.String, dtoIn.PathImage.String,
        dtoIn.CreatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&model.ID, &model.UUIDKey)
    
    return model, err
}