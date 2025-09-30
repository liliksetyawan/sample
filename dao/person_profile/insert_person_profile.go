package person_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will insert a new person profile entry into the database.
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
            person_profile_id, title_id, title_name, first_name, last_name,
            sex, phone, email, created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10, $11,
            $12, $13, $14
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "person_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.PersonProfileModel{}, err
    }

    var model repository.PersonProfileModel
    err = stmt.QueryRow(
        dtoIn.PersonProfileId.Int64, dtoIn.TitleId.Int64, dtoIn.TitleName.String,
        dtoIn.FirstName.String, dtoIn.LastName.String, dtoIn.Sex.String,
        dtoIn.Phone.String, dtoIn.Email.String, dtoIn.CreatedAt.Time,
        ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&model.Id, &model.UuidKey)
    return model, err
}
