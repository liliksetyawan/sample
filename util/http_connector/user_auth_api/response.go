package user_auth_api

// Response Structs
// Response struct for post_/persons
type PostPersonsResponse struct {
	HTTPClientDTO
	OkUuidResponse
}

// Response struct for get_/persons/{uuid}
type GetPersonsUuidResponse struct {
	HTTPClientDTO
	OkPersonResponse
}

