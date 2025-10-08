package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Insert new wardes profile data into database
func (d *wardes_profilePostgresqlSQLDAO) InsertWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     repository.WardesProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3,
            $4, $5, $6
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.WardesProfileModel{}, err
    }

    var model repository.WardesProfileModel
    err = stmt.QueryRow(
        dtoIn.CreatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&model.ID, &model.UUIDKey)
    
    return model, err
}
