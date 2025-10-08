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

// Push user auth
func (c *UserAuthApiClient) PostInternalUsers (
	ctx context.Context,
	pathParam PostInternalUsersPathParamRequest,
	headers PostInternalUsersHeaderRequest,
	body PostInternalUsersBodyRequest,
) (
    PostInternalUsersResponse,
    int,
    error,
) {
	path := "/internal/users"

	var response PostInternalUsersResponse
	statusCode, err := c.apiConnector.HitAPI(
		ctx,
		"POST",
		c.baseURL,
		c.apiConnector.ParseURL(path, &pathParam),
		c.apiConnector.ParseHeader(&headers),
		&body,
		&response,
	)

	return response, statusCode, err
}

// Push user auth
func (c *UserAuthApiClient) GetInternalUsers (
	ctx context.Context,
	pathParam GetInternalUsersPathParamRequest,
	headers GetInternalUsersHeaderRequest,
) (
    GetInternalUsersResponse,
    int,
    error,
) {
	path := "/internal/users"

	var response GetInternalUsersResponse
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

// Push user auth
func (c *UserAuthApiClient) GetInternalUsersId (
	ctx context.Context,
	pathParam GetInternalUsersIdPathParamRequest,
	headers GetInternalUsersIdHeaderRequest,
) (
    GetInternalUsersIdResponse,
    int,
    error,
) {
	path := "/internal/users/{ID}"

	var response GetInternalUsersIdResponse
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

// Upload user auth
func (c *UserAuthApiClient) PostInternalUsersUpload (
	ctx context.Context,
	pathParam PostInternalUsersUploadPathParamRequest,
	headers PostInternalUsersUploadHeaderRequest,
	body PostInternalUsersUploadBodyRequest,
) (
    PostInternalUsersUploadResponse,
    int,
    error,
) {
	path := "/internal/users/upload"

	var response PostInternalUsersUploadResponse
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