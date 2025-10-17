package wardes_profile

import (
    "database/sql"
    "time"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/services/audit_helper"
    "nexsoft.co.id/example/dto/out"
    "nexsoft.co.id/example/repository"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    "github.com/nexsoft-git/nexlogger/log"
    "nexsoft.co.id/example/dto/in"
    "github.com/nexsoft-git/nexcommon/constanta"
    "github.com/nexsoft-git/nexcommon/services/commonError"
)

func (s *wardes_profileService) InsertWardesProfile(
    ctx *context.ContextModel,
    param model.URLParam,
    dto interface{},
) (
    header map[string]string,
    output interface{},
    err error,
) {
    audit, err := s.auditHelper.InitAuditService(ctx, dto, s.doInsertWardesProfile)
    if err != nil {
        return
    }

    output, err = audit.CompletedAuditData()
    return
}

func (s *wardes_profileService) doInsertWardesProfile (
    ctx *context.ContextModel,
    tx *sql.Tx,
    dto interface{},
    timeNow time.Time,
) (
    output interface{},
    audit []audit_helper.AuditSystemModel,
    err error,
) { 
    var (
        dtoIn in.WardesProfileDTOIn
        wardesProfileModel repository.WardesProfileModel
        schema = "nexchief"
    )

    // Parse DTO
    dtoIn = s.parseDTO(dto).(in.WardesProfileDTOIn)

    // Map DTO to entity structure
    wardesProfileModel = repository.WardesProfileModel{
        NexchiefAccountID:        sql.NullInt64{Int64: dtoIn.NexchiefAccountID, Valid: true},
        UserID:                   sql.NullInt64{Int64: dtoIn.UserID, Valid: true},
        Username:                 sql.NullString{String: dtoIn.Username, Valid: true},
        PersonalName:             sql.NullString{String: dtoIn.PersonalName, Valid: true},
        Gender:                   sql.NullString{String: dtoIn.Gender, Valid: true},
        Phone:                    sql.NullString{String: dtoIn.Phone, Valid: true},
        Email:                    sql.NullString{String: dtoIn.Email, Valid: true},
        CreatedAt:                sql.NullTime{Time: timeNow, Valid: true},
        CreatedBy:                sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID, Valid: true},
        CreatedClient:            sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
        UpdatedAt:                sql.NullTime{Time: timeNow, Valid: true},
        UpdatedBy:                sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID, Valid: true},
        UpdatedClient:            sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
    }

    // Insert wardes_profile
    wardesProfileModel, err = s.wardesProfileDAO.InsertWardesProfile(
        ctx,
        tx,
        wardesProfileModel,
    )

    if err != nil {
        log.Error().
            Err(err).
            Caller().
            Msg("Error Found When Inserting wardes_profile data")
        return
    }

    // Add audit entry for insertion
    audit = append(audit, audit_helper.AuditSystemModel{
        TableName:  "wardes_profile",
        PrimaryKey: wardesProfileModel.ID.Int64,
        Action:     audit_helper.AuditServiceActionInsert,
        SchemaName: schema,
    })

    // Map inserted entity to DTO out
    result := out.WardesProfileDTOOut{
        ID:                         wardesProfileModel.ID.Int64,
        UUIDKey:                    wardesProfileModel.UUIDKey.String,
        NexchiefAccountID:          wardesProfileModel.NexchiefAccountID.Int64,
        UserID:                     wardesProfileModel.UserID.Int64,
        Username:                   wardesProfileModel.Username.String,
        PersonalName:               wardesProfileModel.PersonalName.String,
        Gender:                     wardesProfileModel.Gender.String,
        Phone:                      wardesProfileModel.Phone.String,
        Email:                      wardesProfileModel.Email.String,
        ProvinceID:                 wardesProfileModel.ProvinceID.Int64,
        DistrictID:                 wardesProfileModel.DistrictID.Int64,
        SubDistrictID:              wardesProfileModel.SubDistrictID.Int64,
        UrbanVillageID:             wardesProfileModel.UrbanVillageID.Int64,
        Nik:                        wardesProfileModel.Nik.String,
        Age:                        wardesProfileModel.Age.String,
        Hamlet:                     wardesProfileModel.Hamlet.String,
        Neighbour:                  wardesProfileModel.Neighbour.String,
        VillageFamily:              wardesProfileModel.VillageFamily.String,
        VillagePopulation:          wardesProfileModel.VillagePopulation.String,
        BuildingLocation:           wardesProfileModel.BuildingLocation.String,
        BuildingOwnership:          wardesProfileModel.BuildingOwnership.String,
        BuildingType:               wardesProfileModel.BuildingType.String,
        BuildingForm:               wardesProfileModel.BuildingForm.String,
        BuildingWidthFirstFloor:    wardesProfileModel.BuildingWidthFirstFloor.Float64,
        BuildingLengthFirstFloor:   wardesProfileModel.BuildingLengthFirstFloor.Float64,
        BuildingHeightFirstFloor:   wardesProfileModel.BuildingHeightFirstFloor.Float64,
        BuildingWidthSecondFloor:   wardesProfileModel.BuildingWidthSecondFloor.Float64,
        BuildingLengthSecondFloor:  wardesProfileModel.BuildingLengthSecondFloor.Float64,
        BuildingHeightSecondFloor:  wardesProfileModel.BuildingHeightSecondFloor.Float64,
        ParkingLotWidth:            wardesProfileModel.ParkingLotWidth.Float64,
        ParkingLotLength:           wardesProfileModel.ParkingLotLength.Float64,
        VillageInternetConnection:  wardesProfileModel.VillageInternetConnection.String,
        InvestmentCapital:          wardesProfileModel.InvestmentCapital.String,
        InvestmentCapitalSource:    wardesProfileModel.InvestmentCapitalSource.String,
        OwnershipStatus:            wardesProfileModel.OwnershipStatus.String,
        OwnershipFixedAsset:        wardesProfileModel.OwnershipFixedAsset.String,
        PermissionCover:            wardesProfileModel.PermissionCover.String,
        ProfileCompletionStatus:    wardesProfileModel.ProfileCompletionStatus.String,
        CompanyProfileID:           wardesProfileModel.CompanyProfileID.Int64,
        Code:                       wardesProfileModel.Code.String,
        Schema:                     wardesProfileModel.Schema.String,
        Status:                     wardesProfileModel.Status.String,
        ActiveDate:                 wardesProfileModel.ActiveDate.Time,
        ResignDate:                 wardesProfileModel.ResignDate.Time,
        IsNexwise:                  wardesProfileModel.IsNexwise.Bool,
        DeadlineDate:               wardesProfileModel.DeadlineDate.Time,
        CreatedBy:                  wardesProfileModel.CreatedBy.Int64,
        CreatedClient:              wardesProfileModel.CreatedClient.String,
        CreatedAt:                  wardesProfileModel.CreatedAt.Time,
        UpdatedBy:                  wardesProfileModel.UpdatedBy.Int64,
        UpdatedAt:                  wardesProfileModel.UpdatedAt.Time,
        UpdatedClient:              wardesProfileModel.UpdatedClient.String,
        Deleted:                    wardesProfileModel.Deleted.Bool,
        NewProfileApprovalStatus:   wardesProfileModel.NewProfileApprovalStatus.String,
    }

    output = out2.DefaultResponsePayloadMessage{
        Status: out2.DefaultResponsePayloadStatus{
            Code:    "OK",
            Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_INSERT_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
        },
        Data: result,
    }

    return
}