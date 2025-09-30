package user

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *userPostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.UserModel,
)(
    result repository.UserModel,
    err error,
){
    var isUqUserUsername bool

    query := fmt.Sprintf(`
    SELECT
        id, 
        ( username = $1 ) as unique_2 
    FROM %s
    WHERE 
        (username = $1) 
    `, dao.GetDBTable(ctx, "user"))

    args := []interface{}{
        dtoIn.Username.String,
    }

    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2 "
        args = append(args, ctx.Limitation.UserID)
    }

    query += " LIMIT 1 "

    if err = d.db.QueryRow(query, args...).Scan(&result.ID, &isUqUserUsername); err != nil && err != sql.ErrNoRows {
        return
    }

    if result.ID.Int64 != 0 {
        var existValue []string
        
        if isUqUserUsername {
            existValue = append(existValue, "[username]")
        }
        
        err = commonError.ErrDataAlreadyUsed.Param(strings.Join(existValue, ", "))
        return
    }

    err = nil
    return
}
