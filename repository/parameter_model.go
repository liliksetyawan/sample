package repository

import (
	"database/sql"
)

type ParameterModel struct {
	ID                sql.NullInt64
	UUIDKey           sql.NullString
	ParameterGroupID  sql.NullInt64
	Name              sql.NullString
	NameEn            sql.NullString
	Code              sql.NullString
	Sequence          sql.NullInt32
	Level             sql.NullInt16
	Type              sql.NullString
	Length            sql.NullInt16
	MinVal            sql.NullInt32
	MaxVal            sql.NullInt32
	DefaultVal        sql.NullString
	SelectListValue   sql.NullString
	SelectListName    sql.NullString
	Description       sql.NullString
	ParameterID       sql.NullInt64
	CreatedClient     sql.NullString
	UpdatedClient     sql.NullString
	CreatedAt         sql.NullTime
	CreatedBy         sql.NullInt64
	UpdatedAt         sql.NullTime
	UpdatedBy         sql.NullInt64
	Deleted           sql.NullBool
}