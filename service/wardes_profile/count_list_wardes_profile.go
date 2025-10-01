package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/services"
    "nexsoft.co.id/example/repository"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    // TODO: Add other import package if necessary. Delete this line of TODO if you not adding other import package
)

func (s *wardes_profileService) CountListWardesProfile(
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
        - count the number of profiles
        - return count
    */
    // @Method Type: getlist

    count, err := s.wardes_profileDAO.CountListWardesProfile(ctx, dto, searchParam)
    if err != nil {
        return
    }

    result := services.ToInitiateGetListData(s)
    result.CountData = count

    output = out2.DefaultResponsePayloadMessage{
        Status: out2.DefaultResponsePayloadStatus{
            Code:    "OK",
            Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_COUNT_LIST_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
        },
        Data: result,
    }

    return
}