package person_profile

import (
	"database/sql"
	"github.com/nexsoft-git/nexcommon/context"
	"nexsoft.co.id/example/repository"
)

type PersonProfileDAO interface {
	GetPersonProfileByUserID(
		ctx *context.ContextModel,
		id int64,
	) (
		repository.PersonProfileModel,
		error,
	)

	InsertPersonProfile(
		ctx *context.ContextModel,
		tx *sql.Tx,
		dtoIn repository.PersonProfileModel,
	) (
		repository.PersonProfileModel,
		error,
	)

	GetDataByUniqueData(
		ctx *context.ContextModel,
		dtoIn repository.PersonProfileModel,
	) (
		result repository.PersonProfileModel,
		err error,
	)

	GetPersonProfileForUpdate(
		ctx *context.ContextModel,
		tx *sql.Tx,
		uuid_key string,
	) (
		repository.PersonProfileModel,
		error,
	)
}
