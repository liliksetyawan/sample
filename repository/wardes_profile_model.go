package repository

import "database/sql"

type WardesProfileModel struct {
	ID                        sql.NullInt64
	UUIDKey                   sql.NullString
	NexchiefAccountID         sql.NullInt64
	UserID                    sql.NullInt64
	Username                  sql.NullString
	PersonalName              sql.NullString
	Gender                    sql.NullString
	Phone                     sql.NullString
	Email                     sql.NullString
	ProvinceID                sql.NullInt64
	DistrictID                sql.NullInt64
	SubDistrictID             sql.NullInt64
	UrbanVillageID            sql.NullInt64
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
	CompanyProfileID          sql.NullInt64
	Code                      sql.NullString
	Schema                    sql.NullString
	Status                    sql.NullString
	ActiveDate                sql.NullTime
	ResignDate                sql.NullTime
	IsNexwise                 sql.NullBool
	DeadlineDate              sql.NullTime
	CreatedBy                 sql.NullInt64
	CreatedClient             sql.NullString
	CreatedAt                 sql.NullTime
	UpdatedBy                 sql.NullInt64
	UpdatedAt                 sql.NullTime
	UpdatedClient             sql.NullString
	Deleted                   sql.NullBool
	NewProfileApprovalStatus  sql.NullString
}

type UserWardesProfileModel struct {
	WardesProfileModel
	WardesProfileImageModel
}
