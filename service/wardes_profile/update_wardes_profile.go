package wardes_profile

import (
	"database/sql"
	"time"

	"github.com/nexsoft-git/nexcommon/context"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexcommon/services/audit_helper"
	"github.com/nexsoft-git/nexlogger/log"
	"nexsoft.co.id/example/dto/in"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
)

func (s *wardes_profileService) UpdateWardesProfile(
	ctx *context.ContextModel,
	param model.URLParam,
	dto interface{},
) (
	header map[string]string,
	output interface{},
	err error,
) {
	audit, err := s.auditHelper.InitAuditService(ctx, dto, s.doUpdateWardesProfile)
	if err != nil {
		return
	}

	output, err = audit.CompletedAuditData()
	return
}

func (s *wardes_profileService) doUpdateWardesProfile(
	ctx *context.ContextModel,
	tx *sql.Tx,
	dto interface{},
	timeNow time.Time,
) (
	output interface{},
	audit []audit_helper.AuditSystemModel,
	err error,
) {
	dtoIn := s.parseDTO(dto).(in.WardesProfileDTOIn)
	schema := "nexchief"

	wardesProfileOnDB, err := s.wardesProfileDAO.GetWardesProfile(ctx, tx, dtoIn.UUIDKey)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Get Data with the uuid_key wardes_profile data")
		return
	}

	if wardesProfileOnDB.ID.Int64 == 0 {
		err = commonError.ErrUnknownData.Param("id")
		return
	}

	if wardesProfileOnDB.UpdatedAt.Time != dtoIn.UpdatedAt {
		err = commonError.ErrDataLocked.Param("id")
		return
	}

	auditTemp, err := audit_helper.GetDataForAuditByIDTx(
		ctx,
		tx,
		audit_helper.AuditServiceActionUpdate,
		"wardes_profile",
		wardesProfileOnDB.ID.Int64,
		ctx.Limitation.UserID,
		schema,
	)
	if err != nil {
		return
	}
	audit = append(audit, auditTemp...)

	wardesProfileModel, err := s.wardesProfileDAO.UpdateWardesProfileStatus(
		ctx,
		tx,
		repository.WardesProfileModel{
			UUIDKey:       sql.NullString{String: dtoIn.UUIDKey},
			UpdatedAt:     sql.NullTime{Time: timeNow},
			UpdatedBy:     sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			UpdatedClient: sql.NullString{String: ctx.AuthAccessTokenModel.ClientID},
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Updating wardes_profile data")
		return
	}

	result := out.WardesProfileDTOOut{
		ID:                        wardesProfileModel.ID.Int64,
		UUIDKey:                   wardesProfileModel.UUIDKey.String,
		NexchiefAccountID:         wardesProfileModel.NexchiefAccountID.Int64,
		UserID:                    wardesProfileModel.UserID.Int64,
		Username:                  wardesProfileModel.Username.String,
		PersonalName:              wardesProfileModel.PersonalName.String,
		Gender:                    wardesProfileModel.Gender.String,
		Phone:                     wardesProfileModel.Phone.String,
		Email:                     wardesProfileModel.Email.String,
		ProvinceID:                wardesProfileModel.ProvinceID.Int64,
		DistrictID:                wardesProfileModel.DistrictID.Int64,
		SubDistrictID:             wardesProfileModel.SubDistrictID.Int64,
		UrbanVillageID:            wardesProfileModel.UrbanVillageID.Int64,
		Nik:                       wardesProfileModel.Nik.String,
		Age:                       wardesProfileModel.Age.String,
		Hamlet:                    wardesProfileModel.Hamlet.String,
		Neighbour:                 wardesProfileModel.Neighbour.String,
		VillageFamily:             wardesProfileModel.VillageFamily.String,
		VillagePopulation:         wardesProfileModel.VillagePopulation.String,
		BuildingLocation:          wardesProfileModel.BuildingLocation.String,
		BuildingOwnership:         wardesProfileModel.BuildingOwnership.String,
		BuildingType:              wardesProfileModel.BuildingType.String,
		BuildingForm:              wardesProfileModel.BuildingForm.String,
		BuildingWidthFirstFloor:   wardesProfileModel.BuildingWidthFirstFloor.Float64,
		BuildingLengthFirstFloor:  wardesProfileModel.BuildingLengthFirstFloor.Float64,
		BuildingHeightFirstFloor:  wardesProfileModel.BuildingHeightFirstFloor.Float64,
		BuildingWidthSecondFloor:  wardesProfileModel.BuildingWidthSecondFloor.Float64,
		BuildingLengthSecondFloor: wardesProfileModel.BuildingLengthSecondFloor.Float64,
		BuildingHeightSecondFloor: wardesProfileModel.BuildingHeightSecondFloor.Float64,
		ParkingLotWidth:           wardesProfileModel.ParkingLotWidth.Float64,
		ParkingLotLength:          wardesProfileModel.ParkingLotLength.Float64,
		VillageInternetConnection: wardesProfileModel.VillageInternetConnection.String,
		InvestmentCapital:         wardesProfileModel.InvestmentCapital.String,
		InvestmentCapitalSource:   wardesProfileModel.InvestmentCapitalSource.String,
		OwnershipStatus:           wardesProfileModel.OwnershipStatus.String,
		OwnershipFixedAsset:       wardesProfileModel.OwnershipFixedAsset.String,
		PermissionCover:           wardesProfileModel.PermissionCover.String,
		ProfileCompletionStatus:   wardesProfileModel.ProfileCompletionStatus.String,
		CompanyProfileID:          wardesProfileModel.CompanyProfileID.Int64,
		Code:                      wardesProfileModel.Code.String,
		Schema:                    wardesProfileModel.Schema.String,
		Status:                    wardesProfileModel.Status.String,
		ActiveDate:                wardesProfileModel.ActiveDate.Time,
		ResignDate:                wardesProfileModel.ResignDate.Time,
		IsNexwise:                 wardesProfileModel.IsNexwise.Bool,
		DeadlineDate:              wardesProfileModel.DeadlineDate.Time,
		CreatedBy:                 wardesProfileModel.CreatedBy.Int64,
		CreatedClient:             wardesProfileModel.CreatedClient.String,
		CreatedAt:                 wardesProfileModel.CreatedAt.Time,
		UpdatedBy:                 wardesProfileModel.UpdatedBy.Int64,
		UpdatedAt:                 wardesProfileModel.UpdatedAt.Time,
		UpdatedClient:             wardesProfileModel.UpdatedClient.String,
		Deleted:                   wardesProfileModel.Deleted.Bool,
		NewProfileApprovalStatus:  wardesProfileModel.NewProfileApprovalStatus.String,
	}

	output = out2.DefaultResponsePayloadMessage{
		Status: out2.DefaultResponsePayloadStatus{
			Code:    "OK",
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_UPDATE_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return
}
