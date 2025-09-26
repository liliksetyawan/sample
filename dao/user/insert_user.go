package user

import (
	"database/sql"
	"fmt"
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dao"
	"nexsoft.co.id/example/repository"
)

// Descriptions: This function will insert a new user into the database.
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
            created_at, created_by, created_client, 
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, 
            $4, $5, $6
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "user"))

	stmt, err := tx.Prepare(query)
	if err != nil {
		return repository.UserModel{}, err
	}

	var model repository.UserModel
	err = stmt.QueryRow(
		dtoIn.CreatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
		dtoIn.UpdatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
	).Scan(&model.ID, &model.UUIDKey)
	return model, err
}
