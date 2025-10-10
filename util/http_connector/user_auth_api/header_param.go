package user_auth_api

// Header Parameter Structs
// Header parameters for post_/internal/users
type PostInternalUsersHeaderRequest struct {
	Authorization string `header:"Authorization"`
	XRequestId string `header:"X-Request-ID"`
}

// Header parameters for get_/internal/users
type GetInternalUsersHeaderRequest struct {
	Authorization string `header:"Authorization"`
	XRequestId string `header:"X-Request-ID"`
}

// Header parameters for get_/internal/users/{ID}
type GetInternalUsersIdHeaderRequest struct {
	Authorization string `header:"Authorization"`
	XRequestId string `header:"X-Request-ID"`
}

// Header parameters for post_/internal/users/upload
type PostInternalUsersUploadHeaderRequest struct {
	Authorization string `header:"Authorization"`
	XRequestId string `header:"X-Request-ID"`
}
