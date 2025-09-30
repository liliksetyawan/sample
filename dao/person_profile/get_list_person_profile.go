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
               u.id, u.client_id, u.user_code,
               pp.created_at, pp.created_by, pp.created_client,
               pp.updated_at, pp.updated_by, pp.updated_client
        FROM %s pp
        LEFT JOIN %s wp ON pp.person_profile_id = wp.id
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        LEFT JOIN %s u ON pp.created_by = u.id
    `, dao.GetDBTable(ctx, "person_profile"), dao.GetDBTable(ctx, "wardes_profile"),
		dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "user"))

	param := d.NewGetListDataParam(
		query,
		dtoIn,
		searchParams, func(rows *sql.Rows) (interface{}, error) {
			var temp repository.PersonProfileModel
			err := rows.Scan(
				&temp.Id, &temp.TitleId, &temp.TitleName, &temp.FirstName, &temp.LastName,
				&temp.CountryId, &temp.CountryName,
				&temp.ProvinceId, &temp.ProvinceName,
				&temp.DistrictId, &temp.DistrictName,
				&temp.SubDistrictId, &temp.SubDistrictName,
				&temp.UrbanVillageId, &temp.UrbanVillageName,
				&temp.IslandId, &temp.IslandName,
				&temp.PostalCodeId, &temp.PostalCodeName,
				&temp.Phone, &temp.Email, &temp.Photo,
				&temp.CreatedAt, &temp.CreatedBy, &temp.CreatedClient,
				&temp.UpdatedAt, &temp.UpdatedBy, &temp.UpdatedClient,
			)
			return temp, err
		},
	).FieldConverter(dao.NewParaFieldConverter().
		Add("name", "pp.first_name || ' ' || pp.last_name").
		Add("email", "u.user_code"),
	).CreatedBy(ctx.Limitation.UserID)

	return d.GetListDataWithDefaultMustCheck(param)
}
