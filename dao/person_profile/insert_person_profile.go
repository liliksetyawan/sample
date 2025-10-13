
package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Insert a new person profile record
func (d *person_profilePostgresqlSQLDAO) InsertPersonProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.PersonProfileModel,
) (
     repository.PersonProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            person_profile_id, title_id, title_name, first_name, last_name, sex, nik,
            created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7,
            $8, $9, $10,
            $11, $12, $13
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "person_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.PersonProfileModel{}, err
    }

    var model repository.PersonProfileModel
    err = stmt.QueryRow(
        dtoIn.PersonProfileID.Int64, dtoIn.TitleID.Int64, dtoIn.TitleName.String, dtoIn.FirstName.String, dtoIn.LastName.String, dtoIn.Sex.String, dtoIn.NIK.String,
        dtoIn.CreatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&model.ID, &model.UUIDKey)
    
    return model, err
}

