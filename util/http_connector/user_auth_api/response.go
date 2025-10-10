package user_auth_api

// Response Structs
// Response struct for post_/internal/users
type PostInternalUsersResponse struct {
	HTTPClientDTO
	NexsoftResponse
}

// Response struct for get_/internal/users
type GetInternalUsersResponse struct {
	HTTPClientDTO
	NexsoftResponse
}

// Response struct for get_/internal/users/{ID}
type GetInternalUsersIdResponse struct {
	HTTPClientDTO
	NexsoftResponse
}

// Response struct for post_/internal/users/upload
type PostInternalUsersUploadResponse struct {
	HTTPClientDTO
	NexsoftResponse
}

