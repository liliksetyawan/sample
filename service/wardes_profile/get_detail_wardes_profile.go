package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/error"
    "nexsoft.co.id/example/dto/out"
    "nexsoft.co.id/example/repository"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    // TODO: Add other import package if necessary. Delete this line of TODO if you not adding other import package
)

func (s *wardes_profileService) GetDetailWardesProfile(
    ctx *context.ContextModel,
    param model.URLParam,
    dto interface{},
) (
    header map[string]string,
    output interface{},
    err error,
) {

    // TODO: Implement the default function logic here
    // @Descriptions:
    /*
    - call wardes_profileDAO.GetDetailWardesProfile(uuid_key)
    - handle response
    - return details of the wardes profile*/
    // @Audit Enabled: False
    // @DTO Type: default
    // @Method Type: default

    uuidKey := param["uuid_key"]

    wardesProfileOnDB, err := s.wardes_profileDAO.GetDetailWardesProfile(ctx, uuidKey)
    if err != nil {
        log.Error().
            Err(err).
            Caller().
            Msg("Error Found When Get Detail Wardes Profile")
        return
    }

    if wardesProfileOnDB.Id.String == "" {
        err = commonError.ErrUnknownData.Param("uuid_key")
        return
    }

    result := out.WardesProfileDTOOut{
        Id:                wardesProfileOnDB.Id.String,
        Uuidkey:          wardesProfileOnDB.Uuidkey.String,
        Nexchiefaccountid: wardesProfileOnDB.Nexchiefaccountid.String,
        Userid:           wardesProfileOnDB.Userid.String,
        Username:         wardesProfileOnDB.Username.String,
        Personalname:     wardesProfileOnDB.Personalname.String,
        Gender:           wardesProfileOnDB.Gender.String,
    }

    output = out2.DefaultResponsePayloadMessage{
        Status: out2.DefaultResponsePayloadStatus{
            Code:    "OK",
            Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_GET_DETAIL_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
        },
        Data: result,
    }

    return
}
