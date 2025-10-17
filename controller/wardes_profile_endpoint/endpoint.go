package wardes_profile_endpoint

import(
	"fmt"
	"net/http"
	http_validator "github.com/nexsoft-git/nexcommon/controller/http"
	"nexsoft.co.id/example/config"
	wardes_profileService "nexsoft.co.id/example/service/wardes_profile"
)

func NewWardesProfileEndpoint(
	httpValidator *http_validator.HTTPController,
	srv wardes_profileService.WardesProfileService,
	config config.Configuration,
) *wardesProfileEndpoint {
	return &wardesProfileEndpoint{
		httpValidator   : httpValidator,
		srv             : srv,
		config          : config,
	}
}

type wardesProfileEndpoint struct {
	httpValidator *http_validator.HTTPController
	srv           wardes_profileService.WardesProfileService
	config        config.Configuration
}

func (e *wardesProfileEndpoint) RegisterEndpoint() { 

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
		    fmt.Sprintf("/v1/%s/wardes_profile", e.config.Server.ResourceID),
			e.httpValidator.WrapService(
				http_validator.NewWarpServiceParam(
					e.srv,
					e.srv.InsertWardesProfile,
					e.httpValidator.UserAccessValidatorWithKong,
				).Menu("insert").
				Permisson("master.wardes.wardes_profile:insert"),
			),
			http.MethodPost, http.MethodOptions,
		),
	)

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
			fmt.Sprintf("/v1/%s/wardes_profile", e.config.Server.ResourceID),
			e.httpValidator.WrapServiceListData(
				http_validator.NewWarpGetListServiceParam(
					e.srv,
					e.srv.GetWardesProfileList,
					e.httpValidator.UserAccessValidatorWithKong,
				).Permisson("master.wardes.wardes_profile:view"),
			),
			http.MethodGet, http.MethodOptions,
		),
	)

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
		    fmt.Sprintf("/v1/%s/wardes_profile/count", e.config.Server.ResourceID),
			e.httpValidator.WrapService(
				http_validator.NewWarpServiceParam(
					e.srv,
					e.srv.CountWardesProfile,
					e.httpValidator.UserAccessValidatorWithKong,
				).Menu("count").
				Permisson("master.wardes.wardes_profile:view"),
			),
			http.MethodGet, http.MethodOptions,
		),
	)
}