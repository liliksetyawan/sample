package http_connector

import (
	httpClient "github.com/nexsoft-git/nexcommon/http/client"
	"util/http_connector/user_auth_api"
	"util/http_connector/user_auth_api"
)

func NewExternalAPI(
	connector httpClient.APIConnector,
	user_auth_apiBaseURL string,
	user_auth_apiBaseURL string,
) ExternalAPI {
	return ExternalAPI{
		UserAuthApi: user_auth_api.NewUserAuthApiClient(connector, user_auth_apiBaseURL),
		UserAuthApi: user_auth_api.NewUserAuthApiClient(connector, user_auth_apiBaseURL),
	}
}

type ExternalAPI struct {
	UserAuthApi user_auth_api.UserAuthApiClient
	UserAuthApi user_auth_api.UserAuthApiClient
}