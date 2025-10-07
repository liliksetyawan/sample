package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update a new wardes profile record
func (d *wardes_profilePostgresqlSQLDAO) UpdateWardesProfile(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.WardesProfileModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            nexchief_account_id = $1, user_id = $2, username = $3, personal_name = $4,
            gender = $5, phone = $6, email = $7, sub_district_id = $8,
            urban_village_id = $9, nik = $10, age = $11, hamlet = $12,
            neighbour = $13, village_family = $14, village_population = $15,
            building_location = $16, building_ownership = $17, building_type = $18,
            building_form = $19, building_width_first_floor = $20, building_length_first_floor = $21,
            building_height_first_floor = $22, building_width_second_floor = $23, building_length_second_floor = $24,
            building_height_second_floor = $25, parking_lot_width = $26, parking_lot_length = $27,
            village_internet_connection = $28, investment_capital = $29, investment_capital_source = $30,
            ownership_status = $31, ownership_fixed_asset = $32, permission_cover = $33,
            profile_completion_status = $34, company_profile_id = $35, code = $36,
            schema = $37, status = $38, active_date = $39, resign_date = $40,
            is_nexwise = $41, deadline_date = $42, updated_by = $43, updated_at = $44,
            updated_client = $45, deleted = $46, new_profile_approval_status = $47
        WHERE uuid_key = $48
    `, dao.GetDBTable(ctx, "wardes_profile"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.NexchiefAccountID.Int64, dtoIn.UserID.Int64, dtoIn.Username.String, dtoIn.PersonalName.String,
        dtoIn.Gender.String, dtoIn.Phone.String, dtoIn.Email.String, dtoIn.SubDistrictID.Int64,
        dtoIn.UrbanVillageID.Int64, dtoIn.NIK.String, dtoIn.Age.String, dtoIn.Hamlet.String,
        dtoIn.Neighbour.String, dtoIn.VillageFamily.String, dtoIn.VillagePopulation.String,
        dtoIn.BuildingLocation.String, dtoIn.BuildingOwnership.String, dtoIn.BuildingType.String,
        dtoIn.BuildingForm.String, dtoIn.BuildingWidthFirstFloor.Float64, dtoIn.BuildingLengthFirstFloor.Float64,
        dtoIn.BuildingHeightFirstFloor.Float64, dtoIn.BuildingWidthSecondFloor.Float64, dtoIn.BuildingLengthSecondFloor.Float64,
        dtoIn.BuildingHeightSecondFloor.Float64, dtoIn.ParkingLotWidth.Float64, dtoIn.ParkingLotLength.Float64,
        dtoIn.VillageInternetConnection.String, dtoIn.InvestmentCapital.String, dtoIn.InvestmentCapitalSource.String,
        dtoIn.OwnershipStatus.String, dtoIn.OwnershipFixedAsset.String, dtoIn.PermissionCover.String,
        dtoIn.ProfileCompletionStatus.String, dtoIn.CompanyProfileID.Int64, dtoIn.Code.String,
        dtoIn.Schema.String, dtoIn.Status.String, dtoIn.ActiveDate.Time, dtoIn.ResignDate.Time,
        dtoIn.IsNexwise.Bool, dtoIn.DeadlineDate.Time, ctx.Limitation.ServiceUserID, dtoIn.UpdatedAt.Time,
        ctx.AuthAccessTokenModel.ClientID, dtoIn.Deleted.Bool, dtoIn.NewProfileApprovalStatus.String,
        dtoIn.UUIDKey.String,
    )
    return err
}
