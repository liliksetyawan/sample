package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: This function will insert a new wardes_profile record into the database.
func (d *wardes_profilePostgresqlSQLDAO) InsertWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     repository.WardesProfileModel, 
     error, 
) {
    query := fmt.Sprintf(`
        INSERT INTO %s (
            nexchief_account_id, user_id, username, personal_name, gender, phone, email, 
            province_id, district_id, sub_district_id, urban_village_id, nik, age, 
            hamlet, neighbour, village_family, village_population, building_location, 
            building_ownership, building_type, building_form, building_width_first_floor, 
            building_length_first_floor, building_height_first_floor, 
            building_width_second_floor, building_length_second_floor, 
            building_height_second_floor, parking_lot_width, parking_lot_length, 
            village_internet_connection, investment_capital, investment_capital_source, 
            ownership_status, ownership_fixed_asset, permission_cover, 
            profile_completion_status, company_profile_id, code, schema, status, 
            active_date, resign_date, is_nexwise, deadline_date, created_by, 
            created_client, created_at, updated_by, updated_at, updated_client, 
            deleted, new_profile_approval_status
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7,
            $8, $9, $10, $11, $12, $13, 
            $14, $15, $16, $17, $18, 
            $19, $20, $21, $22, $23, 
            $24, $25, $26, $27, $28, 
            $29, $30, $31, $32, $33, 
            $34, $35, $36, $37, $38, 
            $39, $40, $41, $42, $43, 
            $44, $45, $46, $47, $48, 
            $49, $50, $51, $52, $53
        )
        RETURNING id, uuid_key;
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return repository.WardesProfileModel{}, err
    }

    var model repository.WardesProfileModel
    err = stmt.QueryRow(
        dtoIn.NexchiefAccountID.Int64, dtoIn.UserID.Int64, dtoIn.Username.String, 
        dtoIn.PersonalName.String, dtoIn.Gender.String, dtoIn.Phone.String, 
        dtoIn.Email.String, dtoIn.ProvinceID.Int64, dtoIn.DistrictID.Int64, 
        dtoIn.SubDistrictID.Int64, dtoIn.UrbanVillageID.Int64, dtoIn.Nik.String, 
        dtoIn.Age.String, dtoIn.Hamlet.String, dtoIn.Neighbour.String, 
        dtoIn.VillageFamily.String, dtoIn.VillagePopulation.String, 
        dtoIn.BuildingLocation.String, dtoIn.BuildingOwnership.String, 
        dtoIn.BuildingType.String, dtoIn.BuildingForm.String, 
        dtoIn.BuildingWidthFirstFloor.Float64, dtoIn.BuildingLengthFirstFloor.Float64, 
        dtoIn.BuildingHeightFirstFloor.Float64, dtoIn.BuildingWidthSecondFloor.Float64, 
        dtoIn.BuildingLengthSecondFloor.Float64, dtoIn.BuildingHeightSecondFloor.Float64, 
        dtoIn.ParkingLotWidth.Float64, dtoIn.ParkingLotLength.Float64, 
        dtoIn.VillageInternetConnection.String, dtoIn.InvestmentCapital.String, 
        dtoIn.InvestmentCapitalSource.String, dtoIn.OwnershipStatus.String, 
        dtoIn.OwnershipFixedAsset.String, dtoIn.PermissionCover.String, 
        dtoIn.ProfileCompletionStatus.String, dtoIn.CompanyProfileID.Int64, 
        dtoIn.Code.String, dtoIn.Schema.String, dtoIn.Status.String, 
        dtoIn.ActiveDate.Time, dtoIn.ResignDate.Time, dtoIn.IsNexwise.Bool, 
        dtoIn.DeadlineDate.Time, ctx.Limitation.UserID, 
        ctx.AuthAccessTokenModel.ClientID, dtoIn.CreatedAt.Time, 
        ctx.Limitation.UserID, ctx.AuthAccessTokenModel.ClientID, 
        dtoIn.Deleted.Bool, dtoIn.NewProfileApprovalStatus.String,
    ).Scan(&model.ID, &model.UUIDKey)

    return model, err
}
