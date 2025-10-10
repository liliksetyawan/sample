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

// Descriptions: Get a list of wardes profile records with pagination and search
func (d *wardes_profilePostgresqlSQLDAO) GetListWardesProfile(
     ctx *context.ContextModel,
     dtoIn in.GetListRequest,
     searchParams []model.SearchParam,
) (
     []interface{}, 
     error, 
) {
    query := fmt.Sprintf(`
        SELECT wp.id, wp.uuid_key, wp.nexchief_account_id, wp.user_id, wp.username, wp.personal_name, 
               wp.gender, wp.phone, wp.email, wp.province_id, wp.district_id, wp.sub_district_id, 
               wp.urban_village_id, wp.nik, wp.age, wp.hamlet, wp.neighbour, wp.village_family, 
               wp.village_population, wp.building_location, wp.building_ownership, wp.building_type, 
               wp.building_form, wp.building_width_first_floor, wp.building_length_first_floor, 
               wp.building_height_first_floor, wp.building_width_second_floor, wp.building_length_second_floor, 
               wp.building_height_second_floor, wp.parking_lot_width, wp.parking_lot_length, 
               wp.village_internet_connection, wp.investment_capital, wp.investment_capital_source, 
               wp.ownership_status, wp.ownership_fixed_asset, wp.permission_cover, wp.profile_completion_status, 
               wp.company_profile_id, wp.code, wp.schema, wp.status, wp.active_date, wp.resign_date, 
               wp.is_nexwise, wp.deadline_date, wp.created_by, wp.created_client, wp.created_at, 
               wp.updated_by, wp.updated_at, wp.updated_client, wp.deleted, wp.new_profile_approval_status,
               wpi.id, wpi.uuid_key, wpi.nexchief_account_id, wpi.wardes_profile_id, wpi.type, 
               wpi.path_image, wpi.created_by, wpi.created_client, wpi.created_at, wpi.updated_by, 
               wpi.updated_at, wpi.updated_client, wpi.deleted,
               u.id, u.uuid_key, u.auth_user_id, u.person_profile_id, u.client_id, u.user_code, 
               u.username, u.last_token, u.signature_key, u.ip_whitelist, u.device, u.level, 
               u.locale, u.additional_info, u.alias, u.is_user_nex_acc, u.mobile_access, 
               u.website_access, u.wdi_portal_access, u.domain_nd6, u.username_nd6, u.password_nd6, 
               u.domain_nexvin, u.username_nexvin, u.password_nexvin, u.created_by, u.created_client, 
               u.created_at, u.updated_by, u.updated_at, u.updated_client, u.deleted
        FROM %s wp
        LEFT JOIN %s wpi ON wp.id = wpi.wardes_profile_id
        LEFT JOIN %s u ON wp.user_id = u.id
        WHERE wp.deleted = false
    `, dao.GetDBTable(ctx, "wardes_profile"), dao.GetDBTable(ctx, "wardes_profile_image"), dao.GetDBTable(ctx, "user"))

    param := d.NewGetListDataParam(
        query,
        dtoIn,
        searchParams, func(rows *sql.Rows) (interface{}, error) {
            var temp repository.GetListWardesProfileModel
            err := rows.Scan(
                &temp.WardesProfile.ID, &temp.WardesProfile.UUIDKey, &temp.WardesProfile.NexchiefAccountID, 
                &temp.WardesProfile.UserID, &temp.WardesProfile.Username, &temp.WardesProfile.PersonalName,
                &temp.WardesProfile.Gender, &temp.WardesProfile.Phone, &temp.WardesProfile.Email,
                &temp.WardesProfile.ProvinceID, &temp.WardesProfile.DistrictID, &temp.WardesProfile.SubDistrictID,
                &temp.WardesProfile.UrbanVillageID, &temp.WardesProfile.NIK, &temp.WardesProfile.Age,
                &temp.WardesProfile.Hamlet, &temp.WardesProfile.Neighbour, &temp.WardesProfile.VillageFamily,
                &temp.WardesProfile.VillagePopulation, &temp.WardesProfile.BuildingLocation,
                &temp.WardesProfile.BuildingOwnership, &temp.WardesProfile.BuildingType,
                &temp.WardesProfile.BuildingForm, &temp.WardesProfile.BuildingWidthFirstFloor,
                &temp.WardesProfile.BuildingLengthFirstFloor, &temp.WardesProfile.BuildingHeightFirstFloor,
                &temp.WardesProfile.BuildingWidthSecondFloor, &temp.WardesProfile.BuildingLengthSecondFloor,
                &temp.WardesProfile.BuildingHeightSecondFloor, &temp.WardesProfile.ParkingLotWidth,
                &temp.WardesProfile.ParkingLotLength, &temp.WardesProfile.VillageInternetConnection,
                &temp.WardesProfile.InvestmentCapital, &temp.WardesProfile.InvestmentCapitalSource,
                &temp.WardesProfile.OwnershipStatus, &temp.WardesProfile.OwnershipFixedAsset,
                &temp.WardesProfile.PermissionCover, &temp.WardesProfile.ProfileCompletionStatus,
                &temp.WardesProfile.CompanyProfileID, &temp.WardesProfile.Code, &temp.WardesProfile.Schema,
                &temp.WardesProfile.Status, &temp.WardesProfile.ActiveDate, &temp.WardesProfile.ResignDate,
                &temp.WardesProfile.IsNexwise, &temp.WardesProfile.DeadlineDate, &temp.WardesProfile.CreatedBy,
                &temp.WardesProfile.CreatedClient, &temp.WardesProfile.CreatedAt, &temp.WardesProfile.UpdatedBy,
                &temp.WardesProfile.UpdatedAt, &temp.WardesProfile.UpdatedClient, &temp.WardesProfile.Deleted,
                &temp.WardesProfile.NewProfileApprovalStatus,
                &temp.WardesProfileImage.ID, &temp.WardesProfileImage.UUIDKey, &temp.WardesProfileImage.NexchiefAccountID,
                &temp.WardesProfileImage.WardesProfileID, &temp.WardesProfileImage.Type, &temp.WardesProfileImage.PathImage,
                &temp.WardesProfileImage.CreatedBy, &temp.WardesProfileImage.CreatedClient, &temp.WardesProfileImage.CreatedAt,
                &temp.WardesProfileImage.UpdatedBy, &temp.WardesProfileImage.UpdatedAt, &temp.WardesProfileImage.UpdatedClient,
                &temp.WardesProfileImage.Deleted,
                &temp.User.ID, &temp.User.UUIDKey, &temp.User.AuthUserID, &temp.User.PersonProfileID, &temp.User.ClientID,
                &temp.User.UserCode, &temp.User.Username, &temp.User.LastToken, &temp.User.SignatureKey, &temp.User.IPWhitelist,
                &temp.User.Device, &temp.User.Level, &temp.User.Locale, &temp.User.AdditionalInfo, &temp.User.Alias,
                &temp.User.IsUserNexAcc, &temp.User.MobileAccess, &temp.User.WebsiteAccess, &temp.User.WDIPortalAccess,
                &temp.User.DomainND6, &temp.User.UsernameND6, &temp.User.PasswordND6, &temp.User.DomainNexvin,
                &temp.User.UsernameNexvin, &temp.User.PasswordNexvin, &temp.User.CreatedBy, &temp.User.CreatedClient,
                &temp.User.CreatedAt, &temp.User.UpdatedBy, &temp.User.UpdatedAt, &temp.User.UpdatedClient, &temp.User.Deleted,
            )
            return temp, err
        },
    ).Prefix(
        dao.NewParamPrefix().
            Add("id", "wp"),
    ).CreatedBy(ctx.Limitation.UserID)

    return d.GetListDataWithDefaultMustCheck(param)
}
