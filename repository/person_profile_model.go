package repository

import (
	"database/sql"
)

type PersonProfileDetailModel struct {
	PersonProfile      PersonProfileModel
	WardesProfile      WardesProfileModel
	WardesProfileImage WardesProfileImageModel
	User               UserModel
}

type PersonProfileModel struct {
	ID        sql.NullInt64
	TitleID   sql.NullInt64
	TitleName sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
}

type UserModel struct {
	ID       sql.NullInt64
	ClientID sql.NullInt64
	UserCode sql.NullString
}
