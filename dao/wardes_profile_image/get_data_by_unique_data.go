package wardes_profile_image

import (
    "database/sql"
    "fmt"
    "strings"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *wardes_profile_imagePostgresqlSQLDAO) GetDataByUniqueData(
      ctx *context.ContextModel,
      dtoIn repository.WardesProfileImageModel,
)( 
    result repository.WardesProfileImageModel,
    err error,
){
    var isUqWardesProfileImageType bool

    query := fmt.Sprintf(`
    SELECT
        id, 
        ( type = $1 ) as unique_2
    FROM %s
    WHERE 
        (type = $1)
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

    args := []interface{}{
        dtoIn.Type.String,
    }

    if ctx.Limitation.UserID != 0 {
        query += " AND created_by = $2 "
        args = append(args, ctx.Limitation.UserID)
    }

    query += " LIMIT 1 "

    if err = d.db.QueryRow(query, args...).Scan(&result.ID, &isUqWardesProfileImageType); err != nil && err != sql.ErrNoRows {
        return
    }

    if result.ID.Int64 != 0 {
        var existValue []string
        
        if isUqWardesProfileImageType {
            existValue = append(existValue, "[type]")
        }
        
        err = fmt.Errorf("data already used: %s", strings.Join(existValue, ", "))
        return
    }

    err = nil
    return
}
