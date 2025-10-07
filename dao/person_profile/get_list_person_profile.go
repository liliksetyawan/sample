package person_profile

import (
	"database/sql"
	"fmt"
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dao"
	"github.com/nexsoft-git/nexcommon/dto/in"
	"github.com/nexsoft-git/nexcommon/model"
	"nexsoft.co.id/example/repository"
)

// Descriptions: Retrieve a list of person profiles
func (d *person_profilePostgresqlSQLDAO) GetListPersonProfile(
	ctx *context.ContextModel,
	dtoIn in.GetListRequest,
	searchParams []model.SearchParam,
) (
	[]interface{},
	error,
) {
	query := fmt.Sprintf(`
        SELECT pp.id, pp.title_id, pp.title_name, pp.first_name, pp.last_name,
               wp.id, wp.username, wp.personal_name,
               wpi.id, wpi.path_image,
               u.id, u.client_id, u.user_code
        FROM %s pp
        LEFT JOIN %s wp ON pp.id = wp.person_profile_id
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        LEFT JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "person_profile"), dao.GetDBTable(ctx, "wardes_profile"),
		dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "user"))

	param := d.NewGetListDataParam(
		query,
		dtoIn,
		searchParams, func(rows *sql.Rows) (interface{}, error) {
			var temp repository.PersonProfileDetailModel
			err := rows.Scan(
				&temp.PersonProfile.ID, &temp.PersonProfile.TitleID, &temp.PersonProfile.TitleName,
				&temp.PersonProfile.FirstName, &temp.PersonProfile.LastName,
				&temp.WardesProfile.ID, &temp.WardesProfile.Username, &temp.WardesProfile.PersonalName,
				&temp.WardesProfileImage.ID, &temp.WardesProfileImage.PathImage,
				&temp.User.ID, &temp.User.ClientID, &temp.User.UserCode,
			)
			return temp, err
		},
	).FieldConverter(dao.NewParaFieldConverter().
		Add("name", "pp.first_name || ' ' || pp.last_name").
		Add("email", "wp.username"),
	).Prefix(dao.NewParamPrefix().
		Add("id", "pp"),
	)

	return d.GetListDataWithDefaultMustCheck(param)
}
