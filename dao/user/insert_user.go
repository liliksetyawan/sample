package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Insert a new user record into the database
func (d *userPostgresqlSQLDAO) InsertUser(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.UserModel,
) (
     repository.UserModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            auth_user_id, person_profile_id, client_id, user_code, username,
            created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8,
            $9, $10, $11
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "user"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.UserModel{}, err
    }

    var result repository.UserModel
    err = stmt.QueryRow(
        dtoIn.AuthUserID.Int64, dtoIn.PersonProfileID.Int64, dtoIn.ClientID.String, dtoIn.UserCode.String, dtoIn.Username.String,
        dtoIn.CreatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&result.ID, &result.UUIDKey)
    
    return result, err
}
