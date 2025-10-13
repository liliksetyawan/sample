package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/context"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/services"
    "github.com/nexsoft-git/nexlogger/log"
)

func (s *wardes_profileService) GetWardesProfileCount(
	ctx *context.ContextModel,
	param model.URLParam,
	dto interface{},
) (
	header map[string]string,
	output interface{},
	err error,
) {

	var dbResult int

	dbResult, err = s.wardes_profileDao.GetCountWardesProfile(
		ctx,
		dto,
	)

	if err != nil {
		log.Error().
			Err(err).
			Caller().
			Msg("Error Found When Get Count Wardes Profile data")
		return
	}

	result := services.ToInitiateGetListData(s)
	result.CountData = dbResult

	output = out2.DefaultResponsePayloadMessage{
		Status: out2.DefaultResponsePayloadStatus{
			Code:    "OK",
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_GET_WARDES_PROFILE_COUNT_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return

}
