package wardes_profile

import (
	"database/sql"
	"fmt"
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dao"
	"github.com/nexsoft-git/nexcommon/dto/in"
	"github.com/nexsoft-git/nexcommon/model"
	"nexsoft.co.id/example/repository"
)

// Descriptions: This function retrieves a list of wardes profiles.
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfileNexcore(
	ctx *context.ContextModel,
	dtoIn in.GetListRequest,
	searchParams []model.SearchParam,
) (
	[]interface{},
	error,
) {
	query := fmt.Sprintf(`
        SELECT na.id, na.code, wp.id, wp.code, wp.name, wp.username, wp.phone, wp.email, u.id, u.name
        FROM %s wp
        JOIN %s na ON wp.nexchief_account_id = na.id
        JOIN %s u ON wp.user_id = u.id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "nexchief_account"), dao.GetDBTable(ctx, "user"))

	param := d.NewGetListDataParam(
		query,
		dtoIn,
		searchParams, func(rows *sql.Rows) (interface{}, error) {
			var temp repository.WardesProfileModel
			err := rows.Scan(
				&temp.ID, &temp.Code, &temp.Username, &temp.Phone, &temp.Email,
				&temp.CreatedBy, &temp.CreatedClient, &temp.CreatedAt,
				&temp.UpdatedBy, &temp.UpdatedAt, &temp.UpdatedClient,
			)
			return temp, err
		},
	).FieldConverter(dao.NewParaFieldConverter().
		Add("username", "wp.username").
		Add("phone", "wp.phone").
		Add("email", "wp.email"),
	).AdditionalWhere(" wp.deleted = false ").
		CreatedBy(ctx.Limitation.UserID)

	return d.GetListDataWithDefaultMustCheck(param)
}
