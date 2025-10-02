package wardes_profile_image_endpoint

import(
	"fmt"
	"net/http"
	http_validator "github.com/nexsoft-git/nexcommon/controller/http"
	"github.com/your_git/app_name/config"
	wardes_profile_imageService "github.com/your_git/app_name/service/wardes_profile_image"
)

func NewWardesProfileImageEndpoint(
	httpValidator *http_validator.HTTPController,
	srv wardes_profile_imageService.WardesProfileImageService,
	config config.Configuration,
) *wardesProfileImageEndpoint {
	return &wardes_profile_imageEndpoint{
		httpValidator   : httpValidator,
		srv             : srv,
		config          : config,
	}
}

type wardesProfileImageEndpoint struct {
	httpValidator *http_validator.HTTPController
	srv           wardes_profile_imageService.WardesProfileImageService
	config        config.Configuration
}

func (e *wardesProfileImageEndpoint) RegisterEndpoint() { 

	e.httpValidator.HandleFunc(
		http_validator.NewHandleFuncParam(
		    fmt.Sprintf("/v1/%s/wardes_profile_image", e.config.Server.ResourceID),
			e.httpValidator.WrapService(
				http_validator.NewWarpServiceParam(
					e.srv,
					e.srv.InsertWardesProfileImage,
					e.httpValidator.UserAccessValidatorWithKong,
				).Menu("insert").
				Permission("master.wardes_profile_image.wardes_profile_image:insert"),
			),
			http.MethodPost, http.MethodOptions,
		),
	)
}