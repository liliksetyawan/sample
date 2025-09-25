package wardes_profile

import (
    "database/sql"
    "fmt"
    "strings"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
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
    var isUqWardesProfileNik bool

    query := fmt.Sprintf(`
    SELECT
        id, 
        ( nik = $1 ) as unique_2 
    FROM %s
    WHERE 
        (nik = $1) 
    `, dao.GetDBTable(ctx, "wardes_profile"))

    args := []interface{}{
        dtoIn.Nik.String,
    }

    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2 "
        args = append(args, ctx.Limitation.UserID)
    }

    query += " LIMIT 1 "

    if err = d.db.QueryRow(query, args...).Scan(&result.ID, &isUqWardesProfileNik); err != nil && err != sql.ErrNoRows {
        return
    }

    if result.ID.Int64 != 0 {
        var existValue []string
        
        if isUqWardesProfileNik {
            existValue = append(existValue, "[nik]")
        }
        
        err = fmt.Errorf("data already used: %s", strings.Join(existValue, ", "))
        return
    }

    return
}
