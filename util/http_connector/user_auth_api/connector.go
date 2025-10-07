package user_auth_api

import (
	"context"
	"net/http"
	httpClient "github.com/nexsoft-git/nexcommon/http/client"
	"github.com/nexsoft-git/nexcommon/model"
)

// API Client
type UserAuthApiClient struct {
	apiConnector httpClient.APIConnector
	baseURL      string
}

// Constructor
func NewUserAuthApiClient(connector httpClient.APIConnector, baseURL string) *UserAuthApiClient {
	return &UserAuthApiClient{
		apiConnector: connector,
		baseURL:      baseURL,
	}
}

// Push person data (multipart with JSON + photo)
func (c *UserAuthApiClient) PostPersons (
	ctx context.Context,
	pathParam PostPersonsPathParamRequest,
	headers PostPersonsHeaderRequest,
	body PostPersonsBodyRequest,
) (
    PostPersonsResponse,
    int,
    error,
) {
	path := "/persons"

	var response PostPersonsResponse
	statusCode, err := c.apiConnector.HitAPIMultipart(
		ctx,
		"POST",
		c.baseURL,
		c.apiConnector.ParseURL(path, &pathParam),
		c.apiConnector.ParseHeader(&headers),
		c.apiConnector.ParseMultipartStruct(&body),
		&response,
	)

	return response, statusCode, err
}

// Retrieve person data by UUID
func (c *UserAuthApiClient) GetPersonsUuid (
	ctx context.Context,
	pathParam GetPersonsUuidPathParamRequest,
	headers GetPersonsUuidHeaderRequest,
) (
    GetPersonsUuidResponse,
    int,
    error,
) {
	path := "/persons/{uuid}"

	var response GetPersonsUuidResponse
	statusCode, err := c.apiConnector.HitAPI(
		ctx,
		"GET",
		c.baseURL,
		c.apiConnector.ParseURL(path, &pathParam),
		c.apiConnector.ParseHeader(&headers),
		nil,
		&response,
	)

	return response, statusCode, err
}