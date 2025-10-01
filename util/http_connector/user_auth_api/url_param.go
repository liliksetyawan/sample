package user_auth_api

// Path and Query Parameter Structs
// Path and query parameters for post_/persons
type PostPersonsPathParamRequest struct {
}

// Path and query parameters for get_/persons/{uuid}
type GetPersonsUuidPathParamRequest struct {
	Uuid string `field:"path" name:"uuid"`
}
