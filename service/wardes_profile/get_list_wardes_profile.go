package wardes_profile

import (
	"github.com/nexsoft-git/nexcommon/context"
	"github.com/nexsoft-git/nexcommon/dto/in"
	out2 "github.com/nexsoft-git/nexcommon/dto/out"
	"github.com/nexsoft-git/nexcommon/model"
	"nexsoft.co.id/example/dto/out"
	"nexsoft.co.id/example/repository"
	// TODO: Add other import package if necessary. Delete this line of TODO if you not adding other import package
)

func (s *wardes_profileService) GetListWardesProfile(
	ctx *context.ContextModel,
	searchParam []model.SearchParam,
	dto in.GetListRequest,
) (
	header map[string]string,
	output interface{},
	err error,
) {

	// TODO: Implement the getlist function logic here
	// @Descriptions:
	/*
	   - call wardes_profileDAO.GetListWardesProfile()
	   - handle response
	   - return list of wardes profiles
	*/
	// @Method Type: getlist

	dbResult, err := s.wardes_profileDAO.GetListWardesProfile(ctx, dto, searchParam)
	if err != nil {
		return
	}

	var result []out.ListWardesProfileDTOOut
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.WardesProfileModel)
		result = append(result, out.ListWardesProfileDTOOut{
			Id:                repo.Id.Int64,
			Uuidkey:           repo.UuidKey.String,
			Nexchiefaccountid: repo.NexchiefAccountId.Int64,
			Userid:            repo.UserId.Int64,
			Username:          repo.Username.String,
			Personalname:      repo.PersonalName.String,
			Gender:            repo.Gender.String,
		})
	}

	output = out2.DefaultResponsePayloadMessage{
		Status: out2.DefaultResponsePayloadStatus{
			Code:    "OK",
			Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_GET_LIST_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
		},
		Data: result,
	}

	return
}
