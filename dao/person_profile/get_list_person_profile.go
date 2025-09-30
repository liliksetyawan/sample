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
        LEFT JOIN %s wp ON pp.person_profile_id = wp.id
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        RIGHT JOIN %s u ON pp.created_by = u.id
    `, dao.GetDBTable(ctx, "person_profile"),
		dao.GetDBTable(ctx, "wardes_profile"),
		dao.GetDBTable(ctx, "wardes_profile_image"),
		dao.GetDBTable(ctx, "user"))

	param := d.NewGetListDataParam(
		query,
		dtoIn,
		searchParams, func(rows *sql.Rows) (interface{}, error) {
			var temp repository.PersonProfileModel
			err := rows.Scan(
				&temp.Id, &temp.TitleId, &temp.TitleName, &temp.FirstName, &temp.LastName,
				&temp.CountryId, &temp.CountryName, &temp.ProvinceId, &temp.ProvinceName,
				&temp.DistrictId, &temp.DistrictName, &temp.SubDistrictId, &temp.SubDistrictName,
				&temp.UrbanVillageId, &temp.UrbanVillageName, &temp.PostalCodeId, &temp.PostalCodeName,
				&temp.Phone, &temp.Email, &temp.Photo,
				&temp.CreatedClient, &temp.UpdatedClient, &temp.CreatedAt,
				&temp.CreatedBy, &temp.UpdatedAt, &temp.UpdatedBy, &temp.Deleted,
			)
			return temp, err
		},
	).FieldConverter(dao.NewParaFieldConverter().
		Add("name", "pp.first_name || ' ' || pp.last_name").
		Add("email", "u.email"),
	)

	return d.GetListDataWithDefaultMustCheck(param)
}
