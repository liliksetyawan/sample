package wardes_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function counts the total number of wardes_profile records.
func (d *wardes_profilePostgresqlSQLDAO) CountWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     int, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT COUNT(*)
        FROM %s wp
        JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "user"))

    // Adding search conditions based on valid search parameters
    var conditions []string
    var args []interface{}

    if dtoIn.Code.String != "" {
        conditions = append(conditions, "wp.code = $1")
        args = append(args, dtoIn.Code.String)
    }

    if dtoIn.Name.String != "" {
        conditions = append(conditions, "wp.personal_name ILIKE '%' || $2 || '%'")
        args = append(args, dtoIn.Name.String)
    }

    if len(conditions) > 0 {
        query += " WHERE " + fmt.Sprintf("%s", conditions)
    }

    var count int
    err := d.db.QueryRow(query, args...).Scan(&count)
    return count, err
}
