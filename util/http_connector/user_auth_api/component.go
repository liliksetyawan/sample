package user_auth_api

// Component Schemas
// Generated from components.schemas.PersonRequest
type PersonRequest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Job string `json:"job"`
}

// Generated from components.schemas.PersonCreateMultipart
type PersonCreateMultipart struct {
	Data PersonRequest `json:"data"`
	Photo model.FileURLPath `json:"photo"`
}

// Generated from components.schemas.UuidPayload
type UuidPayload struct {
	Uuid string `json:"uuid"`
}

// Generated from components.schemas.PersonPayload
type PersonPayload struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Job string `json:"job"`
}

// Generated from components.schemas.ErrorPayload
type ErrorPayload struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

// Generated from components.schemas.OkUuidResponse
type OkUuidResponse struct {
	Status string `json:"status"`
	Payload UuidPayload `json:"payload"`
}

// Generated from components.schemas.OkPersonResponse
type OkPersonResponse struct {
	Status string `json:"status"`
	Payload PersonPayload `json:"payload"`
}

// Generated from components.schemas.ErrorResponse
type ErrorResponse struct {
	Status string `json:"status"`
	Payload ErrorPayload `json:"payload"`
}
