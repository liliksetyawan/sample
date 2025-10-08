package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/database"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	commonError "github.com/nexsoft-git/nexcommon/error"
	"github.com/nexsoft-git/nexcommon/model"
	"github.com/nexsoft-git/nexlogger/log"
	"nexsoft.co.id/example/dto/in"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
	"nexsoft.co.id/example/util/http_connector/user_auth_api"
)

func (s *wardes_profileService) Insertwardesprofile(
	ctx *context.ContextModel,
	param model.URLParam,
	dto interface{},
) (
	header map[string]string,
	output interface{},
	err error,
) {
	var (
		tx                      database.Transaction
		dtoIn                   in.WardesProfileDTOIn
		wardesProfileModel      repository.WardesProfileModel
		wardesProfileImageModel repository.WardesProfileImageModel
		result                  out.WardesProfileDTOOut
	)

	// Parse DTO
	dtoIn = s.parseDTO(dto).(in.WardesProfileDTOIn)

	// Call User Auth API to upload user data
	pathParam := user_auth_api.PostInternalUsersUploadPathParamRequest{}
	headers := user_auth_api.PostInternalUsersUploadHeaderRequest{
		Authorization: ctx.AuthAccessTokenModel.Token,
		XRequestId:    ctx.RequestID,
	}
	body := user_auth_api.PostInternalUsersUploadBodyRequest{
		UserUploadMultipart: user_auth_api.UserUploadMultipart{
			Data: user_auth_api.UserAuthRequest{
				Username:              dtoIn.Username,
				FirstName:             dtoIn.PersonalName,
				Email:                 dtoIn.Email,
				Phone:                 dtoIn.Phone,
				CountryCode:           "ID",
				ResourceId:            "",
				Scope:                 "",
				Locale:                ctx.AuthAccessTokenModel.Locale,
				AdditionalInformation: []string{},
				EmailMessage:          "",
				EmailLinkMessage:      "",
				PhoneMessage:          "",
			},
			Photo: model.FileURLPath{},
		},
	}

	userAuthResponse, statusCode, err := s.externalAPI.UserAuthApi.PostInternalUsersUpload(
		ctx.Context,
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
		err = commonError.GenerateError("API call failed", statusCode)
		return
	}

	// Begin transaction
	tx, err = database.BeginTransaction(ctx, s.database)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Begin Transaction")
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	// Insert data into wardes_profile using wardes_profileDAO.InsertWardesProfile
	wardesProfileModel, err = s.wardesProfileDAO.InsertWardesProfile(
		tx,
		repository.WardesProfileModel{
			NexchiefAccountID: database.NewNullInt64(dtoIn.NexchiefAccountId),
			UserID:            database.NewNullInt64(dtoIn.UserId),
			Username:          database.NewNullString(dtoIn.Username),
			PersonalName:      database.NewNullString(dtoIn.PersonalName),
			Gender:            database.NewNullString(dtoIn.Gender),
			Phone:             database.NewNullString(dtoIn.Phone),
			Email:             database.NewNullString(dtoIn.Email),
			CreatedBy:         database.NewNullInt64(ctx.AuthAccessTokenModel.ResourceUserID),
			CreatedClient:     database.NewNullString(ctx.AuthAccessTokenModel.ClientID),
			UpdatedBy:         database.NewNullInt64(ctx.AuthAccessTokenModel.ResourceUserID),
			UpdatedClient:     database.NewNullString(ctx.AuthAccessTokenModel.ClientID),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile data")
		return
	}

	// Insert data into wardes_profile_image using wardes_profile_imageDAO.InsertWardesProfileImage
	wardesProfileImageModel, err = s.wardesProfileImageDAO.InsertWardesProfileImage(
		tx,
		repository.WardesProfileImageModel{
			NexchiefAccountID: database.NewNullInt64(dtoIn.NexchiefAccountId),
			WardesProfileID:   database.NewNullInt64(wardesProfileModel.ID.Int64),
			CreatedBy:         database.NewNullInt64(ctx.AuthAccessTokenModel.ResourceUserID),
			CreatedClient:     database.NewNullString(ctx.AuthAccessTokenModel.ClientID),
			UpdatedBy:         database.NewNullInt64(ctx.AuthAccessTokenModel.ResourceUserID),
			UpdatedClient:     database.NewNullString(ctx.AuthAccessTokenModel.ClientID),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Inserting wardes_profile_image data")
		return
	}

	// Construct result
	result = out.WardesProfileDTOOut{
		Id:                        wardesProfileModel.ID.Int64,
		UuidKey:                   wardesProfileModel.UUIDKey.String,
		NexchiefAccountId:         wardesProfileModel.NexchiefAccountID.Int64,
		UserId:                    wardesProfileModel.UserID.Int64,
		Username:                  wardesProfileModel.Username.String,
		PersonalName:              wardesProfileModel.PersonalName.String,
		Gender:                    wardesProfileModel.Gender.String,
		Phone:                     wardesProfileModel.Phone.String,
		Email:                     wardesProfileModel.Email.String,
		ProvinceId:                wardesProfileModel.ProvinceID.Int64,
		DistrictId:                wardesProfileModel.DistrictID.Int64,
		SubDistrictId:             wardesProfileModel.SubDistrictID.Int64,
		UrbanVillageId:            wardesProfileModel.UrbanVillageID.Int64,
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
		CompanyProfileId:          wardesProfileModel.CompanyProfileID.Int64,
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
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_INSERTWARDESPROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return
}
