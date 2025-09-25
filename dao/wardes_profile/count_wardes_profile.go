package wardes_profile

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will count the total number of wardes_profile records.
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
        FROM %s
    `, dao.GetDBTable(ctx, "wardes_profile"))

    var count int
    err := d.db.QueryRow(query).Scan(&count)
    return count, err
}
