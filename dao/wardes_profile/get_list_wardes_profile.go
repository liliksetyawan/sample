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

// Descriptions: Retrieves a list of wardes profiles
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfile(
	ctx *context.ContextModel,
	dtoIn in.GetListRequest,
	searchParams []model.SearchParam,
) (
	[]interface{},
	error,
) {
	query := fmt.Sprintf(`
        SELECT wp.id, wp.code, wp.username, wp.phone, wp.email, 
               u.id, u.username
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
				&temp.UserID, &temp.Username,
			)
			return temp, err
		},
	).FieldConverter(dao.NewParaFieldConverter().
		Add("code", "wp.code").
		Add("name", "wp.personal_name"),
	)

	return d.GetListDataWithDefaultMustCheck(param)
}
