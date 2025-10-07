package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/model"
	commonError "github.com/nexsoft-git/nexcommon/error"
	"github.com/nexsoft-git/nexlogger/log"
	"nexsoft.co.id/example/dto/in"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	"golang.org/x/net/context"
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
	var (
		tx *sql.Tx
		dtoIn in.WardesProfileDTOIn
		wardesProfileOnDB repository.WardesProfileModel
		wardesProfileModel repository.WardesProfileModel
		schema = "nexchief"
	)

	// Parse DTO
	dtoIn = s.parseDTO(dto).(in.WardesProfileDTOIn)

	// Begin transaction
	tx, err = s.BeginTx(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Begin Transaction")
		return
	}
	defer s.RollbackTx(tx, err)

	// Get existing data by UUID
	wardesProfileOnDB, err = s.wardes_profile.GetForUpdate(ctx, tx, dtoIn.UUIDKey)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Get Data with the uuid_key wardes_profile data")
		return
	}

	// Check if data exists
	if wardesProfileOnDB.ID.Int64 == 0 {
		err = commonError.ErrUnknownData.Param("id")
		return
	}

	// Check optimistic lock
	if wardesProfileOnDB.UpdatedAt.Time != dtoIn.UpdatedAt {
		err = commonError.ErrDataLocked.Param("id")
		return
	}

	// Update wardes profile data
	wardesProfileModel = repository.WardesProfileModel{
		ID:                        sql.NullInt64{Int64: wardesProfileOnDB.ID.Int64},
		UUIDKey:                   sql.NullString{String: dtoIn.UUIDKey},
		NexchiefAccountID:         sql.NullInt64{Int64: dtoIn.NexchiefAccountId},
		UserID:                    sql.NullInt64{Int64: dtoIn.UserId},
		Username:                  sql.NullString{String: dtoIn.Username},
		PersonalName:              sql.NullString{String: dtoIn.PersonalName},
		Gender:                    sql.NullString{String: dtoIn.Gender},
		Phone:                     sql.NullString{String: dtoIn.Phone},
		Email:                     sql.NullString{String: dtoIn.Email},
		ProvinceID:                sql.NullInt64{Int64: dtoIn.ProvinceId},
		DistrictID:                sql.NullInt64{Int64: dtoIn.DistrictId},
		SubDistrictID:             sql.NullInt64{Int64: dtoIn.SubDistrictId},
		UrbanVillageID:            sql.NullInt64{Int64: dtoIn.UrbanVillageId},
		Nik:                       sql.NullString{String: dtoIn.Nik},
		Age:                       sql.NullString{String: dtoIn.Age},
		Hamlet:                    sql.NullString{String: dtoIn.Hamlet},
		Neighbour:                 sql.NullString{String: dtoIn.Neighbour},
		VillageFamily:             sql.NullString{String: dtoIn.VillageFamily},
		VillagePopulation:         sql.NullString{String: dtoIn.VillagePopulation},
		BuildingLocation:          sql.NullString{String: dtoIn.BuildingLocation},
		BuildingOwnership:         sql.NullString{String: dtoIn.BuildingOwnership},
		BuildingType:              sql.NullString{String: dtoIn.BuildingType},
		BuildingForm:              sql.NullString{String: dtoIn.BuildingForm},
		BuildingWidthFirstFloor:   sql.NullFloat64{Float64: dtoIn.BuildingWidthFirstFloor},
		BuildingLengthFirstFloor:  sql.NullFloat64{Float64: dtoIn.BuildingLengthFirstFloor},
		BuildingHeightFirstFloor:  sql.NullFloat64{Float64: dtoIn.BuildingHeightFirstFloor},
		BuildingWidthSecondFloor:  sql.NullFloat64{Float64: dtoIn.BuildingWidthSecondFloor},
		BuildingLengthSecondFloor: sql.NullFloat64{Float64: dtoIn.BuildingLengthSecondFloor},
		BuildingHeightSecondFloor: sql.NullFloat64{Float64: dtoIn.BuildingHeightSecondFloor},
		ParkingLotWidth:           sql.NullFloat64{Float64: dtoIn.ParkingLotWidth},
		ParkingLotLength:          sql.NullFloat64{Float64: dtoIn.ParkingLotLength},
		VillageInternetConnection: sql.NullString{String: dtoIn.VillageInternetConnection},
		InvestmentCapital:         sql.NullString{String: dtoIn.InvestmentCapital},
		InvestmentCapitalSource:   sql.NullString{String: dtoIn.InvestmentCapitalSource},
		OwnershipStatus:           sql.NullString{String: dtoIn.OwnershipStatus},
		OwnershipFixedAsset:       sql.NullString{String: dtoIn.OwnershipFixedAsset},
		PermissionCover:           sql.NullString{String: dtoIn.PermissionCover},
		ProfileCompletionStatus:   sql.NullString{String: dtoIn.ProfileCompletionStatus},
		CompanyProfileID:          sql.NullInt64{Int64: dtoIn.CompanyProfileId},
		Code:                      sql.NullString{String: dtoIn.Code},
		Schema:                    sql.NullString{String: dtoIn.Schema},
		Status:                    sql.NullString{String: dtoIn.Status},
		ActiveDate:                sql.NullTime{Time: dtoIn.ActiveDate},
		ResignDate:                sql.NullTime{Time: dtoIn.ResignDate},
		IsNexwise:                 sql.NullBool{Bool: dtoIn.IsNexwise},
		DeadlineDate:              sql.NullTime{Time: dtoIn.DeadlineDate},
		NewProfileApprovalStatus:  sql.NullString{String: dtoIn.NewProfileApprovalStatus},
		UpdatedBy:                 sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
		UpdatedClient:             sql.NullString{String: ctx.AuthAccessTokenModel.ClientID},
		UpdatedAt:                 sql.NullTime{Time: dtoIn.UpdatedAt},
	}

	_, err = s.wardes_profile.Update(tx, wardesProfileModel)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Updating wardes_profile data")
		return
	}

	// Commit transaction
	err = s.CommitTx(tx)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Commit Transaction")
		return
	}

	// Construct result
	output = out2.DefaultResponsePayloadMessage{
		Status: out2.DefaultResponsePayloadStatus{
			Code:    "OK",
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_UPDATE_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: nil,
	}

	return
}