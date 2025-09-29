package wardes_profile

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Retrieve a list of wardes profile records
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT wp.id, wp.uuid_key, wp.nexchief_account_id, wp.user_id, wp.username, 
               wp.personal_name, wp.gender, wp.phone, wp.email, wp.province_id, 
               wp.district_id, wp.sub_district_id, wp.urban_village_id, wp.nik, 
               wp.age, wp.hamlet, wp.neighbour, wp.village_family, wp.village_population, 
               wp.building_location, wp.building_ownership, wp.building_type, 
               wp.building_form, wp.building_width_first_floor, wp.building_length_first_floor, 
               wp.building_height_first_floor, wp.building_width_second_floor, 
               wp.building_length_second_floor, wp.building_height_second_floor, 
               wp.parking_lot_width, wp.parking_lot_length, 
               wp.village_internet_connection, wp.investment_capital, 
               wp.investment_capital_source, wp.ownership_status, 
               wp.ownership_fixed_asset, wp.permission_cover, 
               wp.profile_completion_status, wp.company_profile_id, 
               wp.code, wp.schema, wp.status, wp.active_date, 
               wp.resign_date, wp.is_nexwise, wp.deadline_date, 
               wp.created_by, wp.created_client, wp.created_at, 
               wp.updated_by, wp.updated_at, wp.updated_client, 
               wp.deleted, wp.new_profile_approval_status, 
               wpi.id, wpi.uuid_key, wpi.nexchief_account_id, 
               wpi.wardes_profile_id, wpi.type, wpi.path_image, 
               wpi.created_by, wpi.created_client, wpi.created_at, 
               wpi.updated_by, wpi.updated_at, wpi.updated_client, 
               wpi.deleted
        FROM %s wp
        JOIN %s wpi ON wp.id = wpi.wardes_profile_id
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "wardes_profile_image"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.UserWardesProfileModel
            err := rows.Scan(
                &temp.WardesProfileModel.ID, &temp.WardesProfileModel.UUIDKey, 
                &temp.WardesProfileModel.NexchiefAccountID, &temp.WardesProfileModel.UserID, 
                &temp.WardesProfileModel.Username, &temp.WardesProfileModel.PersonalName, 
                &temp.WardesProfileModel.Gender, &temp.WardesProfileModel.Phone, 
                &temp.WardesProfileModel.Email, &temp.WardesProfileModel.ProvinceID, 
                &temp.WardesProfileModel.DistrictID, &temp.WardesProfileModel.SubDistrictID, 
                &temp.WardesProfileModel.UrbanVillageID, &temp.WardesProfileModel.Nik, 
                &temp.WardesProfileModel.Age, &temp.WardesProfileModel.Hamlet, 
                &temp.WardesProfileModel.Neighbour, &temp.WardesProfileModel.VillageFamily, 
                &temp.WardesProfileModel.VillagePopulation, &temp.WardesProfileModel.BuildingLocation, 
                &temp.WardesProfileModel.BuildingOwnership, &temp.WardesProfileModel.BuildingType, 
                &temp.WardesProfileModel.BuildingForm, &temp.WardesProfileModel.BuildingWidthFirstFloor, 
                &temp.WardesProfileModel.BuildingLengthFirstFloor, &temp.WardesProfileModel.BuildingHeightFirstFloor, 
                &temp.WardesProfileModel.BuildingWidthSecondFloor, &temp.WardesProfileModel.BuildingLengthSecondFloor, 
                &temp.WardesProfileModel.BuildingHeightSecondFloor, &temp.WardesProfileModel.ParkingLotWidth, 
                &temp.WardesProfileModel.ParkingLotLength, &temp.WardesProfileModel.VillageInternetConnection, 
                &temp.WardesProfileModel.InvestmentCapital, &temp.WardesProfileModel.InvestmentCapitalSource, 
                &temp.WardesProfileModel.OwnershipStatus, &temp.WardesProfileModel.OwnershipFixedAsset, 
                &temp.WardesProfileModel.PermissionCover, &temp.WardesProfileModel.ProfileCompletionStatus, 
                &temp.WardesProfileModel.CompanyProfileID, &temp.WardesProfileModel.Code, 
                &temp.WardesProfileModel.Schema, &temp.WardesProfileModel.Status, 
                &temp.WardesProfileModel.ActiveDate, &temp.WardesProfileModel.ResignDate, 
                &temp.WardesProfileModel.IsNexwise, &temp.WardesProfileModel.DeadlineDate, 
                &temp.WardesProfileModel.CreatedBy, &temp.WardesProfileModel.CreatedClient, 
                &temp.WardesProfileModel.CreatedAt, &temp.WardesProfileModel.UpdatedBy, 
                &temp.WardesProfileModel.UpdatedAt, &temp.WardesProfileModel.UpdatedClient, 
                &temp.WardesProfileModel.Deleted, &temp.WardesProfileModel.NewProfileApprovalStatus,
                &temp.WardesProfileImageModel.ID, &temp.WardesProfileImageModel.UUIDKey, 
                &temp.WardesProfileImageModel.NexchiefAccountID, &temp.WardesProfileImageModel.WardesProfileID, 
                &temp.WardesProfileImageModel.Type, &temp.WardesProfileImageModel.PathImage, 
                &temp.WardesProfileImageModel.CreatedBy, &temp.WardesProfileImageModel.CreatedClient, 
                &temp.WardesProfileImageModel.CreatedAt, &temp.WardesProfileImageModel.UpdatedBy, 
                &temp.WardesProfileImageModel.UpdatedAt, &temp.WardesProfileImageModel.UpdatedClient, 
                &temp.WardesProfileImageModel.Deleted,
            )
            return temp, err
        },
    ).FieldConverter(dao.NewParaFieldConverter().
        Add("code", "wp.code").
        Add("name", "wp.personal_name"),
    )

    return d.GetListDataWithDefaultMustCheck(param)
}