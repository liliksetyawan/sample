package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dto/in"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexlogger/log"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
)

func (s *wardes_profileService) InsertwardesProfile(
	ctx *context.ContextModel,
	param model.URLParam,
	dto interface{},
) (
	header map[string]string,
	output interface{},
	err error,
) {
	var (
		tx                      *sql.Tx
		dtoIn                   in.WardesProfileDTOIn
		wardesProfileModel      repository.WardesProfileModel
		wardesProfileImageModel repository.WardesProfileImageModel
		result                  out.WardesProfileDTOOut
		schema                  = "nexchief"
	)

	defer func() {
		if err != nil && tx != nil {
			_ = tx.Rollback()
		} else if err == nil && tx != nil {
			err = tx.Commit()
		}
	}()

	dtoIn = s.parseDTO(dto).(in.WardesProfileDTOIn)

	tx, err = s.nexchiefDB.Begin()
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Begin Transaction")
		return
	}

	wardesProfileModel, err = s.wardes_profileDAO.InsertWardesProfile(
		tx,
		repository.WardesProfileModel{
			NexchiefAccountId:         sql.NullInt64{Int64: dtoIn.NexchiefAccountId, Valid: true},
			UserId:                    sql.NullInt64{Int64: dtoIn.UserId, Valid: true},
			Username:                  sql.NullString{String: dtoIn.Username, Valid: true},
			PersonalName:              sql.NullString{String: dtoIn.PersonalName, Valid: true},
			Gender:                    sql.NullString{String: dtoIn.Gender, Valid: true},
			Phone:                     sql.NullString{String: dtoIn.Phone, Valid: true},
			Email:                     sql.NullString{String: dtoIn.Email, Valid: true},
			ProvinceId:                sql.NullInt64{Int64: dtoIn.ProvinceId, Valid: true},
			DistrictId:                sql.NullInt64{Int64: dtoIn.DistrictId, Valid: true},
			SubDistrictId:             sql.NullInt64{Int64: dtoIn.SubDistrictId, Valid: true},
			UrbanVillageId:            sql.NullInt64{Int64: dtoIn.UrbanVillageId, Valid: true},
			Nik:                       sql.NullString{String: dtoIn.Nik, Valid: true},
			Age:                       sql.NullString{String: dtoIn.Age, Valid: true},
			Hamlet:                    sql.NullString{String: dtoIn.Hamlet, Valid: true},
			Neighbour:                 sql.NullString{String: dtoIn.Neighbour, Valid: true},
			VillageFamily:             sql.NullString{String: dtoIn.VillageFamily, Valid: true},
			VillagePopulation:         sql.NullString{String: dtoIn.VillagePopulation, Valid: true},
			BuildingLocation:          sql.NullString{String: dtoIn.BuildingLocation, Valid: true},
			BuildingOwnership:         sql.NullString{String: dtoIn.BuildingOwnership, Valid: true},
			BuildingType:              sql.NullString{String: dtoIn.BuildingType, Valid: true},
			BuildingForm:              sql.NullString{String: dtoIn.BuildingForm, Valid: true},
			BuildingWidthFirstFloor:   sql.NullFloat64{Float64: dtoIn.BuildingWidthFirstFloor, Valid: true},
			BuildingLengthFirstFloor:  sql.NullFloat64{Float64: dtoIn.BuildingLengthFirstFloor, Valid: true},
			BuildingHeightFirstFloor:  sql.NullFloat64{Float64: dtoIn.BuildingHeightFirstFloor, Valid: true},
			BuildingWidthSecondFloor:  sql.NullFloat64{Float64: dtoIn.BuildingWidthSecondFloor, Valid: true},
			BuildingLengthSecondFloor: sql.NullFloat64{Float64: dtoIn.BuildingLengthSecondFloor, Valid: true},
			BuildingHeightSecondFloor: sql.NullFloat64{Float64: dtoIn.BuildingHeightSecondFloor, Valid: true},
			ParkingLotWidth:           sql.NullFloat64{Float64: dtoIn.ParkingLotWidth, Valid: true},
			ParkingLotLength:          sql.NullFloat64{Float64: dtoIn.ParkingLotLength, Valid: true},
			VillageInternetConnection: sql.NullString{String: dtoIn.VillageInternetConnection, Valid: true},
			InvestmentCapital:         sql.NullString{String: dtoIn.InvestmentCapital, Valid: true},
			InvestmentCapitalSource:   sql.NullString{String: dtoIn.InvestmentCapitalSource, Valid: true},
			OwnershipStatus:           sql.NullString{String: dtoIn.OwnershipStatus, Valid: true},
			OwnershipFixedAsset:       sql.NullString{String: dtoIn.OwnershipFixedAsset, Valid: true},
			PermissionCover:           sql.NullString{String: dtoIn.PermissionCover, Valid: true},
			ProfileCompletionStatus:   sql.NullString{String: dtoIn.ProfileCompletionStatus, Valid: true},
			CompanyProfileId:          sql.NullInt64{Int64: dtoIn.CompanyProfileId, Valid: true},
			Code:                      sql.NullString{String: dtoIn.Code, Valid: true},
			Schema:                    sql.NullString{String: dtoIn.Schema, Valid: true},
			Status:                    sql.NullString{String: dtoIn.Status, Valid: true},
			ActiveDate:                sql.NullTime{Time: dtoIn.ActiveDate, Valid: true},
			ResignDate:                sql.NullTime{Time: dtoIn.ResignDate, Valid: true},
			IsNexwise:                 sql.NullBool{Bool: dtoIn.IsNexwise, Valid: true},
			DeadlineDate:              sql.NullTime{Time: dtoIn.DeadlineDate, Valid: true},
			NewProfileApprovalStatus:  sql.NullString{String: dtoIn.NewProfileApprovalStatus, Valid: true},
			CreatedBy:                 sql.NullInt64{Int64: ctx.Limitation.UserID, Valid: true},
			CreatedClient:             sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
			UpdatedBy:                 sql.NullInt64{Int64: ctx.Limitation.UserID, Valid: true},
			UpdatedClient:             sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile data")
		return
	}

	wardesProfileImageModel, err = s.wardes_profile_imageDAO.InsertWardesProfileImage(
		tx,
		repository.WardesProfileImageModel{
			NexchiefAccountId: sql.NullInt64{Int64: dtoIn.NexchiefAccountId, Valid: true},
			WardesProfileId:   sql.NullInt64{Int64: wardesProfileModel.ID.Int64, Valid: true},
			Type:              sql.NullString{String: "", Valid: false},
			PathImage:         sql.NullString{String: "", Valid: false},
			CreatedBy:         sql.NullInt64{Int64: ctx.Limitation.UserID, Valid: true},
			CreatedClient:     sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
			UpdatedBy:         sql.NullInt64{Int64: ctx.Limitation.UserID, Valid: true},
			UpdatedClient:     sql.NullString{String: ctx.AuthAccessTokenModel.ClientID, Valid: true},
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile_image data")
		return
	}

	result = out.WardesProfileDTOOut{
		Id:                        wardesProfileModel.ID.Int64,
		UuidKey:                   wardesProfileModel.UUIDKey.String,
		NexchiefAccountId:         wardesProfileModel.NexchiefAccountId.Int64,
		UserId:                    wardesProfileModel.UserId.Int64,
		Username:                  wardesProfileModel.Username.String,
		PersonalName:              wardesProfileModel.PersonalName.String,
		Gender:                    wardesProfileModel.Gender.String,
		Phone:                     wardesProfileModel.Phone.String,
		Email:                     wardesProfileModel.Email.String,
		ProvinceId:                wardesProfileModel.ProvinceId.Int64,
		DistrictId:                wardesProfileModel.DistrictId.Int64,
		SubDistrictId:             wardesProfileModel.SubDistrictId.Int64,
		UrbanVillageId:            wardesProfileModel.UrbanVillageId.Int64,
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
		CompanyProfileId:          wardesProfileModel.CompanyProfileId.Int64,
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
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_INSERTWARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return
}
