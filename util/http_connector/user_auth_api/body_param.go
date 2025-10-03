package user_auth_api

// Body Parameter Structs
// Body parameters for post_/persons
type PostPersonsBodyRequest struct {
	PersonCreateMultipart
}

// Body parameters for get_/persons/{uuid}
type GetPersonsUuidBodyRequest struct {
}
