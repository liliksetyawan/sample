package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Get list of wardes profile data with joins to wardes_profile_image and user tables
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT wp.id, wp.username, wp.personal_name, wp.gender, wp.phone, wp.email,
            wpi.path_image, wpi.id, u.id, u.client_id
        FROM %s wp
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        LEFT JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.ListWardesProfileModel
            err := rows.Scan(
                &temp.WardesProfile.ID, &temp.WardesProfile.Username, &temp.WardesProfile.PersonalName, 
                &temp.WardesProfile.Gender, &temp.WardesProfile.Phone, &temp.WardesProfile.Email,
                &temp.WardesProfileImage.PathImage, &temp.WardesProfileImage.ID,
                &temp.User.ID, &temp.User.ClientID,
            )
            return temp, err
        },
    ).Prefix(dao.NewParamPrefix().
        Add("id", "wp"),
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
