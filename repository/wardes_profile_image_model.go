package repository

import (
	"database/sql"
)

type WardesProfileImageModel struct {
	ID                sql.NullInt64
	UUIDKey           sql.NullString
	NexchiefAccountID sql.NullInt64
	WardesProfileID   sql.NullInt64
	Type              sql.NullString
	PathImage         sql.NullString
	CreatedBy         sql.NullInt64
	CreatedClient     sql.NullString
	CreatedAt         sql.NullTime
	UpdatedBy         sql.NullInt64
	UpdatedAt         sql.NullTime
	UpdatedClient     sql.NullString
	Deleted           sql.NullBool
	IsTemporary       sql.NullBool
}