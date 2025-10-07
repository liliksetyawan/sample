package wardes_profile

import (
    "fmt"
    "strings"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/common/commonError"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *wardes_profilePostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.WardesProfileModel,
)(
    result repository.WardesProfileModel,
    err error,
){
    var isUqWardesProfileUsername, isUqWardesProfileNik bool

    query := fmt.Sprintf(`
    SELECT
        id, 
        ( username = $1 ) as unique_2,
        ( nik = $2 ) as unique_3 
    FROM %s
    WHERE 
        (username = $1) 
        OR 
        (nik = $2) 
    `, dao.GetDBTable(ctx, "wardes_profile"))

    args := []interface{}{
        dtoIn.Username.String,
        dtoIn.NIK.String,
    }

    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $3 "
        args = append(args, ctx.Limitation.UserID)
    }

    query += " LIMIT 1 "

    if err = d.db.QueryRow(query, args...).Scan(&result.ID, &isUqWardesProfileUsername, &isUqWardesProfileNik); err != nil && err != sql.ErrNoRows {
        return
    }

    if result.ID.Int64 != 0 {
        var existValue []string
        
        if isUqWardesProfileUsername {
            existValue = append(existValue, "[username]")
        }
        
        if isUqWardesProfileNik {
            existValue = append(existValue, "[nik]")
        }
        
        err = commonError.ErrDataAlreadyUsed.Param(strings.Join(existValue, ", "))
        return
    }

    err = nil
    //No Unique Constraint on the database, return empty result
    return
}
