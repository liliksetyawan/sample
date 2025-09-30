package repository

import (
	"database/sql"
)

type WardesProfileModel struct {
	Id                        sql.NullInt64
	UuidKey                   sql.NullString
	NexchiefAccountId         sql.NullInt64
	UserId                    sql.NullInt64
	Username                  sql.NullString
	PersonalName              sql.NullString
	Gender                    sql.NullString
	Phone                     sql.NullString
	Email                     sql.NullString
	ProvinceId                sql.NullInt64
	DistrictId                sql.NullInt64
	SubDistrictId             sql.NullInt64
	UrbanVillageId            sql.NullInt64
	Nik                       sql.NullString
	Age                       sql.NullString
	Hamlet                    sql.NullString
	Neighbour                 sql.NullString
	VillageFamily             sql.NullString
	VillagePopulation         sql.NullString
	BuildingLocation          sql.NullString
	BuildingOwnership         sql.NullString
	BuildingType              sql.NullString
	BuildingForm              sql.NullString
	BuildingWidthFirstFloor   sql.NullFloat64
	BuildingLengthFirstFloor  sql.NullFloat64
	BuildingHeightFirstFloor  sql.NullFloat64
	BuildingWidthSecondFloor  sql.NullFloat64
	BuildingLengthSecondFloor sql.NullFloat64
	BuildingHeightSecondFloor sql.NullFloat64
	ParkingLotWidth           sql.NullFloat64
	ParkingLotLength          sql.NullFloat64
	VillageInternetConnection sql.NullString
	InvestmentCapital         sql.NullString
	InvestmentCapitalSource   sql.NullString
	OwnershipStatus           sql.NullString
	OwnershipFixedAsset       sql.NullString
	PermissionCover           sql.NullString
	ProfileCompletionStatus   sql.NullString
	CompanyProfileId          sql.NullInt64
	Code                      sql.NullString
	Schema                    sql.NullString
	Status                    sql.NullString
	ActiveDate                sql.NullString
	ResignDate                sql.NullString
	IsNexwise                 sql.NullBool
	DeadlineDate              sql.NullString
	CreatedBy                 sql.NullInt64
	CreatedClient             sql.NullString
	CreatedAt                 sql.NullString
	UpdatedBy                 sql.NullInt64
	UpdatedAt                 sql.NullString
	UpdatedClient             sql.NullString
	Deleted                   sql.NullBool
	NewProfileApprovalStatus  sql.NullString
}

type WardesProfileImageDetailModel struct {
	Id                sql.NullInt64
	UuidKey           sql.NullString
	NexchiefAccountId sql.NullInt64
	WardesProfileId   sql.NullInt64
	Type              sql.NullString
	PathImage         sql.NullString
	CreatedBy         sql.NullInt64
	CreatedClient     sql.NullString
	CreatedAt         sql.NullString
	UpdatedBy         sql.NullInt64
	UpdatedAt         sql.NullString
	UpdatedClient     sql.NullString
	Deleted           sql.NullBool
}
