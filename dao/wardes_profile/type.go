package wardes_profile

import (
	"database/sql"
	"github.com/nexsoft-git/nexcommon/context"
	"nexsoft.co.id/example/repository"
)

type WardesProfileDAO interface {
	UpdateWardesProfile(
		ctx *context.ContextModel,
		tx *sql.Tx,
		dtoIn repository.WardesProfileModel,
	) error

	InsertWardesProfile(
		ctx *context.ContextModel,
		tx *sql.Tx,
		dtoIn repository.WardesProfileModel,
	) (
		repository.WardesProfileModel,
		error,
	)

	GetDataByUniqueData(
		ctx *context.ContextModel,
		dtoIn repository.WardesProfileModel,
	) (
		result repository.WardesProfileModel,
		err error,
	)

	GetWardesProfileForUpdate(
		ctx *context.ContextModel,
		tx *sql.Tx,
		uuid_key string,
	) (
		repository.WardesProfileModel,
		error,
	)
}
