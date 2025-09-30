package repository

import (
	"database/sql"
)

type PersonProfileModel struct {
	Id               sql.NullInt64
	UuidKey          sql.NullString
	PersonProfileId  sql.NullInt64
	TitleId          sql.NullInt64
	TitleName        sql.NullString
	FirstName        sql.NullString
	LastName         sql.NullString
	Sex              sql.NullString
	Nik              sql.NullString
	Address1         sql.NullString
	Address2         sql.NullString
	Address3         sql.NullString
	Hamlet           sql.NullString
	Neighbour        sql.NullString
	CountryId        sql.NullInt64
	CountryName      sql.NullString
	ProvinceId       sql.NullInt64
	ProvinceName     sql.NullString
	DistrictId       sql.NullInt64
	DistrictName     sql.NullString
	SubDistrictId    sql.NullInt64
	SubDistrictName  sql.NullString
	UrbanVillageId   sql.NullInt64
	UrbanVillageName sql.NullString
	IslandId         sql.NullInt64
	IslandName       sql.NullString
	PostalCodeId     sql.NullInt64
	PostalCodeName   sql.NullString
	Phone            sql.NullString
	Email            sql.NullString
	Photo            sql.NullString
	LastSync         sql.NullTime
	CreatedClient     sql.NullString
	UpdatedClient     sql.NullString
	CreatedAt        sql.NullTime
	CreatedBy        sql.NullInt64
	UpdatedAt        sql.NullTime
	UpdatedBy        sql.NullInt64
	Deleted          sql.NullBool
}