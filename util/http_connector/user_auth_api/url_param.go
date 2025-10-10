package user_auth_api

// Path and Query Parameter Structs
// Path and query parameters for post_/internal/users
type PostInternalUsersPathParamRequest struct {
}

// Path and query parameters for get_/internal/users
type GetInternalUsersPathParamRequest struct {
	Page int `field:"query" name:"page"`
	Limit int `field:"query" name:"limit"`
}

// Path and query parameters for get_/internal/users/{ID}
type GetInternalUsersIdPathParamRequest struct {
	Id string `field:"path" name:"ID"`
}

// Path and query parameters for post_/internal/users/upload
type PostInternalUsersUploadPathParamRequest struct {
}
