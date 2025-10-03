package parameter

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Insert a new parameter record
func (d *parameterPostgresqlSQLDAO) InsertParameter(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.ParameterModel,
) (
     repository.ParameterModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            parameter_group_id, name, name_en, code, sequence, level, type,
            created_at, created_by, created_client,
            updated_at, updated_by, updated_client
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7,
            $8, $9, $10,
            $11, $12, $13
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "parameter"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.ParameterModel{}, err
    }

    var model repository.ParameterModel
    err = stmt.QueryRow(
        dtoIn.ParameterGroupID.Int64, dtoIn.Name.String, dtoIn.NameEn.String, dtoIn.Code.String,
        dtoIn.Sequence.Int32, dtoIn.Level.Int16, dtoIn.Type.String,
        dtoIn.CreatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
        dtoIn.UpdatedAt.Time, ctx.Limitation.ServiceUserID, ctx.AuthAccessTokenModel.ClientID,
    ).Scan(&model.ID, &model.UUIDKey)
    
    return model, err
}
