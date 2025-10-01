package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/error"
    "github.com/nexsoft-git/nexlogger/log"
    "nexsoft.co.id/example/dto/out"
    "nexsoft.co.id/example/repository"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    // TODO: Add other import package if necessary. Delete this line of TODO if you not adding other import package
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

    dtoIn := s.parseDTO(dto)

    // TODO: Implement the default function logic here
    // @Descriptions:
    /*
    - call external API to insert data user auth
    - handle response
    - return success or error response*/
    // @Audit Enabled: False
    // @DTO Type: default
    // @Method Type: default

    // TODO: Implement the API function logic here
    // ##API: 
    // -> description:
    //    - call external API to insert data user auth

    result := "" // Placeholder for the result of the function

    output = out2.DefaultResponsePayloadMessage{
        Status: out2.DefaultResponsePayloadStatus{
            Code:    "OK",
            Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_INSERT_WARDES_PROFILE_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
        },
        Data: result,
    }

    return
}
