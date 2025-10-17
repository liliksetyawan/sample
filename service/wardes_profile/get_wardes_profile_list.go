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

func (s *wardes_profileService) GetWardesProfileList(
	ctx *context.ContextModel,
	searchParam []model.SearchParam,
	dto in.GetListRequest,
) (
	header map[string]string,
	output interface{},
	err error,
) {

	// Parse and validate request DTO (List/filter/order/limit/pagination) from input
	if err = dto.Validate(); err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Validating DTO")
		return
	}

	// Build DAO query params (filters, order_by, limit, offset) based on DTO
	queryParams := repository.WardesProfileQueryParams{
		Limit:   dto.Limit,
		Offset:  dto.Offset,
		OrderBy: dto.OrderBy,
		Filters: searchParam,
	}

	// Call wardes_profileDAO.GetListWardesProfile(ctx, queryParams)
	dbResult, err := s.wardesProfileDAO.GetListWardesProfile(ctx, queryParams)
	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Getting Wardes Profile List")
		return
	}

	// Transform DAO result set into ListWardesProfileDTOOut format
	var result []out.ListWardesProfileDTOOut
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.WardesProfileModel)
		result = append(result, out.ListWardesProfileDTOOut{
			ID:              repo.ID.Int64,
			Username:        repo.Username.String,
			PersonalName:    repo.PersonalName.String,
			Gender:          repo.Gender.String,
			Phone:           repo.Phone.String,
			Email:           repo.Email.String,
			WardesProfileID: repo.ID.Int64,
			Type:            "",
			PathImage:       "",
			UserID:          repo.UserID.Int64,
			AuthUserID:      0,
			ClientID:        0,
		})
	}

	output = out2.DefaultResponsePayloadMessage{
		Status: out2.DefaultResponsePayloadStatus{
			Code:    "OK",
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_GET_WARDES_PROFILE_LIST_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return

}
