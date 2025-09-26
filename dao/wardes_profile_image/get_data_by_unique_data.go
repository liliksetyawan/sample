package wardes_profile_image

import (
	"database/sql"
	"fmt"
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dao"
	"nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *wardes_profile_imagePostgresqlSQLDAO) GetDataByUniqueData(
	ctx *context.ContextModel,
	dtoIn repository.WardesProfileImageModel,
) (
	result repository.WardesProfileImageModel,
	err error,
) {
	query := fmt.Sprintf(`
        SELECT id, uuid_key
        FROM %s
        WHERE nexchief_account_id = $1 AND wardes_profile_id = $2
    `, dao.GetDBTable(ctx, "wardes_profile_image"))

	err = d.db.QueryRow(query, dtoIn.NexchiefAccountID.Int64, dtoIn.WardesProfileID.Int64).Scan(
		&result.ID, &result.UUIDKey,
	)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}
	return result, nil
}
