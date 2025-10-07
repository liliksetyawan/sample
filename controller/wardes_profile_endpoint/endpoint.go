package wardes_profile_endpoint

import (
	"fmt"
	http_validator "github.com/nexsoft-git/nexcommon/controller/http"
	"net/http"
	"nexsoft.co.id/example/config"
	wardes_profileService "nexsoft.co.id/example/service/wardes_profile"
)

func NewWardesProfileEndpoint(
	httpValidator *http_validator.HTTPController,
	srv wardes_profileService.WardesProfileService,
	config config.Configuration,
) *wardesProfileEndpoint {
	return &wardesProfileEndpoint{
		httpValidator: httpValidator,
		srv:           srv,
		config:        config,
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
					Permisson("[master].[wardes_profile].[wardes_profile]:insert"),
			),
			http.MethodPost, http.MethodOptions,
		),
	)
}
