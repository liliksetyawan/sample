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

// Descriptions: Get list of wardes_profile records
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
               wpi.id, wpi.type, wpi.path_image,
               u.id, u.auth_user_id, u.client_id
        FROM %s wp
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        LEFT JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.GetListWardesProfileModel
            err := rows.Scan(
                &temp.WardesProfile.ID, &temp.WardesProfile.Username, &temp.WardesProfile.PersonalName, 
                &temp.WardesProfile.Gender, &temp.WardesProfile.Phone, &temp.WardesProfile.Email,
                &temp.WardesProfileImage.ID, &temp.WardesProfileImage.Type, &temp.WardesProfileImage.PathImage,
                &temp.User.ID, &temp.User.AuthUserID, &temp.User.ClientID,
            )
            return temp, err
        },
    ).Prefix(
        dao.NewParamPrefix().
            Add("id", "wp"),
    ).FieldConverter(
        dao.NewParaFieldConverter().
            Add("wp.username", "wp.username"),
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
