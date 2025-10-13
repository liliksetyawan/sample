package person_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"nexsoft.co.id/example/repository"
)

// Descriptions: Check is data is exist before inserting/updating
func (d *person_profilePostgresqlSQLDAO) GetDataByUniqueData(
	ctx *context.ContextModel,
	dtoIn repository.PersonProfileModel,
) (
	result repository.PersonProfileModel,
	err error,
) {
	//No Unique Constraint on the database, return empty result
	return
}
