package wardes_profile

import (
	"database/sql"
	"time"

	"github.com/nexsoft-git/nexcommon/context"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexcommon/services/audit_helper"
	"github.com/nexsoft-git/nexlogger/log"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
	"nexsoft.co.id/example/util/http_connector/user_auth_api"
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

func (s *wardes_profileService) doInsertWardesProfile(
	ctx *context.ContextModel,
	tx *sql.Tx,
	dto interface{},
	timeNow time.Time,
) (
	output interface{},
	audit []audit_helper.AuditSystemModel,
	err error,
) {
	dtoIn := s.parseDTO(dto)

	// Call external API user_auth_api.PostInternalUserUpload
	pathParam := user_auth_api.PostInternalUsersUploadPathParamRequest{}
	headers := user_auth_api.PostInternalUsersUploadHeaderRequest{
		Authorization: ctx.AuthAccessTokenModel.AccessToken,
		XRequestId:    ctx.RequestID,
	}
	body := user_auth_api.PostInternalUsersUploadBodyRequest{
		UserUploadMultipart: user_auth_api.UserUploadMultipart{
			Data: user_auth_api.UserAuthRequest{
				Username:  dtoIn.Username,
				FirstName: dtoIn.PersonalName,
				Email:     dtoIn.Email,
				Phone:     dtoIn.Phone,
				Locale:    ctx.AuthAccessTokenModel.Locale,
			},
		},
	}

	userAuthResponse, statusCode, err := s.httpConnector.UserAuthApi.PostInternalUsersUpload(
		ctx.ToContext(),
		pathParam,
		headers,
		body,
	)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error Found When Calling User Auth API")
		return
	}

	if statusCode != 200 && statusCode != 201 {
		log.Error().Int("statusCode", statusCode).Caller().Msg("User Auth API returned non-success status")
		err = errors.New("failed to upload user data")
		return
	}

	var userModel repository.UserModel
	userModel, err = s.userDAO.InsertUser(
		ctx,
		tx,
		repository.UserModel{
			AuthUserID:    sql.NullInt64{Int64: int64(userAuthResponse.Payload.Data.Content.UserId)},
			ClientID:      sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ClientID},
			UserCode:      sql.NullString{String: dtoIn.Code},
			Username:      sql.NullString{String: dtoIn.Username},
			CreatedAt:     sql.NullTime{Time: timeNow},
			CreatedBy:     sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			CreatedClient: sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ClientID},
			UpdatedAt:     sql.NullTime{Time: timeNow},
			UpdatedBy:     sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			UpdatedClient: sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ClientID},
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting user data")
		return
	}

	audit = append(audit, audit_helper.AuditSystemModel{
		TableName:  "user",
		PrimaryKey: userModel.ID.Int64,
		Action:     audit_helper.AuditServiceActionInsert,
		SchemaName: "nexchief",
	})

	var wardesProfileModel repository.WardesProfileModel
	wardesProfileModel, err = s.wardesProfileDAO.InsertWardesProfile(
		ctx,
		tx,
		repository.WardesProfileModel{
			NexchiefAccountID:         sql.NullInt64{Int64: dtoIn.NexchiefAccountID},
			UserID:                    sql.NullInt64{Int64: userModel.ID.Int64},
			Username:                  sql.NullString{String: dtoIn.Username},
			PersonalName:              sql.NullString{String: dtoIn.PersonalName},
			Gender:                    sql.NullString{String: dtoIn.Gender},
			Phone:                     sql.NullString{String: dtoIn.Phone},
			Email:                     sql.NullString{String: dtoIn.Email},
			ProvinceID:                sql.NullInt64{Int64: dtoIn.ProvinceID},
			DistrictID:                sql.NullInt64{Int64: dtoIn.DistrictID},
			SubDistrictID:             sql.NullInt64{Int64: dtoIn.SubDistrictID},
			UrbanVillageID:            sql.NullInt64{Int64: dtoIn.UrbanVillageID},
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
			CompanyProfileID:          sql.NullInt64{Int64: dtoIn.CompanyProfileID},
			Code:                      sql.NullString{String: dtoIn.Code},
			Schema:                    sql.NullString{String: dtoIn.Schema},
			Status:                    sql.NullString{String: dtoIn.Status},
			ActiveDate:                sql.NullTime{Time: dtoIn.ActiveDate},
			ResignDate:                sql.NullTime{Time: dtoIn.ResignDate},
			IsNexwise:                 sql.NullBool{Bool: dtoIn.IsNexwise},
			DeadlineDate:              sql.NullTime{Time: dtoIn.DeadlineDate},
			NewProfileApprovalStatus:  sql.NullString{String: dtoIn.NewProfileApprovalStatus},
			CreatedAt:                 sql.NullTime{Time: timeNow},
			CreatedBy:                 sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			CreatedClient:             sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ClientID},
			UpdatedAt:                 sql.NullTime{Time: timeNow},
			UpdatedBy:                 sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			UpdatedClient:             sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ClientID},
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile data")
		return
	}

	audit = append(audit, audit_helper.AuditSystemModel{
		TableName:  "wardes_profile",
		PrimaryKey: wardesProfileModel.ID.Int64,
		Action:     audit_helper.AuditServiceActionInsert,
		SchemaName: "nexchief",
	})

	var wardesProfileImageModel repository.WardesProfileImageModel
	wardesProfileImageModel, err = s.wardesProfileImageDAO.InsertWardesProfileImage(
		ctx,
		tx,
		repository.WardesProfileImageModel{
			NexchiefAccountID: sql.NullInt64{Int64: dtoIn.NexchiefAccountID},
			WardesProfileID:   sql.NullInt64{Int64: wardesProfileModel.ID.Int64},
			Type:              sql.NullString{String: "profile"},
			PathImage:         sql.NullString{String: ""},
			CreatedAt:         sql.NullTime{Time: timeNow},
			CreatedBy:         sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			CreatedClient:     sql.NullString{String: ctx.AuthAccessTokenModel.ClientID},
			UpdatedAt:         sql.NullTime{Time: timeNow},
			UpdatedBy:         sql.NullInt64{Int64: ctx.AuthAccessTokenModel.ResourceUserID},
			UpdatedClient:     sql.NullString{String: ctx.AuthAccessTokenModel.ClientID},
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile_image data")
		return
	}

	audit = append(audit, audit_helper.AuditSystemModel{
		TableName:  "wardes_profile_image",
		PrimaryKey: wardesProfileImageModel.ID.Int64,
		Action:     audit_helper.AuditServiceActionInsert,
		SchemaName: "nexchief",
	})

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
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_INSERT_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return
}
