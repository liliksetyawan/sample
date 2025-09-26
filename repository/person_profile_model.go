package repository

import "database/sql"

type PersonProfileModel struct {
    ID             sql.NullInt64
    UUIDKey        sql.NullString
    PersonProfileID sql.NullInt64
    TitleID        sql.NullInt64
    TitleName      sql.NullString
    FirstName      sql.NullString
    LastName       sql.NullString
    Sex            sql.NullString
    NIK            sql.NullString
    Address1       sql.NullString
    Address2       sql.NullString
    Address3       sql.NullString
    Hamlet         sql.NullString
    Neighbour      sql.NullString
    CountryID      sql.NullInt64
    CountryName    sql.NullString
    ProvinceID     sql.NullInt64
    ProvinceName   sql.NullString
    DistrictID     sql.NullInt64
    DistrictName   sql.NullString
    SubDistrictID  sql.NullInt64
    SubDistrictName sql.NullString
    UrbanVillageID sql.NullInt64
    UrbanVillageName sql.NullString
    IslandID       sql.NullInt64
    IslandName     sql.NullString
    PostalCodeID   sql.NullInt64
    PostalCodeName sql.NullString
    Phone          sql.NullString
    Email          sql.NullString
    Photo          sql.NullString
    LastSync       sql.NullTime
    CreatedClient  sql.NullString
    UpdatedClient  sql.NullString
    CreatedAt      sql.NullTime
    CreatedBy      sql.NullInt64
    UpdatedAt      sql.NullTime
    UpdatedBy      sql.NullInt64
    Deleted        sql.NullBool
}