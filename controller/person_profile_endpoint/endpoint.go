package person_profile_endpoint

import(
	"fmt"
	"net/http"
	http_validator "github.com/nexsoft-git/nexcommon/controller/http"
	"github.com/your_git/app_name/config"
	person_profileService "github.com/your_git/app_name/service/person_profile"
)

func NewPersonProfileEndpoint(
	httpValidator *http_validator.HTTPController,
	srv person_profileService.PersonProfileService,
	config config.Configuration,
) *personProfileEndpoint {
	return &person_profileEndpoint{
		httpValidator   : httpValidator,
		srv             : srv,
		config          : config,
	}
}

type personProfileEndpoint struct {
	httpValidator *http_validator.HTTPController
	srv           person_profileService.PersonProfileService
	config        config.Configuration
}

func (e *personProfileEndpoint) RegisterEndpoint() { 

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
		    fmt.Sprintf("/v1/%s/person_profile", e.config.Server.ResourceID),
			e.httpValidator.WrapService(
				http_validator.NewWarpServiceParam(
					e.srv,
					e.srv.InsertPersonProfile,
					e.httpValidator.UserAccessValidatorWithKong,
				).Menu("insert").
				Permission("master.person_profile.person_profile:insert"),
			),
			http.MethodPost, http.MethodOptions,
		),
	)

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
			fmt.Sprintf("/v1/%s/person_profile", e.config.Server.ResourceID),
			e.httpValidator.WrapServiceListData(
				http_validator.NewWarpGetListServiceParam(
					e.srv,
					e.srv.ListPersonProfiles,
					e.httpValidator.UserAccessValidatorWithKong,
				).Permission("master.person_profile.person_profile:view"),
			),
			http.MethodGet, http.MethodOptions,
		),
	)

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
			fmt.Sprintf("/v1/%s/person_profile/count", e.config.Server.ResourceID),
			e.httpValidator.WrapServiceListData(
				http_validator.NewWarpGetListServiceParam(
					e.srv,
					e.srv.CountPersonProfiles,
					e.httpValidator.UserAccessValidatorWithKong,
				).Permission("master.person_profile.person_profile:view"),
			),
			http.MethodGet, http.MethodOptions,
		),
	)
}