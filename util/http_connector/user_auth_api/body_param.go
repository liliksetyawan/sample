package user_auth_api

// Body Parameter Structs
// Body parameters for post_/internal/users
type PostInternalUsersBodyRequest struct {
	UserAuthRequest
}

// Body parameters for get_/internal/users
type GetInternalUsersBodyRequest struct {
}

// Body parameters for get_/internal/users/{ID}
type GetInternalUsersIdBodyRequest struct {
}

// Body parameters for post_/internal/users/upload
type PostInternalUsersUploadBodyRequest struct {
	UserUploadMultipart
}
